package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strux_api/internal/config"
	"strux_api/internal/config/schema"
	"strux_api/pkg/db"
	"strux_api/pkg/logging"
	"strux_api/services/protofiles/baseproto"
	"strux_api/services/protofiles/pkgproto"
	"strux_api/services/utils"
)

func ShowVersions(requestData *pkgproto.RequestShowVersions) *pkgproto.MutateShowVersionBaseResponse {
	client, err := db.GetMongoClient()
	if err != nil {
		panic(err)
	}
	connect, ctx, err := client.Connect()
	if err != nil {
		panic(err)
	}
	// get usr id
	userId, err := getUserId(connect, ctx, requestData)
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "ShowVersions", "", err.Error())
		return sendMutateShowVersionsBaseResponseError(err.Error())
	}

	//get package id
	packageId, err := getPackageId(connect, ctx, requestData, userId)
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "ShowVersions", "", err.Error())
		return sendMutateShowVersionsBaseResponseError(err.Error())
	}

	// get versions
	packageVersionOperation := utils.GetPackageVersionServiceOperation(connect, ctx)
	var packageVersions []schema.PackageVersion
	err = packageVersionOperation.GetMultipleValues(bson.D{
		{"parentPackageId", packageId},
	}, &packageVersions)
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "ShowVersions", "", err.Error())
		return sendMutateShowVersionsBaseResponseError(err.Error())
	}
	var versions []string
	for i := 0; i < len(packageVersions); i++ {
		versions = append(versions, packageVersions[i].Version)
	}
	return &pkgproto.MutateShowVersionBaseResponse{
		BaseResponse: &baseproto.BaseResponse{
			Message: "OK",
			Success: true,
			Status:  baseproto.ResponseStatus_StatusOk,
		},
		Versions: versions,
	}
}

func sendMutateShowVersionsBaseResponseError(msg string) *pkgproto.MutateShowVersionBaseResponse {
	return &pkgproto.MutateShowVersionBaseResponse{
		BaseResponse: &baseproto.BaseResponse{
			Message: msg,
			Success: false,
			Status:  baseproto.ResponseStatus_StatusError,
		},
		Versions: nil,
	}
}

func getPackageId(clientConnection *mongo.Client, ctx context.Context, requestData *pkgproto.RequestShowVersions, userId string) (string, error) {
	packageOperation := utils.GetPackageServiceOperation(clientConnection, ctx)
	var userPackage map[string]string
	err := packageOperation.FindOneByMultipleValues(bson.D{
		{"parentUserId", userId},
		{"packageName", requestData.PkgName},
	}, &userPackage)
	if err == mongo.ErrNoDocuments {
		return "", errors.New(fmt.Sprintf("Package %s/%s not exist", requestData.Username, requestData.PkgName))
	}
	if err != nil {
		return "", err
	}
	return userPackage["_id"], nil
}

func getUserId(clientConnection *mongo.Client, ctx context.Context, requestData *pkgproto.RequestShowVersions) (string, error) {
	userOperation := utils.GetUserServiceOperation(clientConnection, ctx)
	var user map[string]string
	err := userOperation.FindOneByValue("username", requestData.Username, &user)
	if err == mongo.ErrNoDocuments {
		return "", errors.New(fmt.Sprintf("User %s not exist", requestData.Username))
	}
	if err != nil {
		return "", err
	}
	return user["_id"], nil
}
