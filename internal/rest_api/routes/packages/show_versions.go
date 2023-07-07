package packages

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"mime/multipart"
	"net/http"
	"reflect"
	"strux_api/internal/config"
	"strux_api/internal/rest_api/routes/errors"
	"strux_api/internal/rest_api/routes/utils"
	"strux_api/pkg/logging"
	"strux_api/services/protofiles/pkgproto"
)

func ShowVersions(w http.ResponseWriter, r *http.Request) {
	connection, err := utils.CheckFormKeyAndGetPackageServiceConnection(w, r, []string{}, []string{"username", "pkgName"})
	if reflect.DeepEqual(err, errors.ErrFormKeyNotExist{}) {
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "ShowVersions", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "ShowVersions", "", err.Error())
			utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)

	values, _, err := utils.GetFormData(r)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "ShowVersions", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := pkgproto.NewPackageClient(connection)
	request := &pkgproto.RequestShowVersions{
		Username: values["username"][0],
		PkgName:  values["pkgName"][0],
	}
	response, err := client.ShowVersions(context.Background(), request)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "ShowVersions", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if response.BaseResponse.Success {
		err = sendForm(w, response.Versions)
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "ShowVersions", "", err.Error())
			utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	err = utils.CreateResponse(w, http.StatusOK, response.BaseResponse)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "ShowVersions", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// sendForm sends the "versions" field with a list of all versions of the package
func sendForm(w http.ResponseWriter, versions []string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	versionJsonBytes, err := json.Marshal(versions)
	if err != nil {
		return err
	}
	err = writer.WriteField("versions", string(versionJsonBytes))
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", writer.FormDataContentType())
	_, err = w.Write(body.Bytes())
	if err != nil {
		return err
	}
	return nil
}
