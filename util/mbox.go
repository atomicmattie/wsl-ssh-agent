//go:build windows
// +build windows

package util

import (
	"log"
	"unsafe"

	"golang.org/x/sys/windows"
)

// MsgType specifies how message box will look.
type MsgType uint32

// Actual values.
const (
	MsgError       MsgType = 0x00000010
	MsgExclamation MsgType = 0x00000030
	MsgInformation MsgType = 0x00000040
)

// ShowOKMessage shows MB_OK message box.
func ShowOKMessage(t MsgType, title, text string) {

	log.Print(text)

	var (
		mod  = windows.NewLazySystemDLL("user32")
		proc = mod.NewProc("MessageBoxW")
		mb   = 0x00000000 + t
	)
	_, _, _ = proc.Call(0,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(title))),
		uintptr(mb))
}
