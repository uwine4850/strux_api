package internal

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/uwine4850/strux_api/internal/config"
	"github.com/uwine4850/strux_api/internal/config/schema"
	"github.com/uwine4850/strux_api/pkg/logging"
	"github.com/uwine4850/strux_api/services/protofiles/baseproto"
	"github.com/uwine4850/strux_api/services/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
)

func UserDelete(username string, password string) *baseproto.BaseResponse {
	// connect to database
	clientConnection, ctx, errResponse := utils.GetDbClientConnection()
	if errResponse != nil {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserDelete", "", errResponse.Message)
		return errResponse
	}
	defer func(clientConnection *mongo.Client, ctx context.Context) {
		err := clientConnection.Disconnect(ctx)
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserDelete", "", err.Error())
			panic(err)
		}
	}(clientConnection, ctx)

	// get user db operations
	operation := GetUserOperation(clientConnection, ctx)
	var user schema.User
	err := operation.FindOneByValue("username", username, &user)
	if err != nil && err != mongo.ErrNoDocuments {
		logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserDelete", "", err.Error())
		return utils.SendResponseError(err.Error())
	}

	if user.Username != "" {
		// check password match
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
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserDelete", "", err.Error())
			return utils.SendResponseError(err.Error())
		}
		var user map[string]string
		err = operation.FindOneByValue("username", username, &user)
		if err != nil {
			panic(err)
		}

		// delete packages
		deletedPackages, err := deleteDbPackages(bson.D{
			{"parentUserId", user["_id"]},
		}, clientConnection, ctx)
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserDelete", "", err.Error())
			return utils.SendResponseError(err.Error())
		}

		// delete package versions
		err = deleteDbPackageVersions(deletedPackages, clientConnection, ctx)
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserDelete", "", err.Error())
			return utils.SendResponseError(err.Error())
		}

		// delete user package directory
		err = deleteUserDir(username)
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserDelete", "", err.Error())
			return utils.SendResponseError(err.Error())
		}

		// delete user
		_, err = operation.DeleteOneEntry(bson.D{{"username", username}})
		if err != nil {
			logging.CreateLog(config.UserServiceLogFileName, logrus.ErrorLevel, "user_service.internal", "UserDelete", "", err.Error())
			return utils.SendResponseError(err.Error())
		} else {
			resp := &baseproto.BaseResponse{
				Message: fmt.Sprintf("User %s deleted.", username),
				Success: true,
				Status:  baseproto.ResponseStatus_StatusOk,
			}
			return resp
		}
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

// deleteDbPackages deleting all of the user's packages from the database.
// All data about deleted records is returned.
func deleteDbPackages(filter bson.D, client *mongo.Client, ctx context.Context) ([]map[string]string, error) {
	operation := utils.GetPackageServiceOperation(client, ctx)
	var userPackages []map[string]string
	err := operation.GetMultipleValues(filter, &userPackages)
	if err != nil {
		return nil, err
	}
	err = operation.DeleteManyDocuments(filter)
	if err != nil {
		return nil, err
	}
	return userPackages, nil
}

// deleteDbPackageVersions remove package versions of all deleted packages.
func deleteDbPackageVersions(userPackages []map[string]string, client *mongo.Client, ctx context.Context) error {
	operation := utils.GetPackageVersionServiceOperation(client, ctx)
	for i := 0; i < len(userPackages); i++ {
		userPackagesId := userPackages[i]["_id"]
		err := operation.DeleteManyDocuments(bson.D{{"parentPackageId", userPackagesId}})
		if err != nil {
			return err
		}
	}
	return nil
}

// deleteUserDir removing the user directory and, accordingly, all packages.
func deleteUserDir(username string) error {
	err := os.RemoveAll(filepath.Join("/home/uwine/Documents/GolandProjects/strux_api/services/package_service/UserPackages", username))
	if err != nil {
		return err
	}
	return nil
}
