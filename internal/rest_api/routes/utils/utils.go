package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/uwine4850/strux_api/internal/config"
	"github.com/uwine4850/strux_api/internal/rest_api/routes/errors"
	"github.com/uwine4850/strux_api/pkg/logging"
	"github.com/uwine4850/strux_api/services/protofiles/baseproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"mime/multipart"
	"net/http"
)

// CreateResponse Creates and sends a response to the client
func CreateResponse(w http.ResponseWriter, httpStatus int, response *baseproto.BaseResponse) error {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	jm := protojson.MarshalOptions{EmitUnpopulated: true, UseEnumNumbers: true}
	jsonResp, err := jm.Marshal(response)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		return err
	}
	return nil
}

// GetFormData Returns form data. Values and Files are grouped.
func GetFormData(r *http.Request) (map[string][]string, map[string][]*multipart.FileHeader, error) {
	err := r.ParseMultipartForm(64 << 20)
	return r.MultipartForm.Value, r.MultipartForm.File, err
}

// SendResponseError Returns an error response
func SendResponseError(w http.ResponseWriter, errText string, httpStatus int) {
	resp := &baseproto.BaseResponse{
		Message: errText,
		Success: false,
		Status:  baseproto.ResponseStatus_StatusError,
	}
	err := CreateResponse(w, http.StatusInternalServerError, resp)
	if err != nil {
		resp.Message = err.Error()
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
	}
}

// CheckPostKeysExist Checks if the required keys are in the form.
func CheckPostKeysExist(values map[string][]string, keys []string) (string, bool) {
	for i := 0; i < len(keys); i++ {
		exist := false
		for s, _ := range values {
			if s == keys[i] {
				exist = true
			}
		}
		if !exist {
			return keys[i], false
		}
	}
	return "", true
}

func CheckPostFilesKeysExist(values map[string][]*multipart.FileHeader, keys []string) (string, bool) {
	for i := 0; i < len(keys); i++ {
		exist := false
		for s, _ := range values {
			if s == keys[i] {
				exist = true
			}
		}
		if !exist {
			return keys[i], false
		}
	}
	return "", true
}

// ConnectToUserService Connecting to the user microservice.
func ConnectToUserService() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(config.GetUserServiceAddress(), opts...)
	return conn, err
}

// ConnectToPackageService Connecting to the user microservice.
func ConnectToPackageService() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(config.GetPkgServiceAddress(), opts...)
	return conn, err
}

// CheckFormKeyAndGetUserServiceConnection a frequently used template is placed in a separate function.
// Checks if the key is in the form, and also resets the connection to the grpc service.
func CheckFormKeyAndGetUserServiceConnection(w http.ResponseWriter, r *http.Request, keys []string) (*grpc.ClientConn, error) {
	values, _, err := GetFormData(r)
	if err != nil {
		return nil, err
	}
	key, ok := CheckPostKeysExist(values, keys)
	if !ok {
		return nil, &errors.ErrFormKeyNotExist{KeyName: key}
	}

	// connect to user service
	connection, err := ConnectToUserService()
	if err != nil {
		return nil, err
	}
	return connection, nil
}

// CheckFormKeyAndGetPackageServiceConnection a frequently used template is placed in a separate function.
// Checks if the key is in the form, and also resets the connection to the grpc service.
func CheckFormKeyAndGetPackageServiceConnection(w http.ResponseWriter, r *http.Request, filesKeys []string, valuesKeys []string) (*grpc.ClientConn, error) {
	values, files, err := GetFormData(r)
	if err != nil {
		return nil, err
	}
	key, ok := CheckPostFilesKeysExist(files, filesKeys)
	if !ok {
		return nil, &errors.ErrFormKeyNotExist{KeyName: key}
	}

	key, ok = CheckPostKeysExist(values, valuesKeys)
	if !ok {
		return nil, &errors.ErrFormKeyNotExist{KeyName: key}
	}

	// connect to user service
	connection, err := ConnectToPackageService()
	if err != nil {
		return nil, err
	}
	return connection, nil
}
