package internal

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"strux_api/internal/config"
	"strux_api/internal/config/schema"
	"strux_api/pkg/db"
	"strux_api/pkg/logging"
	"strux_api/services/user_service/protobufs"
)

// UserExist checks if the selected username exists in the database.
// If no errors occur during the check (see status), success: true means that the user was found, and false is not found.
func UserExist(username string) *protobufs.BaseResponse {
	// connect to database
	client, err := db.GetMongoClient()
	if err != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserExist", "", err.Error())
		return SendResponseError(err.Error())
	}
	clientConnect, ctx, err := client.Connect()
	if err != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserExist", "", err.Error())
		return SendResponseError(err.Error())
	}
	defer func(clientConnect *mongo.Client, ctx context.Context) {
		err := clientConnect.Disconnect(ctx)
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserExist", "", err.Error())
		}
	}(clientConnect, ctx)

	// get user operation
	operation := GetUserOperation(clientConnect, ctx)
	var user schema.User
	err = operation.FindOneByValue("username", username, &user)
	if err != nil && err != mongo.ErrNoDocuments {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserExist", "", err.Error())
		return SendResponseError(err.Error())
	}
	// user exist
	if user.Username != "" {
		resp := &protobufs.BaseResponse{
			Message: "Exist",
			Success: true,
			Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusOk},
		}
		return resp
		//	user not exist
	} else {
		resp := &protobufs.BaseResponse{
			Message: "Not exist",
			Success: false,
			Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusOk},
		}
		return resp
	}
}
