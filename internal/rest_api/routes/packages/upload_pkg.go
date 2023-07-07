package packages

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"strux_api/internal/config"
	"strux_api/internal/rest_api/routes/errors"
	"strux_api/internal/rest_api/routes/utils"
	"strux_api/pkg/logging"
	"strux_api/services/protofiles/pkgproto"
)

func UploadPkgService(w http.ResponseWriter, r *http.Request) {
	connection, err := utils.CheckFormKeyAndGetPackageServiceConnection(w, r, []string{"files_data"},
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
	err = setUploadFiles(files, &uplFilesData1)
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

// setUploadFiles populates the appropriate structure with file and byte data.
func setUploadFiles(filesData map[string][]*multipart.FileHeader, uploadFiles *[]*pkgproto.UploadFile) error {
	for names, value := range filesData {
		dirName := strings.Split(names, "DIR$")
		if len(dirName) > 2 {
			return &errors.ErrInvalidFilePath{Path: names}
		}
		if len(dirName) < 2 {
			continue
		}
		for j := 0; j < len(value); j++ {
			open, err := value[j].Open()
			if err != nil {
				return err
			}
			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(open)
			if err != nil {
				return err
			}
			if buf == nil {
				continue
			}
			*uploadFiles = append(*uploadFiles, &pkgproto.UploadFile{
				FileName:      filepath.Join(dirName[1], value[j].Filename),
				FileBytesData: buf.Bytes(),
			})
		}
	}
	return nil
}
