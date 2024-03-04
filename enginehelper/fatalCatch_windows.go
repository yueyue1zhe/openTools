package core

import (
	"os"
	"syscall"
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

func setStdHandle(stdhandle int32, handle syscall.Handle) error {
	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}

// RedirectStderr to the file passed in
func RedirectStderr() (err error) {
	logFile, err := getLogFile()
	err = setStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(logFile.Fd()))
	if err != nil {
		return
	}
	// SetStdHandle does not affect prior references to stderr
	os.Stderr = logFile
	return
}
