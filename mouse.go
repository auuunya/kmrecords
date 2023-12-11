package kmrecords

var (
	WM_CAPTURECHANGED  = 0x0215
	WM_LBUTTONDBLCLK   = 0x0203
	WM_LBUTTONDOWN     = 0x0201
	WM_LBUTTONUP       = 0x0202
	WM_MBUTTONDBLCLK   = 0x0209
	WM_MBUTTONDOWN     = 0x0207
	WM_MBUTTONUP       = 0x0208
	WM_MOUSEACTIVATE   = 0x0021
	WM_MOUSEHOVER      = 0x02A1
	WM_MOUSEHWHEEL     = 0x020E
	WM_MOUSELEAVE      = 0x02A3
	WM_MOUSEMOVE       = 0x0200
	WM_MOUSEWHEEL      = 0x020A
	WM_NCHITTEST       = 0x0084
	WM_NCLBUTTONDBLCLK = 0x00A3
	WM_NCLBUTTONDOWN   = 0x00A1
	WM_NCLBUTTONUP     = 0x00A2
	WM_NCMBUTTONDBLCLK = 0x00A9
	WM_NCMBUTTONDOWN   = 0x00A7
	WM_NCMBUTTONUP     = 0x00A8
	WM_NCMOUSEHOVER    = 0x02A0
	WM_NCMOUSELEAVE    = 0x02A2
	WM_NCMOUSEMOVE     = 0x00A0
	WM_NCRBUTTONDBLCLK = 0x00A6
	WM_NCRBUTTONDOWN   = 0x00A4
	WM_NCRBUTTONUP     = 0x00A5
	WM_NCXBUTTONDBLCLK = 0x00AD
	WM_NCXBUTTONDOWN   = 0x00AB
	WM_NCXBUTTONUP     = 0x00AC
	WM_RBUTTONDBLCLK   = 0x0206
	WM_RBUTTONDOWN     = 0x0204
	WM_RBUTTONUP       = 0x0205
	WM_XBUTTONDBLCLK   = 0x020D
	WM_XBUTTONDOWN     = 0x020B
	WM_XBUTTONUP       = 0x020C
)

const (
	// WM_NCXBUTTONDBLCLK
	XBUTTON1 = 0x0001 // 双击第一个 X 按钮
	XBUTTON2 = 0x0002 // 双击第二个 X 按钮
)

var (
	MOUSEEVENTF_MOVE            = 0x0001
	MOUSEEVENTF_LEFTDOWN        = 0x0002
	MOUSEEVENTF_LEFTUP          = 0x0004
	MOUSEEVENTF_RIGHTDOWN       = 0x0008
	MOUSEEVENTF_RIGHTUP         = 0x0010
	MOUSEEVENTF_MIDDLEDOWN      = 0x0020
	MOUSEEVENTF_MIDDLEUP        = 0x0040
	MOUSEEVENTF_XDOWN           = 0x0080
	MOUSEEVENTF_XUP             = 0x0100
	MOUSEEVENTF_WHEEL           = 0x0800
	MOUSEEVENTF_HWHEEL          = 0x1000
	MOUSEEVENTF_MOVE_NOCOALESCE = 0x2000
	MOUSEEVENTF_VIRTUALDESK     = 0x4000
	MOUSEEVENTF_ABSOLUTE        = 0x8000
)

var (
	// WM_LBUTTONDBLCLK
	MK_CONTROL  = 0x0008 // 按下了 CTRL 键
	MK_LBUTTON  = 0x0001 // 按下了鼠标左键
	MK_MBUTTON  = 0x0010 // 按下了鼠标中键
	MK_RBUTTON  = 0x0002 // 按下了鼠标右键
	MK_SHIFT    = 0x0004 // 按下了 SHIFT 键
	MK_XBUTTON1 = 0x0020 // 按下了第一个 X 按钮
	MK_XBUTTON2 = 0x0040 // 按下了第二个 X 按钮
)

const (
	MA_ACTIVATE         int = iota + 1 // 激活窗口，并且不丢弃鼠标消息
	MA_ACTIVATEANDEAT                  // 激活窗口，并丢弃鼠标消息
	MA_NOACTIVATE                      // 不激活窗口，并且不丢弃鼠标消息
	MA_NOACTIVATEANDEAT                // 不激活窗口，但丢弃鼠标消息
)

const (
	INPUT_MOUSE = iota
	INPUT_KEYBOARD
	INPUT_HARDWARE
)

type Point struct {
	X, Y int32
}

type MouseInput struct {
	Dx          int32
	Dy          int32
	MouseData   uintptr
	DwFlags     uintptr
	Time        uintptr
	DwExtraInfo uintptr
}
type TagInput struct {
	Type           int
	DummyUoionName *DummyUoionName
}

type DummyUoionName struct {
	*MouseInput
	*KeyBDInput
	*HardWareInput
}

type KeyBDInput struct {
	Wvk         uint16
	Wscan       uint16
	DwFlags     uintptr
	Time        uintptr
	DwExtraInfo uintptr
}

type HardWareInput struct {
	UMsg    uintptr
	WParamL uint16
	WParamH uint16
}
