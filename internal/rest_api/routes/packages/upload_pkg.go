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

func UploadPkgService(w http.ResponseWriter, r *http.Request) {
	connection, err := utils.CheckFormKeyAndGetPackageServiceConnection(w, r, []string{"files_data", "files_info"},
		[]string{"username", "password", "version"})

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
	var uplFilesData []*pkgproto.UploadFile
	err = setUploadFiles(files["files_data"], &uplFilesData)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "UploadPkgService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// parsing files_info and setting dirs info
	f := files["files_info"][0]
	uplDirInfo := &pkgproto.UploadDirInfo{}
	err = setUploadDirInfo(f, uplDirInfo)
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
		UplFiles:   uplFilesData,
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

// setUploadDirInfo decodes the project directory tree into a select structure.
func setUploadDirInfo(readFile *multipart.FileHeader, uplDirInfo *pkgproto.UploadDirInfo) error {
	file, err := readFile.Open()
	if err != nil {
		return err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "packages", "setUploadDirInfo", "", err.Error())
			panic(err)
		}
	}(file)
	err = json.NewDecoder(file).Decode(uplDirInfo)
	if err != nil {
		return err
	}
	return nil
}

// setUploadFiles populates the appropriate structure with file and byte data.
func setUploadFiles(filesData []*multipart.FileHeader, uploadFiles *[]*pkgproto.UploadFile) error {
	for i := 0; i < len(filesData); i++ {
		f := filesData[i]
		fOpen, err := f.Open()
		if err != nil {
			return err
		}
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(fOpen)
		if err != nil {
			return err
		}
		if buf == nil {
			continue
		}
		u := pkgproto.UploadFile{
			FileName:      f.Filename,
			FileBytesData: buf.Bytes(),
		}
		*uploadFiles = append(*uploadFiles, &u)
		err = fOpen.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
