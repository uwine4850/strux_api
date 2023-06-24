package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"strux_api/internal/config"
	"strux_api/pkg/logging"
	"strux_api/services/user_service/protobufs"
)

func UsersInit() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/create-user/", createUserService)
	r.Get("/user-exist/", userExistService)
	return r
}

// createUserService The function connects to the microservice responsible for operations with users and will send
// a request to create a new user.
// In this case, the answer is either various kinds of errors, or messages about the correct completion of the operation.
func createUserService(w http.ResponseWriter, r *http.Request) {
	// check post keys exist
	values, _, err := GetFormData(r)
	if err != nil {
		panic(err)
	}
	key, ok := CheckPostKeysExist(values, []string{"username", "password"})
	if !ok {
		msg := fmt.Sprintf("Post key '%s' not exist", key)
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", msg)
		SendResponseError(w, msg, http.StatusInternalServerError)
		return
	}

	// connect to user service
	connection, err := ConnectToUserService()
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
			SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)
	client := protobufs.NewUserClient(connection)

	// send request and processing response
	request := &protobufs.RequestCreateUser{
		Username: values["username"][0],
		Password: values["password"][0],
	}
	response, err := client.CreateUser(context.Background(), request)

	// catch errors and send response
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the user is not created
	if !response.Success {
		// if user already exist
		if response.Status[0] != protobufs.ResponseStatus_StatusError {
			err := CreateResponse(w, http.StatusOK, response)
			if err != nil {
				logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
				resp := &protobufs.BaseResponse{
					Message: err.Error(),
					Success: false,
					Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusError},
				}
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set("Content-Type", "application/json")
				jsonResp, err := json.Marshal(resp)
				if err != nil {
					panic(err)
				}
				_, err = w.Write(jsonResp)
				if err != nil {
					panic(err)
				}
			}
			return
			//	if some error
		} else {
			SendResponseError(w, response.Message, http.StatusInternalServerError)
		}
		// if user successfully created
	} else {
		err := CreateResponse(w, http.StatusCreated, response)
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
			resp := &protobufs.BaseResponse{
				Message: err.Error(),
				Success: false,
				Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusError},
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				panic(err)
			}
			_, err = w.Write(jsonResp)
			if err != nil {
				panic(err)
			}
		}
	}
}

// userExistService sends a user existence request to the user service.
func userExistService(w http.ResponseWriter, r *http.Request) {
	// check valid form keys
	values, _, _ := GetFormData(r)
	key, ok := CheckPostKeysExist(values, []string{"username"})
	if !ok {
		msg := fmt.Sprintf("Post key '%s' not exist", key)
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", msg)
		SendResponseError(w, msg, http.StatusInternalServerError)
		return
	}
	checkExistUsername := values["username"][0]

	// connect to user service
	userServiceConnect, err := ConnectToUserService()
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
	}
	defer func(userServiceConnect *grpc.ClientConn) {
		err := userServiceConnect.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
			SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(userServiceConnect)
	client := protobufs.NewUserClient(userServiceConnect)

	// send request
	response, err := client.UserExist(context.Background(), &protobufs.RequestExistUser{Username: checkExistUsername})

	// catch errors and send response
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// processing error
	if response.Status[0] == protobufs.ResponseStatus_StatusError {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", response.Message)
		SendResponseError(w, response.Message, http.StatusInternalServerError)
		return
	}

	// No errors found, return user existence response.
	err = CreateResponse(w, http.StatusOK, response)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
	}
}
