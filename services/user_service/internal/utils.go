package internal

import (
	"go.mongodb.org/mongo-driver/mongo"
	"strux_api/pkg/db"
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
