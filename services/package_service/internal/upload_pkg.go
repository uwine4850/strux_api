package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"path/filepath"
	"strings"
	"strux_api/internal/config"
	"strux_api/internal/config/schema"
	"strux_api/pkg/db"
	"strux_api/pkg/logging"
	"strux_api/services/protofiles/baseproto"
	"strux_api/services/protofiles/pkgproto"
	"strux_api/services/protofiles/userproto"
	"strux_api/services/utils"
)

func UploadPkg(uploadPackage *pkgproto.RequestUploadPackage) *baseproto.BaseResponse {
	var createdPackageVersionId string
	var createdPackageId string
	var packageDirPathForRollBack string
	var userFolderPath string
	// log in user
	userLogInBaseResponse := logInUser(uploadPackage.User.Username, uploadPackage.User.Password)
	if userLogInBaseResponse.Status == baseproto.ResponseStatus_StatusError {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", userLogInBaseResponse.Message)
		return utils.SendResponseError(userLogInBaseResponse.Message)
	}
	if userLogInBaseResponse.Status == baseproto.ResponseStatus_StatusError || userLogInBaseResponse.Success == false {
		return userLogInBaseResponse
	}

	// get db operations
	client, err := db.GetMongoClient()
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
		return utils.SendResponseError(err.Error())
	}
	clientConnection, ctx, err := client.Connect()
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
		return utils.SendResponseError(err.Error())
	}
	defer func(clientConnection *mongo.Client, ctx context.Context) {
		err := clientConnection.Disconnect(ctx)
		if err != nil {
			logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
		}
	}(clientConnection, ctx)
	userOperation := utils.GetUserServiceOperation(clientConnection, ctx)
	packageOperation := utils.GetPackageServiceOperation(clientConnection, ctx)
	pkgVersionOperation := utils.GetPackageVersionServiceOperation(clientConnection, ctx)

	// get current user
	var user map[string]string
	err = userOperation.FindOneByValue("username", uploadPackage.User.Username, &user)
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
		return utils.SendResponseError(err.Error())
	}

	// init roll back
	userId := user["_id"]
	rollBack := RollBack{
		CreatedPackageId:         &createdPackageId,
		CreatedPackageVersionId:  &createdPackageVersionId,
		PackagesOperation:        &packageOperation,
		PackageVersionsOperation: &pkgVersionOperation,
		PackageDirPath:           &packageDirPathForRollBack,
		UserId:                   &userId,
		UserFolderPath:           &userFolderPath,
	}

	// find duplicated version
	res := findDuplicatedVersion(packageOperation, pkgVersionOperation, uploadPackage, user)
	if res != nil {
		return res
	}

	// create main package dir
	if !utils.PathExist(config.PackagesDirPath) {
		err := os.Mkdir(config.PackagesDirPath, os.ModePerm)
		if err != nil {
			logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
			return utils.SendResponseError(err.Error())
		}
	}

	// create user folder For example, UserPackages/username.
	userFolderPath = filepath.Join(config.PackagesDirPath, uploadPackage.User.Username)
	if !utils.PathExist(userFolderPath) {
		err := os.MkdirAll(userFolderPath, os.ModePerm)
		if err != nil {
			logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
			errRollBack := rollBack.run()
			var rollBackErrText string
			if errRollBack != nil {
				rollBackErrText = err.Error()
			}
			return utils.SendResponseError(err.Error() + "&& RollBackErr: " + rollBackErrText)
		}
	}

	// create package version dir. For example, UserPackages/username/projectName/0.0.1v.
	packageDirPath := filepath.Join(config.PackagesDirPath, uploadPackage.User.Username, uploadPackage.UplDirInfo.Name, uploadPackage.Version)
	packageDirPathForRollBack = packageDirPath
	if !utils.PathExist(packageDirPath) {
		err := os.MkdirAll(packageDirPath, os.ModePerm)
		if err != nil {
			logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
			errRollBack := rollBack.run()
			var rollBackErrText string
			if errRollBack != nil {
				rollBackErrText = err.Error()
			}
			return utils.SendResponseError(err.Error() + "&& RollBackErr: " + rollBackErrText)
		}
	}
	if !utils.PathExist(packageDirPath) {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
		errRollBack := rollBack.run()
		var rollBackErrText string
		if errRollBack != nil {
			rollBackErrText = err.Error()
		}
		return utils.SendResponseError("Package directory path not found." + "&& RollBackErr: " + rollBackErrText)
	}

	// creating a directory tree from the passed json file.
	dirTreeMap := make(map[string][]string)
	err = createDirTree(packageDirPath, "", uploadPackage.UplDirInfo, &dirTreeMap)
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
		errRollBack := rollBack.run()
		var rollBackErrText string
		if errRollBack != nil {
			rollBackErrText = err.Error()
		}
		return utils.SendResponseError(err.Error() + "&& RollBackErr: " + rollBackErrText)
	}

	// create files
	err = createFiles(packageDirPath, &uploadPackage.UplFiles, dirTreeMap)
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
		errRollBack := rollBack.run()
		var rollBackErrText string
		if errRollBack != nil {
			rollBackErrText = err.Error()
		}
		return utils.SendResponseError(err.Error() + "&& RollBackErr: " + rollBackErrText)
	}

	// commit all changes to db
	// create or get new package
	insertedPkgId, err := CreateNewPackage(packageOperation, uploadPackage, user["_id"])
	createdPackageId = insertedPkgId
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
		errRollBack := rollBack.run()
		var rollBackErrText string
		if errRollBack != nil {
			rollBackErrText = err.Error()
		}
		return utils.SendResponseError(err.Error() + "&& RollBackErr: " + rollBackErrText)
	}

	// create pkg version
	pkgVersion := schema.PackageVersion{
		ParentPackageId: insertedPkgId,
		Version:         uploadPackage.Version,
		DownloadPath:    packageDirPath,
	}
	insertedPkgVersionId, err := pkgVersionOperation.InsertOne(pkgVersion)
	createdPackageVersionId = insertedPkgVersionId.InsertedID.(primitive.ObjectID).Hex()
	if err != nil {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
		errRollBack := rollBack.run()
		var rollBackErrText string
		if errRollBack != nil {
			rollBackErrText = err.Error()
		}
		return utils.SendResponseError(err.Error() + "&& RollBackErr: " + rollBackErrText)
	}

	resp := &baseproto.BaseResponse{
		Message: fmt.Sprintf("Package %s succesfully created.", uploadPackage.UplDirInfo.Name),
		Success: true,
		Status:  baseproto.ResponseStatus_StatusOk,
	}
	return resp
}

