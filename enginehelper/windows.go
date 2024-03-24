package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

// RECT 对应 Windows API 中的 RECT 结构体
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

const (
	user32 = "user32.dll"

	// Windows API 函数标识符
	procFindWindowA   = "FindWindowA"
	procFindWindowExA = "FindWindowExA"
	//procGetWindowRect = "GetWindowRect"
)

var (
	// 加载 user32.dll 中的函数
	modUser32         = syscall.NewLazyDLL(user32)
	procFindWindow    = modUser32.NewProc(procFindWindowA)
	procFindWindowEx  = modUser32.NewProc(procFindWindowExA)
	procGetWindowRect = modUser32.NewProc("GetWindowRect")
)

func findWindow(className, windowName string) (hwnd syscall.Handle, err error) {
	// 调用 FindWindowA 函数
	r0, _, e1 := syscall.Syscall(procFindWindow.Addr(), 2,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(className))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowName))),
		0)
	hwnd = syscall.Handle(r0)
	if hwnd == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func findWindowEx(parent, childAfter syscall.Handle, className, windowName string) (hwnd syscall.Handle, err error) {
	// 调用 FindWindowExA 函数
	r0, _, e1 := syscall.Syscall6(procFindWindowEx.Addr(), 4,
		uintptr(parent), uintptr(childAfter),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(className))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowName))),
		0, 0)
	hwnd = syscall.Handle(r0)
	if hwnd == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getWindowRect(hwnd syscall.Handle, rect *RECT) (err error) {
	// 调用 GetWindowRect 函数
	r1, _, e1 := syscall.Syscall(procGetWindowRect.Addr(), 2,
		uintptr(hwnd), uintptr(unsafe.Pointer(rect)),
		0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func main() {
	// 假设你知道目标窗口的类名（ClassName）和窗口名（WindowName）
	className := "Chrome_WidgetWin_1" // 记事本窗口的类名
	windowName := "直播伴侣"              // 窗口的标题

	// 查找目标窗口的句柄
	hwnd, err := findWindow(className, windowName)
	if err != nil {
		fmt.Println("未找到窗口:", err)
		return
	}

	// 假设你知道要查找的控件的类名（ChildClassName）
	childClassName := "EDIT" // 记事本中的编辑框的类名

	// 从目标窗口开始查找子控件
	hwndChild, err := findWindowEx(hwnd, 0, childClassName, "")
	if err != nil {
		fmt.Println("未找到子控件:", err)
		return
	}
	// 定义 RECT 结构体实例来接收控件的位置和大小
	var rect RECT

	// 调用 GetWindowRect 函数获取控件的位置和大小
	err = getWindowRect(hwndChild, &rect)
	if err != nil {
		fmt.Println("获取控件位置失败:", err)
		return
	}

	// 打印控件的位置和大小
	fmt.Printf("控件位置: Left=%d, Top=%d, Right=%d, Bottom=%d\n", rect.Left, rect.Top, rect.Right, rect.Bottom)
	fmt.Printf("控件宽度: %d, 控件高度: %d\n", rect.Right-rect.Left, rect.Bottom-rect.Top)
}
