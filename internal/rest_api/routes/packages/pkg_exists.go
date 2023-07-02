package packages

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"reflect"
	"strux_api/internal/config"
	"strux_api/internal/rest_api/routes/errors"
	"strux_api/internal/rest_api/routes/utils"
	"strux_api/pkg/logging"
	"strux_api/services/protofiles/baseproto"
	"strux_api/services/protofiles/pkgproto"
)

func PackageExistsService(w http.ResponseWriter, r *http.Request) {
	connection, err := utils.CheckFormKeyAndGetPackageServiceConnection(w, r, []string{}, []string{"username", "pkgName", "version"})

	if reflect.DeepEqual(err, errors.ErrFormKeyNotExist{}) {
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "PackageExistsService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "PackageExistsService", "", err.Error())
			utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)

	values, _, err := utils.GetFormData(r)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "PackageExistsService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := pkgproto.NewPackageClient(connection)
	req := &pkgproto.RequestPackageExists{
		Username: values["username"][0],
		PkgName:  values["pkgName"][0],
		Version:  values["version"][0],
	}
	existsPackage, err := client.ExistsPackage(context.Background(), req)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "PackageExistsService", "", err.Error())
		utils.SendResponseError(w, existsPackage.Message, http.StatusInternalServerError)
		return
	}
	if existsPackage.Status == baseproto.ResponseStatus_StatusError {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "PackageExistsService", "", err.Error())
		utils.SendResponseError(w, existsPackage.Message, http.StatusInternalServerError)
		return
	}
	err = utils.CreateResponse(w, http.StatusOK, existsPackage)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "PackageExistsService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
	}
}
