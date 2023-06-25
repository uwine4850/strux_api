package routes

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"reflect"
	"strux_api/internal/config"
	"strux_api/pkg/logging"
	"strux_api/services/user_service/protobufs"
)

func UsersInit() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/create-user/", createUserService)
	r.Get("/user-exist/", userExistService)
	r.Delete("/user-delete/", userDeleteService)
	r.Put("/user-password-update/", userPasswordUpdateService)
	return r
}

// createUserService The function connects to the microservice responsible for operations with users and will send
// a request to create a new user.
// In this case, the answer is either various kinds of errors, or messages about the correct completion of the operation.
func createUserService(w http.ResponseWriter, r *http.Request) {
	connection, err := CheckFormKeyAndGetUserServiceConnection(w, r, []string{"username", "password"})
	if err != nil {
		if reflect.DeepEqual(err, ErrFormKeyNotExist{}) {
			logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "createUserService", "", err.Error())
		} else {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
		}
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
			SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)
	client := protobufs.NewUserClient(connection)

	values, _, _ := GetFormData(r)
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
	connection, err := CheckFormKeyAndGetUserServiceConnection(w, r, []string{"username"})
	if err != nil {
		if reflect.DeepEqual(err, ErrFormKeyNotExist{}) {
			logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "userExistService", "", err.Error())
		} else {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
		}
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(userServiceConnect *grpc.ClientConn) {
		err := userServiceConnect.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
			SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)
	client := protobufs.NewUserClient(connection)

	values, _, _ := GetFormData(r)
	checkExistUsername := values["username"][0]
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

func userDeleteService(w http.ResponseWriter, r *http.Request) {
	// connect to user service and check form keys
	connection, err := CheckFormKeyAndGetUserServiceConnection(w, r, []string{"username", "password"})
	if err != nil {
		if reflect.DeepEqual(err, ErrFormKeyNotExist{}) {
			logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "userDeleteService", "", err.Error())
		} else {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userDeleteService", "", err.Error())
		}
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userDeleteService", "", err.Error())
			SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)

	// get form data
	values, _, _ := GetFormData(r)

	// exec and processing UserDelete function
	client := protobufs.NewUserClient(connection)
	response, err := client.UserDelete(context.Background(), &protobufs.RequestDeleteUser{
		Username: values["username"][0],
		Password: values["password"][0],
	})
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userDeleteService", "", err.Error())
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if response.Status[0] == protobufs.ResponseStatus_StatusError {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userDeleteService", "", response.Message)
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// No errors found, return user existence response.
	err = CreateResponse(w, http.StatusOK, response)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
	}
}

func userPasswordUpdateService(w http.ResponseWriter, r *http.Request) {
	// connect to user service and check form keys
	connection, err := CheckFormKeyAndGetUserServiceConnection(w, r, []string{"username", "password", "newPassword"})
	if err != nil {
		if reflect.DeepEqual(err, ErrFormKeyNotExist{}) {
			logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "userPasswordUpdateService", "", err.Error())
		} else {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userPasswordUpdateService", "", err.Error())
		}
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userPasswordUpdateService", "", err.Error())
			SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)

	// get form data
	values, _, _ := GetFormData(r)

	client := protobufs.NewUserClient(connection)
	request := &protobufs.RequestUpdatePassword{
		Username:    values["username"][0],
		Password:    values["password"][0],
		NewPassword: values["newPassword"][0],
	}
	response, err := client.UserUpdatePassword(context.Background(), request)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "userPasswordUpdateService", "", err.Error())
		SendResponseError(w, err.Error(), http.StatusInternalServerError)
	}
	err = CreateResponse(w, http.StatusOK, response)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userPasswordUpdateService", "", err.Error())
	}
}
