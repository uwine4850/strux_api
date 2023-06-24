package routes

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"mime/multipart"
	"net/http"
	"strux_api/internal/config"
	"strux_api/pkg/logging"
	"strux_api/services/user_service/protobufs"
)

// CreateResponse Creates and sends a response to the client
func CreateResponse(w http.ResponseWriter, httpStatus int, response *protobufs.BaseResponse) error {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	jm := protojson.MarshalOptions{EmitUnpopulated: true}
	jsonResp, err := jm.Marshal(response)
	//jsonResp, err := json.Marshal(response)
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
	resp := &protobufs.BaseResponse{
		Message: errText,
		Success: false,
		Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusError},
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

// ConnectToUserService Connecting to the user microservice.
func ConnectToUserService() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(config.UserServiceAddress, opts...)
	return conn, err
}
