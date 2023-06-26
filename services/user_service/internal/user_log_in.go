package internal

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"strux_api/internal/config"
	"strux_api/internal/config/schema"
	"strux_api/pkg/logging"
	"strux_api/services/user_service/protobufs"
)

func UserLogIn(username string, password string) *protobufs.BaseResponse {
	// connect to database
	clientConnection, ctx, errResponse := GetDbClientConnection()
	if errResponse != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserLogIn", "", errResponse.Message)
		return errResponse
	}
	defer func(clientConnection *mongo.Client, ctx context.Context) {
		err := clientConnection.Disconnect(ctx)
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserLogIn", "", err.Error())
			panic(err)
		}
	}(clientConnection, ctx)
	// get user db operations
	operation := GetUserOperation(clientConnection, ctx)
	var user schema.User
	err := operation.FindOneByValue("username", username, &user)
	if err != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserLogIn", "", err.Error())
		SendResponseError(err.Error())
	}

	if user.Username != "" {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			resp := &protobufs.BaseResponse{
				Message: "Password mismatch.",
				Success: false,
				Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusOk},
			}
			return resp
		}
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "PasswordUpdate", "", err.Error())
			return SendResponseError(err.Error())
		}
		resp := &protobufs.BaseResponse{
			Message: "OK",
			Success: true,
			Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusOk},
		}
		return resp
	} else {
		// user not exist
		resp := &protobufs.BaseResponse{
			Message: fmt.Sprintf("User %s not exist.", username),
			Success: false,
			Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusOk},
		}
		return resp
	}
}