// RollBack structure with which you can undo ALL changes in the package.
// To do this, you only need to run the run method.
// All parameters are passed by reference, so the structure can only be initialized once and used everywhere in the scope.
// The "rollBack" methods will themselves work as the necessary parameters appear.
type RollBack struct {
	CreatedPackageId         *string
	CreatedPackageVersionId  *string
	PackagesOperation        *db.DatabaseOperation
	PackageVersionsOperation *db.DatabaseOperation
	PackageDirPath           *string
	UserId                   *string
	UserFolderPath           *string
}

// run the main method that runs all the helper methods as data arrives.
func (rb *RollBack) run() error {
	if *rb.PackageDirPath != "" {
		err := rb.rollBackPackageDirPath()
		if err != nil {
			return err
		}
		err = rb.rollBackCreatedProjectFolder()
		if err != nil {
			return err
		}
	}
	if *rb.CreatedPackageId != "" {
		err := rb.rollBackCreatedPackage()
		if err != nil {
			return err
		}
	}
	if *rb.CreatedPackageVersionId != "" {
		err := rb.rollBackCreatedPackageVersion()
		if err != nil {
			return err
		}
	}
	if *rb.UserFolderPath != "" {
		err := rb.rollBackCreatedUserFolder()
		if err != nil {
			return err
		}
	}
	return nil
}

