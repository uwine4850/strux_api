package internal

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"strux_api/internal/config"
	"strux_api/internal/config/schema"
	"strux_api/pkg/db"
	"strux_api/pkg/logging"
	"strux_api/services/protofiles/baseproto"
	"strux_api/services/utils"
)

// CreateUser Adding a new user to the database if it was not previously found
func CreateUser(username string, password string) *baseproto.BaseResponse {
	// connect to database
	clientConnection, ctx, errResponse := utils.GetDbClientConnection()
	if errResponse != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "CreateUser", "", errResponse.Message)
		return errResponse
	}
	defer func(clientConnect *mongo.Client, ctx context.Context) {
		err := clientConnect.Disconnect(ctx)
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.PanicLevel, "user_service.internal", "CreateUser", "", err.Error())
			panic(err)
		}
	}(clientConnection, ctx)

	// get user
	operation := GetUserOperation(clientConnection, ctx)
	user := schema.User{}
	resp, err := utils.FindOneWitchResponse(operation, "username", username, &user)
	if err != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "CreateUser", "", err.Error())
		return utils.SendResponseError(resp.Message)
	}

	// add user
	if user.Username != "" {
		msg := fmt.Sprintf("User %s already exist.", user.Username)
		resp := &baseproto.BaseResponse{
			Message: msg,
			Success: false,
			Status:  baseproto.ResponseStatus_StatusOk,
		}
		return resp
	} else {
		// hashing password
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 4)
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "CreateUser", "", err.Error())
			return utils.SendResponseError(err.Error())
		}

		_, err = operation.InsertOne(
			schema.User{
				Username: username,
				Password: string(passwordHash),
			})
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "CreateUser", "", err.Error())
			return utils.SendResponseError(err.Error())
		} else {
			resp := &baseproto.BaseResponse{
				Message: fmt.Sprintf("User %s created successfuly.", username),
				Status:  baseproto.ResponseStatus_StatusOk,
				Success: true,
			}
			return resp
		}
	}
}

// GetUserOperation returns a class for database operations
func GetUserOperation(clientConnect *mongo.Client, ctx context.Context) db.DatabaseOperation {
	operation := db.DatabaseOperation{
		DbName:         config.Dbname,
		CollectionName: config.UsersCN,
		Client:         clientConnect,
		Ctx:            ctx,
	}
	return operation
}
