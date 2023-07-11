package uplutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uwine4850/strux_api/services/protofiles/pkgproto"
	"github.com/uwine4850/strux_api/services/utils"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
	"strings"
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

// CreateDirTree Recursively creates a directory tree at the specified path.
// After that it returns a map with directory addresses and names of their files.
func CreateDirTree(packageDirPath string, uploadDirInfo *pkgproto.UploadDirInfo, dirTreeMap *map[string][]string) error {
	dirPath := uploadDirInfo.Name
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
			err := CreateDirTree(packageDirPath, uploadDirInfo.InnerDir[i], dirTreeMap)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// SetUploadFiles extracts files from the form and writes them to the pkgproto.UploadFile structure.
// If the form field name starts with DIR$ it means that this is the path to the directory with files.
// The value of the files in this field corresponds to the files in the same directory of the file system.
func SetUploadFiles(filesData map[string][]*multipart.FileHeader, uploadFiles *[]*pkgproto.UploadFile) error {
	for names, value := range filesData {
		dirName := strings.Split(names, "DIR$")
		if len(dirName) > 2 {
			return errors.New(fmt.Sprintf("The path \"%s\" could not be processed.", names))
		}
		if len(dirName) < 2 {
			continue
		}
		for j := 0; j < len(value); j++ {
			open, err := value[j].Open()
			if err != nil {
				return err
			}
			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(open)
			if err != nil {
				return err
			}
			if buf == nil {
				continue
			}
			*uploadFiles = append(*uploadFiles, &pkgproto.UploadFile{
				FileName:      filepath.Join(dirName[1], value[j].Filename),
				FileBytesData: buf.Bytes(),
			})
		}
	}
	return nil
}

// CreateFiles creates files in their parent directories.
// Skips a directory if there are no files in it.
func CreateFiles(packageDirPath string, files *[]*pkgproto.UploadFile, dirTree map[string][]string) error {
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
