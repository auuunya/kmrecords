package kmrecords

import "syscall"

type MSG struct {
	Hwnd     syscall.Handle
	Message  uint32
	WParam   uintptr
	LParam   uintptr
	Time     uint32
	Pt       *Point
	LPrivate uint32
}
