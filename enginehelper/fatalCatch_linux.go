package core

import (
	"os"
	"syscall"
)

func RedirectStderr() (err error) {
	logFile, err := getLogFile()
	err = syscall.Dup3(int(logFile.Fd()), int(os.Stderr.Fd()), 0)
	if err != nil {
		return
	}
	return
}
