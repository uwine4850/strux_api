package internal

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strux_api/internal/config"
	"strux_api/pkg/db"
	"strux_api/pkg/logging"
	"strux_api/services/protofiles/baseproto"
	"strux_api/services/protofiles/pkgproto"
	"strux_api/services/utils"
)

func ExistsPackage(requestData *pkgproto.RequestPackageExists) *baseproto.BaseResponse {
	client, err := db.GetMongoClient()
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "ExistsPackage", "", err.Error())
		return utils.SendResponseError(err.Error())
	}
	clientConnection, ctx, err := client.Connect()
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "ExistsPackage", "", err.Error())
		return utils.SendResponseError(err.Error())
	}
	defer func(clientConnection *mongo.Client, ctx context.Context) {
		err := clientConnection.Disconnect(ctx)
		if err != nil {
			logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "ExistsPackage", "", err.Error())
		}
	}(clientConnection, ctx)

	userOperation := utils.GetUserServiceOperation(clientConnection, ctx)
	packageOperation := utils.GetPackageServiceOperation(clientConnection, ctx)
	pkgVersionOperation := utils.GetPackageVersionServiceOperation(clientConnection, ctx)

	// checking the presence of a user
	var user map[string]string
	err = userOperation.FindOneByValue("username", requestData.Username, &user)
	if err == mongo.ErrNoDocuments {
		return utils.SendResponseWarning(fmt.Sprintf("User %s not exist.", requestData.Username))
	}
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "ExistsPackage", "", err.Error())
		return utils.SendResponseError(err.Error())
	}

	// checking the presence of a package
	var userPackage map[string]string
	err = packageOperation.FindOneByMultipleValues(bson.D{
		{"parentUserId", user["_id"]},
		{"packageName", requestData.PkgName},
	}, &userPackage)
	if err == mongo.ErrNoDocuments {
		return utils.SendResponseWarning(fmt.Sprintf("Package %s not exist.", requestData.PkgName))
	}
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "ExistsPackage", "", err.Error())
		return utils.SendResponseError(err.Error())
	}

	// checking the presence of a package version
	var packageVersion map[string]string
	err = pkgVersionOperation.FindOneByMultipleValues(bson.D{
		{"parentPackageId", userPackage["_id"]},
		{"version", requestData.Version},
	}, &packageVersion)
	if err == mongo.ErrNoDocuments {
		return utils.SendResponseWarning(fmt.Sprintf("Package version %s not exist.", requestData.Version))
	}
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "ExistsPackage", "", err.Error())
		return utils.SendResponseError(err.Error())
	}

	resp := &baseproto.BaseResponse{
		Message: fmt.Sprintf("Package %s successfully found", requestData.PkgName+"/"+requestData.Version),
		Success: true,
		Status:  baseproto.ResponseStatus_StatusOk,
	}
	return resp
}
