package internal

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"strux_api/internal/config"
	"strux_api/pkg/db"
	"strux_api/pkg/logging"
	"strux_api/services/user_service/protobufs"
)

// FindOneWitchResponse Searches for a single value in the database. Returns the corresponding response.
func FindOneWitchResponse(operation db.DatabaseOperation, colName string, value string, res interface{}) (*protobufs.BaseResponse, error) {
	err := operation.FindOneByValue(colName, value, res)
	if err != nil && err != mongo.ErrNoDocuments {
		resp := &protobufs.BaseResponse{
			Message: err.Error(),
			Success: false,
			Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusError},
		}
		return resp, err
	}
	resp := &protobufs.BaseResponse{
		Message: "OK",
		Success: true,
		Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusOk},
	}
	return resp, nil
}

// SendResponseError Returns an error response
func SendResponseError(errText string) *protobufs.BaseResponse {
	resp := &protobufs.BaseResponse{
		Message: errText,
		Status:  []protobufs.ResponseStatus{protobufs.ResponseStatus_StatusError},
		Success: false,
	}
	return resp
}

// GetDbClientConnection Returns the connection to the client
func GetDbClientConnection() (*mongo.Client, context.Context, *protobufs.BaseResponse) {
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
