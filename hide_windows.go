package systray

import (
	. "log"
	"syscall"
)

var (
	GetConsoleWindow = syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	ShowWindow       = syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
)

const (
	SW_HIDE = 0
	SW_SHOW = 5
)

func HideConsole() {
	handle, _, _ := GetConsoleWindow.Call()
	ShowWindow.Call(handle, SW_HIDE)
}

func ShowConsole() {
	handle, _, _ := GetConsoleWindow.Call()
	ShowWindow.Call(handle, SW_SHOW)

}

func Syscall3ByName(library syscall.Handle, funcName string, nargs, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	c, e := syscall.GetProcAddress(library, funcName)
	if e != nil {
		Println(e)
		return
	}
	return syscall.Syscall(c, nargs, a1, a2, a3)

}
