package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	user32 = "user32.dll"
	//procEnumWindows    = "EnumWindows"
	procGetWindowTextW = "GetWindowTextW"
	procGetClassNameW  = "GetClassNameW"
)

var (
	modUser32         = syscall.NewLazyDLL(user32)
	procEnumWindows   = modUser32.NewProc("EnumWindows")
	procGetWindowText = modUser32.NewProc(procGetWindowTextW)
	procGetClassName  = modUser32.NewProc(procGetClassNameW)
)

// callback 用于 EnumWindows 函数，遍历所有窗口
func callback(hwnd syscall.Handle, lParam uintptr) uintptr {
	//var titleLen uint32
	//var classNameLen uint32

	// 获取窗口标题
	titleBuf := make([]uint16, 256)
	_, _, _ = procGetWindowText.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&titleBuf[0])), uintptr(len(titleBuf)))
	title := syscall.UTF16ToString(titleBuf)

	// 获取窗口类名
	classNameBuf := make([]uint16, 256)
	_, _, _ = procGetClassName.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&classNameBuf[0])), uintptr(len(classNameBuf)))
	className := syscall.UTF16ToString(classNameBuf)
	if title == "直播伴侣" || title == "抖音聊天" {
		fmt.Printf("HWND: %v, Class: %s, Title: %s\n", hwnd, className, title)
	}

	// 继续枚举
	return 1
}

func main() {
	// 枚举所有顶级窗口
	// 使用 syscall.NewCallbackCDecl 创建回调函数的 Windows 兼容版本
	cb := syscall.NewCallbackCDecl(callback)

	// 枚举所有顶级窗口
	r, _, _ := procEnumWindows.Call(uintptr(cb), 0)
	if r == 0 {
		fmt.Println("EnumWindows failed")
	}
}
