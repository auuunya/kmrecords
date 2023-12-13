package kmrecords

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32 = syscall.NewLazyDLL("User32.dll")

	procGetCapture          = user32.NewProc("GetCapture")
	procGetCursorPos        = user32.NewProc("GetCursorPos")
	procGetDesktopWindow    = user32.NewProc("GetDesktopWindow")
	procDragDetect          = user32.NewProc("DragDetect")
	procCallWindowProcA     = user32.NewProc("CallWindowProcA")
	procSetWindowsHookExA   = user32.NewProc("SetWindowsHookExA")
	procCallNextHookEx      = user32.NewProc("CallNextHookEx")
	procGetMessageA         = user32.NewProc("GetMessageA")
	procSendInput           = user32.NewProc("SendInput")
	procRegisterHotKey      = user32.NewProc("RegisterHotKey")
	procGetAsyncKeyState    = user32.NewProc("GetAsyncKeyState")
	procGetKeyState         = user32.NewProc("GetKeyState")
	procGetKeyNameTextA     = user32.NewProc("GetKeyNameTextA")
	procTranslateMessage    = user32.NewProc("TranslateMessage")
	procPeekMessageA        = user32.NewProc("PeekMessageA")
	procUnhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	procPostQuitMessage     = user32.NewProc("PostQuitMessage")
	MouseEventChan          = make(chan interface{})
)

func GetDesktopWindowHWDN() uintptr {
	window, _, _ := procGetDesktopWindow.Call()
	return window
}

func SetHooks(h *HookData) uintptr {
	hook, _, err := procSetWindowsHookExA.Call(
		uintptr(h.Type),
		uintptr(syscall.NewCallback(h.HookProc)),
		uintptr(0),
		uintptr(0),
	)
	if err != nil && err.Error() != "The operation completed successfully." {
		log.Fatalf("error: %v\n", err)
	}
	return hook
}

func GetMsgProc(nCode int, wParam, lParam uintptr) uintptr {
	return CallNextHookEx(0, nCode, wParam, lParam)
}

func HookMProcCallback(nCode int, wParam, lParam uintptr) (ret uintptr) {
	if nCode < 0 {
		return CallNextHookEx(0, nCode, wParam, lParam)
	}

	point := (*MsllHookStruct)(unsafe.Pointer(lParam)).Pt
	mouseinfo := &MouseInfo{
		X:         point.X,
		Y:         point.Y,
		Ctrl:      GetAsyncKeyState(VK_CONTROL),
		Shift:     GetAsyncKeyState(VK_SHIFT),
		Alt:       GetAsyncKeyState(VK_MENU),
		TimeStamp: time.Now().Unix(),
	}
	switch wParam {
	case uintptr(WM_LBUTTONDOWN):
		mouseinfo.Event = WM_LBUTTONDOWN
	case uintptr(WM_LBUTTONUP):
		mouseinfo.Event = WM_LBUTTONUP
	case uintptr(WM_MOUSEMOVE):
		mouseinfo.Event = WM_MOUSEMOVE
	case uintptr(WM_MOUSEWHEEL):
		mouseinfo.Event = WM_MOUSEWHEEL
	case uintptr(WM_RBUTTONDOWN):
		mouseinfo.Event = WM_RBUTTONDOWN
	case uintptr(WM_RBUTTONUP):
		mouseinfo.Event = WM_RBUTTONUP
	}
	// fmt.Printf("鼠标操作，鼠标信息: %#v\n", mouseinfo)
	MouseEventChan <- mouseinfo
	return CallNextHookEx(0, nCode, wParam, lParam)
}

func HookKProcCallback(nCode int, wParam, lParam uintptr) (ret uintptr) {
	if nCode < 0 {
		return CallNextHookEx(0, nCode, wParam, lParam)
	}

	param := (*KbDllHookStruct)(unsafe.Pointer(lParam))
	keyboradinfo := &KeyBoradInfo{
		VK_Code:   param.VkCode & 0xff,
		Ctrl:      GetAsyncKeyState(VK_CONTROL),
		Shift:     GetAsyncKeyState(VK_SHIFT),
		Alt:       GetAsyncKeyState(VK_MENU),
		Win:       GetAsyncKeyState(VK_LWIN | VK_RWIN),
		CapsLock:  GetAsyncKeyState(VK_CAPITAL),
		TimeStamp: time.Now().Unix(),
	}
	switch wParam {
	case uintptr(WM_KEYDOWN):
		keyboradinfo.VK_Type = WM_KEYDOWN
	case uintptr(WM_KEYUP):
		keyboradinfo.VK_Type = WM_KEYUP
	case uintptr(WM_SYSKEYDOWN):
		keyboradinfo.VK_Type = WM_SYSKEYDOWN
	case uintptr(WM_SYSKEYUP):
		keyboradinfo.VK_Type = WM_SYSKEYUP
	}
	// fmt.Printf("键盘操作, %#v\n", keyboradinfo)
	MouseEventChan <- keyboradinfo
	return CallNextHookEx(0, nCode, wParam, lParam)
}

func CallNextHookEx(hhk uintptr, nCode int, wParam uintptr, lParam uintptr) uintptr {
	ret, _, _ := procCallNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return ret
}

func GetMessageA(msg *MSG) bool {
	ret, _, _ := procGetMessageA.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(0),
		uintptr(0),
		uintptr(0),
	)
	return ret != 0
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

func GetAsyncKeyState(key int) int {
	state, _, _ := procGetAsyncKeyState.Call(
		uintptr(key),
	)
	return int(state)
}

func GetKeyState(key int) int {
	state, _, _ := procGetKeyState.Call(
		uintptr(key),
	)
	return int(state)
}

func GetKeyName() string {
	var name string
	textA, _, _ := procGetKeyNameTextA.Call(
		uintptr(WM_KEYDOWN),
		uintptr(unsafe.Pointer(&name)),
		uintptr(255),
	)
	fmt.Printf("textA: %#v\n", textA)
	fmt.Printf("name: %#v\n", name)
	return name
}

func TransLateMessage(msg *MSG) {
	ret, _, _ := procTranslateMessage.Call(
		uintptr(unsafe.Pointer(msg)),
	)
	fmt.Printf("Ret: %#v\n", ret)
}

func PeekMessage(msg *MSG) int {
	ret, _, _ := procPeekMessageA.Call(
		uintptr(unsafe.Pointer(&msg)),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(PM_REMOVE),
	)
	fmt.Printf("msg: %#v\n", msg)
	return int(ret)
}

func UnWindowHook(hhk uintptr) {
	procUnhookWindowsHookEx.Call(
		uintptr(hhk),
	)
}

func PostQuitMsg() {
	procPostQuitMessage.Call(
		uintptr(WM_QUIT),
	)
}
