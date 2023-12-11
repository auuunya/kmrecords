package kmrecords

var (
	// https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-setwindowshookexa
	WH_CALLWNDPROC     = 4
	WH_CALLWNDPROCRET  = 12
	WH_CBT             = 5
	WH_DEBUG           = 9
	WH_FOREGROUNDIDLE  = 11
	WH_GETMESSAGE      = 3
	WH_JOURNALPLAYBACK = 1
	WH_JOURNALRECORD   = 0
	WH_KEYBOARD        = 2
	WH_KEYBOARD_LL     = 13
	WH_MOUSE           = 7
	WH_MOUSE_LL        = 14
	WH_MSGFILTER       = -1
	WH_SHELL           = 10
	WH_SYSMSGFILTER    = 6
)

type MsllHookStruct struct {
	Pt          Point
	MouseData   uintptr
	Flags       uintptr
	Time        uintptr
	DwExtraInfo uintptr
}

type KbDllHookStruct struct {
	VkCode      uintptr
	ScanCode    uintptr
	Flags       uintptr
	Time        uintptr
	DwExtraInfo uintptr
}

type HookData struct {
	Type     int
	HookProc func(nCode int, wParam, lParam uintptr) uintptr
}

func NewHookData(wh int, proc func(nCode int, wParam, lParam uintptr) uintptr) *HookData {
	return &HookData{
		Type:     wh,
		HookProc: proc,
	}
}
