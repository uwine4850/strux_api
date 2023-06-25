package internal

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"strux_api/internal/config"
	"strux_api/internal/config/schema"
	"strux_api/pkg/logging"
	"strux_api/services/user_service/protobufs"
)

func PasswordUpdate(username string, password string, newPassword string) *protobufs.BaseResponse {
	if password == newPassword {
		resp := &protobufs.BaseResponse{
			Message: "The old and new passwords are the same.",
			Success: false,
			Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusOk},
		}
		return resp
	}

	// connect to database
	clientConnection, ctx, errResponse := GetDbClientConnection()
	if errResponse != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "PasswordUpdate", "", errResponse.Message)
		return errResponse
	}
	defer func(clientConnection *mongo.Client, ctx context.Context) {
		err := clientConnection.Disconnect(ctx)
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "PasswordUpdate", "", err.Error())
			panic(err)
		}
	}(clientConnection, ctx)

	hashPassword, errResponse := HashPassword(newPassword)
	if errResponse != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "PasswordUpdate", "", errResponse.Message)
		return errResponse
	}

	// get user db operations
	operation := GetUserOperation(clientConnection, ctx)

	// match passwords
	var user schema.User
	err := operation.FindOneByValue("username", username, &user)
	if err != nil && err != mongo.ErrNoDocuments {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "PasswordUpdate", "", err.Error())
		return SendResponseError(err.Error())
	}
	if user.Username != "" {
		// check password match
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			resp := &protobufs.BaseResponse{
				Message: "Old password mismatch.",
				Success: false,
				Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusOk},
			}
			return resp
		}
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "PasswordUpdate", "", err.Error())
			return SendResponseError(err.Error())
		}
	} else {
		// user not exist
		resp := &protobufs.BaseResponse{
			Message: fmt.Sprintf("User %s not exist.", username),
			Success: false,
			Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusOk},
		}
		return resp
	}

	// update
	updUser := schema.User{
		Username: username,
		Password: string(hashPassword),
	}
	upd := bson.M{"$set": updUser}
	res := operation.FindOneAndUpdate("username", username, &upd)
	err = res.Err()
	if err != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "PasswordUpdate", "", err.Error())
		return SendResponseError(err.Error())
	}
	resp := &protobufs.BaseResponse{
		Message: "Password updated.",
		Success: true,
		Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusOk},
	}
	return resp
}
