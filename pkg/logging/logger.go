package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// CreateLog Creates one entry in the log file.
func CreateLog(filePath string, logLevel log.Level, pkgName string, callFunc string, descr string, error string) {
	f, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	log.SetOutput(f)
	log.SetLevel(logLevel)
	log.WithFields(log.Fields{
		"package": pkgName,
		"func":    callFunc,
		"descr":   descr,
		"error":   error,
	})
}
