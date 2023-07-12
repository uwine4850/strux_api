package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/uwine4850/strux_api/internal/config"
	"github.com/uwine4850/strux_api/internal/config/schema"
	"github.com/uwine4850/strux_api/pkg/db"
	"github.com/uwine4850/strux_api/pkg/logging"
	"github.com/uwine4850/strux_api/pkg/uplutils"
	"github.com/uwine4850/strux_api/services/protofiles/baseproto"
	"github.com/uwine4850/strux_api/services/protofiles/pkgproto"
	"github.com/uwine4850/strux_api/services/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"path/filepath"
)

func DownloadPackage(requestData *pkgproto.RequestDownloadPackage) *pkgproto.MutateDownloadBaseResponse {
	client, err := db.GetMongoClient()
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "DownloadPackage", "", err.Error())
		return sendMutateDownloadBaseResponseError(err.Error())
	}
	clientConnection, ctx, err := client.Connect()
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "DownloadPackage", "", err.Error())
		return sendMutateDownloadBaseResponseError(err.Error())
	}
	defer func(clientConnection *mongo.Client, ctx context.Context) {
		err := clientConnection.Disconnect(ctx)
		if err != nil {
			logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "DownloadPackage", "", err.Error())
		}
	}(clientConnection, ctx)

	// Checking the existence of the package. If it is found, return the path to download it.
	downloadPath, err := checkingPackageExistsAndGetDownloadPath(requestData, clientConnection, ctx)
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "DownloadPackage", "", err.Error())
		return sendMutateDownloadBaseResponseError(err.Error())
	}

	// Getting data about the package (directory).
	dirInfo, err := uplutils.GetDirsInfo(filepath.Join(downloadPath, requestData.PkgName), requestData.PkgName)
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "DownloadPackage", "", err.Error())
		return sendMutateDownloadBaseResponseError(err.Error())
	}

	// Convert pkgproto.UploadDirInfo(dirInfo) to json(map).
	var uploadDirJson []byte
	err = uplutils.UploadDirInfoToJson(dirInfo, &uploadDirJson)
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "DownloadPackage", "", err.Error())
		return sendMutateDownloadBaseResponseError(err.Error())
	}

	// Creates a new formatted structure pkgproto.UploadDirInfo from dirInfo
	uplFiles, err := uplutils.CreateUploadFilePaths(dirInfo, downloadPath)
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "DownloadPackage", "", err.Error())
		return sendMutateDownloadBaseResponseError(err.Error())
	}
	resp := &pkgproto.MutateDownloadBaseResponse{
		BaseResponse: &baseproto.BaseResponse{
			Message: "OK",
			Success: true,
			Status:  baseproto.ResponseStatus_StatusOk,
		},
		UplFiles:           uplFiles,
		UplDirInfo:         dirInfo,
		UploadDirsInfoJson: uploadDirJson,
	}
	return resp
}

// sendMutateDownloadBaseResponseError Special method to send pkgproto.MutateDownloadBaseResponse error
func sendMutateDownloadBaseResponseError(msg string) *pkgproto.MutateDownloadBaseResponse {
	resp := &pkgproto.MutateDownloadBaseResponse{
		BaseResponse: &baseproto.BaseResponse{
			Message: msg,
			Success: false,
			Status:  baseproto.ResponseStatus_StatusError,
		},
		UplFiles:   nil,
		UplDirInfo: nil,
	}
	return resp
}

// checkingPackageExistsAndGetDownloadPath Accessing the Packet Checker service.
// If a package is found, return the path for downloading it.
func checkingPackageExistsAndGetDownloadPath(requestData *pkgproto.RequestDownloadPackage, dbClient *mongo.Client,
	ctx context.Context) (string, error) {
	service, err := utils.ConnectToPackageService()
	if err != nil {
		return "", err
	}
	packageClient := pkgproto.NewPackageClient(service)
	request := &pkgproto.RequestPackageExists{
		Username: requestData.Username,
		PkgName:  requestData.PkgName,
		Version:  requestData.Version,
	}
	existsPackage, err := packageClient.ExistsPackage(context.Background(), request)
	if err != nil {
		return "", err
	}

	if existsPackage.Success {
		userOperation := utils.GetUserServiceOperation(dbClient, ctx)
		var user map[string]string
		err := userOperation.FindOneByValue("username", requestData.Username, &user)
		if err != nil {
			return "", err
		}
		packageOperation := utils.GetPackageServiceOperation(dbClient, ctx)
		var userPackage map[string]string
		err = packageOperation.FindOneByMultipleValues(bson.D{
			{"parentUserId", user["_id"]},
			{"packageName", requestData.PkgName},
		}, &userPackage)
		if err != nil {
			return "", err
		}
		packageVersionOperation := utils.GetPackageVersionServiceOperation(dbClient, ctx)
		var userPackageVersion schema.PackageVersion
		err = packageVersionOperation.FindOneByMultipleValues(bson.D{
			{"parentPackageId", userPackage["_id"]},
			{"version", requestData.Version},
		}, &userPackageVersion)
		if err != nil {
			return "", err
		}
		return userPackageVersion.DownloadPath, nil
	}

	err = service.Close()
	if err != nil {
		return "", err
	}
	return "", errors.New(fmt.Sprintf("Package %s/%s not exist.", requestData.PkgName, requestData.Version))
}
