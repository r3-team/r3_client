//go:build windows

package open

import (
	"syscall"
	"unsafe"
)

func chooseApp(filePath string) error {

	// https://docs.microsoft.com/en-us/windows/win32/api/shlobj_core/nf-shlobj_core-shopenwithdialog
	dllShell32 := syscall.NewLazyDLL("shell32.dll")
	procOpenWith := dllShell32.NewProc("SHOpenWithDialog")

	// https://docs.microsoft.com/en-us/windows/win32/api/shlobj_core/ns-shlobj_core-openasinfo
	type OPENASINFO struct {
		pcszFile    *uint16
		pcszClass   *uint16
		oaifInFlags uint32
	}
	filePathPtr, err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		return err
	}
	var data = OPENASINFO{
		pcszFile:    filePathPtr,
		oaifInFlags: 4,
	}
	if _, _, err := procOpenWith.Call(0, uintptr(unsafe.Pointer(&data))); err != nil {
		return err
	}
	return nil
}
