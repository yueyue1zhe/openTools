package core

import (
	"os"
)

func fileExists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func getLogFile() (logFile *os.File, err error) {
	logPath := "./err.log"
	if !fileExists(logPath) {
		logFile, err = os.Create(logPath)
		if err != nil {
			return
		}
	} else {
		logFile, err = os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0644)
		if err != nil {
			return
		}
	}
	return
}
