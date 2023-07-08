package internal

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/uwine4850/strux_api/internal/config"
	"github.com/uwine4850/strux_api/internal/config/schema"
	"github.com/uwine4850/strux_api/pkg/logging"
	"github.com/uwine4850/strux_api/services/protofiles/baseproto"
	"github.com/uwine4850/strux_api/services/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserExist checks if the selected username exists in the database.
// If no errors occur during the check (see status), success: true means that the user was found, and false is not found.
func UserExist(username string) *baseproto.BaseResponse {
	// connect to database
	clientConnection, ctx, errResponse := utils.GetDbClientConnection()
	if errResponse != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserExist", "", errResponse.Message)
		return errResponse
	}
	defer func(clientConnect *mongo.Client, ctx context.Context) {
		err := clientConnect.Disconnect(ctx)
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserExist", "", err.Error())
		}
	}(clientConnection, ctx)

	// get user operation
	operation := GetUserOperation(clientConnection, ctx)
	var user schema.User
	err := operation.FindOneByValue("username", username, &user)
	if err != nil && err != mongo.ErrNoDocuments {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserExist", "", err.Error())
		return utils.SendResponseError(err.Error())
	}
	// user exist
	if user.Username != "" {
		resp := &baseproto.BaseResponse{
			Message: "Exist",
			Success: true,
			Status:  baseproto.ResponseStatus_StatusOk,
		}
		return resp
		//	user not exist
	} else {
		resp := &baseproto.BaseResponse{
			Message: "Not exist",
			Success: false,
			Status:  baseproto.ResponseStatus_StatusOk,
		}
		return resp
	}
}
