package kmrecords

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("User32.dll")

	procGetCapture        = user32.NewProc("GetCapture")
	procGetCursorPos      = user32.NewProc("GetCursorPos")
	procGetDesktopWindow  = user32.NewProc("GetDesktopWindow")
	procDragDetect        = user32.NewProc("DragDetect")
	procCallWindowProcA   = user32.NewProc("CallWindowProcA")
	procSetWindowsHookExA = user32.NewProc("SetWindowsHookExA")
	procCallNextHookEx    = user32.NewProc("CallNextHookEx")
	procGetMessageW       = user32.NewProc("GetMessageW")
	procSendInput         = user32.NewProc("SendInput")
	procRegisterHotKey    = user32.NewProc("RegisterHotKey")
)

type HookProc func(int, uintptr, uintptr) uintptr

func GetDesktopWindowHWDN() uintptr {
	window, _, _ := procGetDesktopWindow.Call()
	return window
}

func SetHooks(wh int, callback HookProc) uintptr {
	hook, _, _ := procSetWindowsHookExA.Call(
		uintptr(wh),
		uintptr(syscall.NewCallback(callback)),
		uintptr(0),
		uintptr(0),
	)
	return hook
}

func HookMProcCallback(nCode int, wParam, lParam uintptr) (ret uintptr) {
	if nCode < 0 {
		ret, _ = CallNextHookEx(0, nCode, wParam, lParam)
	}
	switch wParam {
	case uintptr(WM_LBUTTONDOWN):
		fmt.Printf("鼠标左键按下，鼠标信息: %v\n", (*MsllHookStruct)(unsafe.Pointer(lParam)))
	case uintptr(WM_LBUTTONUP):
		fmt.Printf("鼠标左键释放，鼠标信息: %v\n", (*MsllHookStruct)(unsafe.Pointer(lParam)))
	case uintptr(WM_MOUSEMOVE):
		fmt.Printf("鼠标移动，鼠标信息: %v\n", (*MsllHookStruct)(unsafe.Pointer(lParam)))
	case uintptr(WM_MOUSEWHEEL):
		fmt.Printf("鼠标滚动，鼠标信息: %v\n", (*MsllHookStruct)(unsafe.Pointer(lParam)))
	case uintptr(WM_RBUTTONDOWN):
		fmt.Printf("鼠标右键按下，鼠标信息: %v\n", (*MsllHookStruct)(unsafe.Pointer(lParam)))
	case uintptr(WM_RBUTTONUP):
		fmt.Printf("鼠标右键释放，鼠标信息: %v\n", (*MsllHookStruct)(unsafe.Pointer(lParam)))
	}
	return
}

func HookKProcCallback(nCode int, wParam, lParam uintptr) (ret uintptr) {
	if nCode < 0 {
		ret, _ = CallNextHookEx(0, nCode, wParam, lParam)
	}
	switch wParam {
	case uintptr(WM_KEYDOWN):
		fmt.Printf("键盘按下, %#v\n", (*KbDllHookStruct)(unsafe.Pointer(lParam)))
	case uintptr(WM_KEYUP):
		fmt.Printf("键盘抬起, %#v\n", (*KbDllHookStruct)(unsafe.Pointer(lParam)))
	case uintptr(WM_SYSKEYDOWN):
		fmt.Printf("键盘SYSDOWN, %#v\n", (*KbDllHookStruct)(unsafe.Pointer(lParam)))
	case uintptr(WM_SYSKEYUP):
		fmt.Printf("键盘SYSUP, %#v\n", (*KbDllHookStruct)(unsafe.Pointer(lParam)))
	}
	return
}

func CallNextHookEx(hhk uintptr, nCode int, wParam uintptr, lParam uintptr) (uintptr, error) {
	ret, _, err := procCallNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	if ret == 0 {
		return 0, err
	}
	return ret, nil
}

func GetMessageW(msg *MSG) int {
	ret, _, _ := procGetMessageW.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(0),
		uintptr(0),
		uintptr(0),
	)
	fmt.Printf("ret: %#v\n", ret)
	return int(ret)
}

func GetCursorPoint() *Point {
	var pointer *Point
	ret, _, _ := procGetCursorPos.Call(
		uintptr(unsafe.Pointer(&pointer)),
	)
	fmt.Printf("ret: %#v\n", ret)
	fmt.Printf("pointer: %#v\n", pointer)
	return pointer
}
