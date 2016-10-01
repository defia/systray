package systray

import (
	. "log"
	"syscall"
)

var kernel32, user32 syscall.Handle

func init() {
	var err error
	kernel32, err = syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		Fatal(err)
	}
	user32, err = syscall.LoadLibrary("user32.dll")
	if err != nil {
		Fatal(err)
	}
}

var (
	SW_HIDE = uintptr(0)
	SW_SHOW = uintptr(5)
)

func HideConsole() {
	handle, _, _ := Syscall3ByName(kernel32, "GetConsoleWindow", 0, 0, 0, 0)

	Syscall3ByName(user32, "ShowWindow", 2, handle, SW_HIDE, 0)

}

func ShowConsole() {
	handle, _, _ := Syscall3ByName(kernel32, "GetConsoleWindow", 0, 0, 0, 0)

	Syscall3ByName(user32, "ShowWindow", 2, handle, SW_SHOW, 0)

}

func Syscall3ByName(library syscall.Handle, funcName string, nargs, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	c, e := syscall.GetProcAddress(library, funcName)
	if e != nil {
		Println(e)
		return
	}
	return syscall.Syscall(c, nargs, a1, a2, a3)

}
