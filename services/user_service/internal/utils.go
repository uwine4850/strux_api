package internal

import (
	"go.mongodb.org/mongo-driver/mongo"
	"strux_api/pkg/db"
	"strux_api/services/user_service/protobufs"
)

type Response struct {
	Message string
	Success bool
}

// FindOneWitchResponse Searches for a single value in the database. Returns the corresponding response.
func FindOneWitchResponse(operation db.DatabaseOperation, colName string, value string, res interface{}) (Response, error) {
	err := operation.FindOneByValue(colName, value, res)
	if err != nil && err != mongo.ErrNoDocuments {
		resp := Response{
			Message: err.Error(),
			Success: false,
		}
		return resp, err
	}
	resp := Response{
		Message: "OK",
		Success: true,
	}
	return resp, nil
}

// SendResponseError Returns an error response
func SendResponseError(errText string) *protobufs.BaseResponse {
	resp := &protobufs.BaseResponse{
		Message: errText,
		Success: false,
	}
	return resp
}