// rollBackCreatedUserFolder cancels the creation of the user folder. For example UserPackages/username.
func (rb *RollBack) rollBackCreatedUserFolder() error {
	count, err := rb.PackagesOperation.GetCountDocuments(bson.D{{"parentUserId", *rb.UserId}})
	if err != nil {
		return err
	}
	if count == 0 && !utils.PathExist(*rb.PackageDirPath) {
		err = os.RemoveAll(*rb.UserFolderPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// rollBackCreatedProjectFolder cancels the creation of the project folder. For example UserPackages/username/projectName.
func (rb *RollBack) rollBackCreatedProjectFolder() error {
	count, err := rb.PackageVersionsOperation.GetCountDocuments(bson.D{{"parentPackageId", *rb.CreatedPackageId}})
	if err != nil {
		return err
	}
	if count >= 0 && count < 2 && !utils.PathExist(*rb.PackageDirPath) {
		splitPackageDirPath := strings.Split(*rb.PackageDirPath, string(os.PathSeparator))
		pkgPath := filepath.Join(splitPackageDirPath[:len(splitPackageDirPath)-1]...)
		err = os.RemoveAll(pkgPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// rollBackPackageDirPath undoes the creation of the project version folder. For example, UserPackages/username/projectName/0.0.1v.
func (rb *RollBack) rollBackPackageDirPath() error {
	if utils.PathExist(*rb.PackageDirPath) {
		err := os.RemoveAll(*rb.PackageDirPath)
		return err
	}
	return errors.New(fmt.Sprintf("Path %s not exist", *rb.PackageDirPath))
}

// rollBackCreatedPackage removes an entry from the package database.
func (rb *RollBack) rollBackCreatedPackage() error {
	count, err := rb.PackageVersionsOperation.GetCountDocuments(bson.D{{"parentPackageId", *rb.CreatedPackageId}})
	if err != nil {
		return err
	}
	if count == 1 || count == 0 {
		objectId, err := primitive.ObjectIDFromHex(*rb.CreatedPackageId)
		if err != nil {
			return err
		}
		_, err = rb.PackagesOperation.DeleteOneEntry(bson.D{{"_id", objectId}})
		if err != nil {
			return err
		}
	}
	return nil
}

// rollBackCreatedPackageVersion removes an entry from the database for the version of the selected package.
func (rb *RollBack) rollBackCreatedPackageVersion() error {
	objectId, err := primitive.ObjectIDFromHex(*rb.CreatedPackageVersionId)
	if err != nil {
		return err
	}
	_, err = rb.PackageVersionsOperation.DeleteOneEntry(bson.D{{"_id", objectId}})
	return err
}

// findDuplicatedVersion looks for duplicate versions of a package.
// Duplicates are searched only for the package of a specific user.
func findDuplicatedVersion(packageOperation db.DatabaseOperation, pkgVersionOperation db.DatabaseOperation,
	uploadPackage *pkgproto.RequestUploadPackage, user map[string]string) *baseproto.BaseResponse {

	values := bson.D{
		{"packageName", uploadPackage.UplDirInfo.Name},
		{"parentUserId", user["_id"]},
	}
	err := packageOperation.FindOneByMultipleValues(values, map[string]string{})
	if err != nil && err != mongo.ErrNoDocuments {
		logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
		return utils.SendResponseError(err.Error())
	} else {
		var currentPackage map[string]string
		err = packageOperation.FindOneByValue("parentUserId", user["_id"], &currentPackage)
		if err != nil && err != mongo.ErrNoDocuments {
			logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
			return utils.SendResponseError(err.Error())
		}
		err = pkgVersionOperation.FindOneByMultipleValues(
			bson.D{
				{"parentPackageId", currentPackage["_id"]},
				{"version", uploadPackage.Version},
			},
			map[string]string{})
		if err != nil && err != mongo.ErrNoDocuments {
			logging.CreateLog(config.PackageServiceLogFileName, logrus.ErrorLevel, "package_service.internal", "UploadPkg", "", err.Error())
			return utils.SendResponseError(err.Error())
		}
		if err != mongo.ErrNoDocuments {
			return utils.SendResponseWarning(fmt.Sprintf("Version %s of the %s package already exists.",
				uploadPackage.Version, uploadPackage.UplDirInfo.Name))
		}
	}
	return nil
}

// CreateNewPackage creates a new package and returns its ID.
// If the package has already been created, simply returns the id.
func CreateNewPackage(packageOperation db.DatabaseOperation, uploadPackage *pkgproto.RequestUploadPackage, userId string) (string, error) {
	var findPackage map[string]string
	values := bson.D{
		{"packageName", uploadPackage.UplDirInfo.Name},
		{"parentUserId", userId},
	}
	err := packageOperation.FindOneByMultipleValues(values, &findPackage)
	if err != nil && err != mongo.ErrNoDocuments {
		return "", err
	}
	var insertedPkgId string
	if findPackage == nil {
		pkg := schema.Package{
			ParentUserId: userId,
			PackageName:  uploadPackage.UplDirInfo.Name,
		}
		insertedPkg, err := packageOperation.InsertOne(pkg)
		if err != nil {
			return "", err
		}
		insertedPkgId = insertedPkg.InsertedID.(primitive.ObjectID).Hex()
	} else {
		insertedPkgId = findPackage["_id"]
	}
	return insertedPkgId, nil
}

// createFiles creates files from a list in a pre-created directory tree.
func createFiles(packageDirPath string, files *[]*pkgproto.UploadFile, dirTree map[string][]string) error {
	for i := 0; i < len(*files); i++ {
		for dirPath, dirFiles := range dirTree {
			for j := 0; j < len(dirFiles); j++ {
				filePath := filepath.Join(packageDirPath, (*files)[i].FileName)
				// the paths of the files from the directory tree and the files received from the form are the same.
				if filepath.Join(packageDirPath, dirPath, dirFiles[j]) == filePath && !utils.PathExist(filePath) {
					err := os.WriteFile(filePath, (*files)[i].FileBytesData, os.ModePerm)
					if err != nil {
						return err
					}
					dirTree[dirPath] = append(dirTree[dirPath][:j], dirTree[dirPath][j+1:]...)
				}
			}
		}
	}
	return nil
}

// createDirTree recursively creates and returns a package directory tree.
func createDirTree(packageDirPath string, parentDirName string, uploadDirInfo *pkgproto.UploadDirInfo, dirTreeMap *map[string][]string) error {
	dirPath := ""
	if parentDirName == "" {
		dirPath = uploadDirInfo.Name
	} else {
		dirPath = filepath.Join(parentDirName, uploadDirInfo.Name)
	}

	if !utils.PathExist(filepath.Join(packageDirPath, dirPath)) {
		err := os.MkdirAll(filepath.Join(packageDirPath, dirPath), os.ModePerm)
		if err != nil {
			return err
		}
	}
	for i := 0; i < len(uploadDirInfo.FileNames); i++ {
		if utils.PathExist(filepath.Join(packageDirPath, dirPath)) {
			(*dirTreeMap)[dirPath] = append((*dirTreeMap)[dirPath], uploadDirInfo.FileNames[i])
		}
	}
	if len(uploadDirInfo.FileNames) == 0 {
		if utils.PathExist(filepath.Join(packageDirPath, dirPath)) {
			(*dirTreeMap)[dirPath] = append((*dirTreeMap)[dirPath], "")
		}
	}

	if uploadDirInfo.InnerDir != nil {
		for i := 0; i < len(uploadDirInfo.InnerDir); i++ {
			err := createDirTree(packageDirPath, dirPath, uploadDirInfo.InnerDir[i], dirTreeMap)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func connectToUserService() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	connection, err := grpc.Dial(config.UserServiceAddress, opts...)
	return connection, err
}

// logInUser verify that the user is logged in correctly. A user is considered logged in if the password and username are correct.
func logInUser(username string, password string) *baseproto.BaseResponse {
	userServiceConnection, err := connectToUserService()
	if err != nil {
		return utils.SendResponseError(err.Error())
	}
	userClient := userproto.NewUserClient(userServiceConnection)
	user := &userproto.RequestUserLogIn{
		Username: username,
		Password: password,
	}
	response, err := userClient.UserLogIn(context.Background(), user)
	if err != nil {
		return utils.SendResponseError(err.Error())
	}

	// close connection
	err = userServiceConnection.Close()
	if err != nil {
		return utils.SendResponseError(err.Error())
	}
	return response
}
