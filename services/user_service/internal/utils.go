package internal

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"strux_api/internal/config"
	"strux_api/pkg/db"
	"strux_api/pkg/logging"
	"strux_api/services/protofiles/baseproto"
)

// FindOneWitchResponse Searches for a single value in the database. Returns the corresponding response.
func FindOneWitchResponse(operation db.DatabaseOperation, colName string, value string, res interface{}) (*baseproto.BaseResponse, error) {
	err := operation.FindOneByValue(colName, value, res)
	if err != nil && err != mongo.ErrNoDocuments {
		resp := &baseproto.BaseResponse{
			Message: err.Error(),
			Success: false,
			Status:  baseproto.ResponseStatus_StatusError,
		}
		return resp, err
	}
	resp := &baseproto.BaseResponse{
		Message: "OK",
		Success: true,
		Status:  baseproto.ResponseStatus_StatusOk,
	}
	return resp, nil
}

// SendResponseError Returns an error response
func SendResponseError(errText string) *baseproto.BaseResponse {
	resp := &baseproto.BaseResponse{
		Message: errText,
		Status:  baseproto.ResponseStatus_StatusError,
		Success: false,
	}
	return resp
}

// GetDbClientConnection Returns the connection to the client
func GetDbClientConnection() (*mongo.Client, context.Context, *baseproto.BaseResponse) {
	client, err := db.GetMongoClient()
	if err != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "CreateUser", "", err.Error())
		return nil, nil, SendResponseError(err.Error())
	}
	clientConnect, ctx, err := client.Connect()
	if err != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "CreateUser", "", err.Error())
		return nil, nil, SendResponseError(err.Error())
	}
	return clientConnect, ctx, nil
}

func HashPassword(password string) ([]byte, *baseproto.BaseResponse) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "CreateUser", "", err.Error())
		return nil, SendResponseError(err.Error())
	}
	return passwordHash, nil
}
