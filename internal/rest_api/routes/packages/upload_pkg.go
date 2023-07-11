package packages

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/uwine4850/strux_api/internal/config"
	"github.com/uwine4850/strux_api/internal/rest_api/routes/errors"
	"github.com/uwine4850/strux_api/internal/rest_api/routes/utils"
	"github.com/uwine4850/strux_api/pkg/logging"
	"github.com/uwine4850/strux_api/pkg/uplutils"
	"github.com/uwine4850/strux_api/services/protofiles/pkgproto"
	"google.golang.org/grpc"
	"net/http"
	"reflect"
	"strings"
)

func UploadPkgService(w http.ResponseWriter, r *http.Request) {
	connection, err := utils.CheckFormKeyAndGetPackageServiceConnection(w, r, []string{},
		[]string{"username", "password", "version", "files_info"})

	if reflect.DeepEqual(err, errors.ErrFormKeyNotExist{}) {
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "UploadPkgService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "UploadPkgService", "", err.Error())
			utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)

	client := pkgproto.NewPackageClient(connection)

	values, files, err := utils.GetFormData(r)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "UploadPkgService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// setting upload files
	var uplFilesData1 []*pkgproto.UploadFile
	err = uplutils.SetUploadFiles(files, &uplFilesData1)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "UploadPkgService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// parsing files_info and setting dirs info
	f := values["files_info"][0]
	reader := strings.NewReader(f)
	uplDirInfo := &pkgproto.UploadDirInfo{}
	err = json.NewDecoder(reader).Decode(&uplDirInfo)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "UploadPkgService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := &pkgproto.User{
		Username: values["username"][0],
		Password: values["password"][0],
	}

	up := &pkgproto.RequestUploadPackage{
		UplFiles:   uplFilesData1,
		UplDirInfo: uplDirInfo,
		User:       user,
		Version:    values["version"][0],
	}

	response, err := client.UploadPackage(context.Background(), up)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "UploadPkgService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = utils.CreateResponse(w, http.StatusCreated, response)

	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "UploadPkgService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
