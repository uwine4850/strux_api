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
	"strux_api/services/protofiles/baseproto"
	"strux_api/services/utils"
)

func UserLogIn(username string, password string) *baseproto.BaseResponse {
	// connect to database
	clientConnection, ctx, errResponse := utils.GetDbClientConnection()
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
	if err != nil && err != mongo.ErrNoDocuments {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserLogIn", "", err.Error())
		return utils.SendResponseError(err.Error())
	}

	if user.Username != "" {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			resp := &baseproto.BaseResponse{
				Message: "Password mismatch.",
				Success: false,
				Status:  baseproto.ResponseStatus_StatusOk,
			}
			return resp
		}
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "PasswordUpdate", "", err.Error())
			return utils.SendResponseError(err.Error())
		}
		resp := &baseproto.BaseResponse{
			Message: "OK",
			Success: true,
			Status:  baseproto.ResponseStatus_StatusOk,
		}
		return resp
	} else {
		// user not exist
		resp := &baseproto.BaseResponse{
			Message: fmt.Sprintf("User %s not exist.", username),
			Success: false,
			Status:  baseproto.ResponseStatus_StatusOk,
		}
		return resp
	}
}
