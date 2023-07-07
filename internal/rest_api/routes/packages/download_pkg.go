package packages

import (
	"bytes"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strux_api/internal/config"
	"strux_api/internal/rest_api/routes/errors"
	"strux_api/internal/rest_api/routes/utils"
	"strux_api/pkg/logging"
	"strux_api/pkg/upload_package"
	"strux_api/services/protofiles/baseproto"
	"strux_api/services/protofiles/pkgproto"
)

func DownloadPackage(w http.ResponseWriter, r *http.Request) {
	connection, err := utils.CheckFormKeyAndGetPackageServiceConnection(w, r, []string{}, []string{"username", "pkgName", "version"})

	if reflect.DeepEqual(err, errors.ErrFormKeyNotExist{}) {
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "DownloadPackage", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "DownloadPackage", "", err.Error())
			utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)

	values, _, err := utils.GetFormData(r)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "DownloadPackage", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := pkgproto.NewPackageClient(connection)
	request := &pkgproto.RequestDownloadPackage{
		Username: values["username"][0],
		PkgName:  values["pkgName"][0],
		Version:  values["version"][0],
	}
	downloadPackageResponse, err := client.DownloadPackage(context.Background(), request)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "DownloadPackage", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if downloadPackageResponse.BaseResponse.Status == baseproto.ResponseStatus_StatusError {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "DownloadPackage",
			"", downloadPackageResponse.BaseResponse.Message)
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = createAndSendForm(w, upload_package.UplFilesToMap(downloadPackageResponse.UplFiles), downloadPackageResponse.UploadDirsInfoJson)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "DownloadPackage",
			"", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
	}
}

// createAndSendForm Creates a form and sends it.
// The form contains all the information for downloading the package, namely:
// Directory names as field names (beginning with DIR$), and as values the files in that directory.
// The text field files_info, which contains json text describing the directory and file structure of the package.
func createAndSendForm(w http.ResponseWriter, uplFilesMap map[string][]string, uplDirsInfoJson []byte) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, filePaths := range uplFilesMap {
		for i := 0; i < len(filePaths); i++ {
			file, err := os.Open(filePaths[i])
			if err != nil {
				panic(err)
			}
			formFile, err := writer.CreateFormFile(key, filepath.Base(filePaths[i]))
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(formFile, file)
			if err != nil {
				panic(err)
			}

			err = file.Close()
			if err != nil {
				panic(err)
			}
		}
	}
	err := writer.WriteField("files_info", string(uplDirsInfoJson))
	if err != nil {
		return err
	}
	// Finishes recording the file. MUST be closed before submitting the form.
	err = writer.Close()
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", writer.FormDataContentType())
	_, err = w.Write(body.Bytes())
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	return nil
}
