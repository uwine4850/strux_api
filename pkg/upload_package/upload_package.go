package upload_package

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strux_api/services/protofiles/pkgproto"
	"strux_api/services/utils"
)

// GetDirsInfo Recursively populates the pkgproto.UploadDirInfo structure.
// The name contains the relative path to one directory of the package. For example, the first directory is packageDir,
// the second is packageDir2 and so on.
// FileNames contains the name of all files in the current directory.
// InnerDir is a slice of pkgproto.UploadDirInfo which contains internal directories.
func GetDirsInfo(dir string, parentDirName string) (*pkgproto.UploadDirInfo, error) {
	readDir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var innerDirs []string
	tee := &pkgproto.UploadDirInfo{}
	tee.Name = parentDirName

	for i := 0; i < len(readDir); i++ {
		if readDir[i].IsDir() {
			innerDirs = append(innerDirs, readDir[i].Name())
		} else {
			tee.FileNames = append(tee.FileNames, readDir[i].Name())
		}
	}
	for i := 0; i < len(innerDirs); i++ {
		iDir, err := GetDirsInfo(filepath.Join(dir, innerDirs[i]), filepath.Join(parentDirName, innerDirs[i]))
		if err != nil {
			return nil, err
		}
		tee.InnerDir = append(tee.InnerDir, iDir)
	}
	return tee, nil
}

// UploadDirInfoToJson converts the pkgproto.UploadDirInfo structure into json(map) format.
func UploadDirInfoToJson(uploadDirInfo *pkgproto.UploadDirInfo, mapData *[]byte) error {
	marshal, err := json.Marshal(uploadDirInfo)
	if err != nil {
		return err
	}
	*mapData = marshal
	//err = json.Unmarshal(marshal, mapData)
	//if err != nil {
	//	return err
	//}
	return nil
}

// CreateUploadFilePaths It formats the information about the batch files that are passed through the form.
// The file name is a formatted key for the form, and FileNames contains the absolute paths to the loaded files.
// InnerDir contains identically formatted directories.
func CreateUploadFilePaths(dirInfo *pkgproto.UploadDirInfo, downloadPath string) (*pkgproto.UploadDirInfo, error) {
	uplDirInfo := &pkgproto.UploadDirInfo{}
	absDownload, err := filepath.Abs(downloadPath)
	if err != nil {
		return nil, err
	}
	typ := reflect.Indirect(reflect.ValueOf(dirInfo)).Type()
	val := reflect.Indirect(reflect.ValueOf(dirInfo))
	for i := 0; i < typ.NumField(); i++ {
		fieldType := typ.Field(i)
		fieldValue := val.Field(i)
		if fieldType.IsExported() {
			if fieldType.Name == "Name" {
				uplDirInfo.Name = fieldValue.String()
			}
			if fieldType.Name == "FileNames" {
				for j := 0; j < fieldValue.Len(); j++ {
					filePath := filepath.Join(absDownload, uplDirInfo.Name, fieldValue.Index(j).String())
					if utils.PathExist(filePath) {
						uplDirInfo.FileNames = append(uplDirInfo.FileNames, filePath)
					} else {
						err := errors.New(fmt.Sprintf("File path %s not exist.", filePath))
						return nil, err
					}
				}
			}
			if fieldType.Name == "InnerDir" {
				for j := 0; j < fieldValue.Len(); j++ {
					uplFile, err := CreateUploadFilePaths(fieldValue.Index(j).Interface().(*pkgproto.UploadDirInfo), downloadPath)
					if err != nil {
						return nil, err
					}
					uplDirInfo.InnerDir = append(uplDirInfo.InnerDir, uplFile)
				}
			}
		}
	}
	uplDirInfo.Name = "DIR$" + uplDirInfo.Name
	return uplDirInfo, nil
}

// UplFilesToMap Converts the response of the CreateUploadFilePaths function into a map in the following format:
// The key is the name of the package directory and at the same time the name for the field in the form.
// The slice values are the absolute paths to the files of the given directory.
// Thus, data is generated for uploading files to the form, where the key is the name of the field and the slice is
// the paths for reading and adding files to the field.
func UplFilesToMap(uplFiles *pkgproto.UploadDirInfo) map[string][]string {
	uplFilesMap := make(map[string][]string)
	typ := reflect.Indirect(reflect.ValueOf(uplFiles)).Type()
	val := reflect.Indirect(reflect.ValueOf(uplFiles))
	var pkgName string
	for i := 0; i < typ.NumField(); i++ {
		fieldType := typ.Field(i)
		fieldValue := val.Field(i)
		if fieldType.IsExported() {
			if fieldType.Name == "Name" {
				uplFilesMap[fieldValue.String()] = []string{}
				pkgName = fieldValue.String()
			}
			if fieldType.Name == "FileNames" {
				for j := 0; j < fieldValue.Len(); j++ {
					uplFilesMap[pkgName] = append(uplFilesMap[pkgName], fieldValue.Index(j).String())
				}
			}
			if fieldType.Name == "InnerDir" {
				for j := 0; j < fieldValue.Len(); j++ {
					for key, value := range UplFilesToMap(fieldValue.Index(j).Interface().(*pkgproto.UploadDirInfo)) {
						uplFilesMap[key] = value
					}
				}
			}
		}
	}
	return uplFilesMap
}
