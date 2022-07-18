package displayController

import (
	"fmt"
	"syscall"
	"unsafe"
)

func GetMonitorBrightness(hPhysicalMonitor syscall.Handle) (currentValue int, minimumValue int, maximumValue int, err error) {
	_, _, callErr := syscall.SyscallN(procGetMonitorBrightness, uintptr(hPhysicalMonitor), uintptr(unsafe.Pointer(&minimumValue)), uintptr(unsafe.Pointer(&currentValue)), uintptr(unsafe.Pointer(&maximumValue)))
	if callErr != 0 {
		return currentValue, minimumValue, maximumValue, fmt.Errorf(callErr.Error())
	}
	return currentValue, minimumValue, maximumValue, nil
}

func SetMonitorBrightness(hPhysicalMonitor syscall.Handle, value int) (err error) {
	_, _, callErr := syscall.SyscallN(procSetMonitorBrightness, uintptr(value))
	if callErr != 0 {
		return fmt.Errorf(callErr.Error())
	}
	return nil
}
