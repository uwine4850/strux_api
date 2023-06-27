package users

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"reflect"
	"strux_api/internal/config"
	"strux_api/internal/rest_api/routes/errors"
	"strux_api/internal/rest_api/routes/utils"
	"strux_api/pkg/logging"
	"strux_api/services/protofiles/baseproto"
	"strux_api/services/protofiles/userproto"
)

// CreateUserService The function connects to the microservice responsible for operations with users and will send
// a request to create a new user.
// In this case, the answer is either various kinds of errors, or messages about the correct completion of the operation.
func CreateUserService(w http.ResponseWriter, r *http.Request) {
	connection, err := utils.CheckFormKeyAndGetUserServiceConnection(w, r, []string{"username", "password"})
	if err != nil {
		if reflect.DeepEqual(err, errors.ErrFormKeyNotExist{}) {
			logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "createUserService", "", err.Error())
		} else {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
		}
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
			utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)
	client := userproto.NewUserClient(connection)

	values, _, _ := utils.GetFormData(r)
	// send request and processing response
	request := &userproto.RequestCreateUser{
		Username: values["username"][0],
		Password: values["password"][0],
	}
	response, err := client.CreateUser(context.Background(), request)

	// catch errors and send response
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the user is not created
	if !response.Success {
		// if user already exist
		if response.Status != baseproto.ResponseStatus_StatusError {
			err := utils.CreateResponse(w, http.StatusOK, response)
			if err != nil {
				logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
				resp := &baseproto.BaseResponse{
					Message: err.Error(),
					Success: false,
					Status:  baseproto.ResponseStatus_StatusError,
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
			utils.SendResponseError(w, response.Message, http.StatusInternalServerError)
		}
		// if user successfully created
	} else {
		err := utils.CreateResponse(w, http.StatusCreated, response)
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "createUserService", "", err.Error())
			resp := &baseproto.BaseResponse{
				Message: err.Error(),
				Success: false,
				Status:  baseproto.ResponseStatus_StatusError,
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

// UserExistService sends a user existence request to the user service.
func UserExistService(w http.ResponseWriter, r *http.Request) {
	connection, err := utils.CheckFormKeyAndGetUserServiceConnection(w, r, []string{"username"})
	if err != nil {
		if reflect.DeepEqual(err, errors.ErrFormKeyNotExist{}) {
			logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "userExistService", "", err.Error())
		} else {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
		}
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(userServiceConnect *grpc.ClientConn) {
		err := userServiceConnect.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
			utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)
	client := userproto.NewUserClient(connection)

	values, _, _ := utils.GetFormData(r)
	checkExistUsername := values["username"][0]
	// send request
	response, err := client.UserExist(context.Background(), &userproto.RequestExistUser{Username: checkExistUsername})

	// catch errors and send response
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// processing error
	if response.Status == baseproto.ResponseStatus_StatusError {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", response.Message)
		utils.SendResponseError(w, response.Message, http.StatusInternalServerError)
		return
	}

	// No errors found, return user existence response.
	err = utils.CreateResponse(w, http.StatusOK, response)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
	}
}

func UserDeleteService(w http.ResponseWriter, r *http.Request) {
	// connect to user service and check form keys
	connection, err := utils.CheckFormKeyAndGetUserServiceConnection(w, r, []string{"username", "password"})
	if err != nil {
		if reflect.DeepEqual(err, errors.ErrFormKeyNotExist{}) {
			logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "userDeleteService", "", err.Error())
		} else {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userDeleteService", "", err.Error())
		}
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userDeleteService", "", err.Error())
			utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)

	// get form data
	values, _, _ := utils.GetFormData(r)

	// exec and processing UserDelete function
	client := userproto.NewUserClient(connection)
	response, err := client.UserDelete(context.Background(), &userproto.RequestDeleteUser{
		Username: values["username"][0],
		Password: values["password"][0],
	})
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userDeleteService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if response.Status == baseproto.ResponseStatus_StatusError {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userDeleteService", "", response.Message)
		utils.SendResponseError(w, response.Message, http.StatusInternalServerError)
		return
	}

	// No errors found, return user existence response.
	err = utils.CreateResponse(w, http.StatusOK, response)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userExistService", "", err.Error())
	}
}

func UserPasswordUpdateService(w http.ResponseWriter, r *http.Request) {
	// connect to user service and check form keys
	connection, err := utils.CheckFormKeyAndGetUserServiceConnection(w, r, []string{"username", "password", "newPassword"})
	if err != nil {
		if reflect.DeepEqual(err, errors.ErrFormKeyNotExist{}) {
			logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "userPasswordUpdateService", "", err.Error())
		} else {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userPasswordUpdateService", "", err.Error())
		}
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userPasswordUpdateService", "", err.Error())
			utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)

	// get form data
	values, _, _ := utils.GetFormData(r)

	client := userproto.NewUserClient(connection)
	request := &userproto.RequestUpdatePassword{
		Username:    values["username"][0],
		Password:    values["password"][0],
		NewPassword: values["newPassword"][0],
	}
	response, err := client.UserUpdatePassword(context.Background(), request)
	if response.Status == baseproto.ResponseStatus_StatusError {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userPasswordUpdateService", "", response.Message)
		utils.SendResponseError(w, response.Message, http.StatusInternalServerError)
		return
	}
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "userPasswordUpdateService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = utils.CreateResponse(w, http.StatusOK, response)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "userPasswordUpdateService", "", err.Error())
	}
}

func UserLogInService(w http.ResponseWriter, r *http.Request) {
	connection, err := utils.CheckFormKeyAndGetUserServiceConnection(w, r, []string{"username", "password"})
	if err != nil {
		if reflect.DeepEqual(err, errors.ErrFormKeyNotExist{}) {
			logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "UserLogInService", "", err.Error())
		} else {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "UserLogInService", "", err.Error())
		}
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "UserLogInService", "", err.Error())
			utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		}
	}(connection)

	values, _, _ := utils.GetFormData(r)

	client := userproto.NewUserClient(connection)
	request := &userproto.RequestUserLogIn{
		Username: values["username"][0],
		Password: values["password"][0],
	}
	response, err := client.UserLogIn(context.Background(), request)

	if response.Status == baseproto.ResponseStatus_StatusError {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "UserLogInService", "", response.Message)
		utils.SendResponseError(w, response.Message, http.StatusInternalServerError)
		return
	}
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.WarnLevel, "routes", "UserLogInService", "", err.Error())
		utils.SendResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = utils.CreateResponse(w, http.StatusOK, response)
	if err != nil {
		logging.CreateLog(config.APILogFileName, logrus.ErrorLevel, "routes", "UserLogInService", "", err.Error())
	}
}
