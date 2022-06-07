package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func getAllDisplayMonitors() (monitors []DisplayMonitorInfo, err error) {
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-monitorenumproc
	var fnCallback = func(handleMonitor syscall.Handle, hdc syscall.Handle, rect *Rect, lParam uintptr) int {
		monitors = append(monitors, DisplayMonitorInfo{handle: handleMonitor, deviceContext: hdc, rectAngle: *rect})
		// 继续枚举下一个显示器,1代表true
		return 1
	}

	// 两者结果一致 uintptr(unsafe.Pointer(nil)) or uintptr(syscall.Handle(0))
	r1, _, errno := syscall.SyscallN(procEnumDisplayMonitors, uintptr(unsafe.Pointer(nil)), uintptr(unsafe.Pointer(nil)), syscall.NewCallback(fnCallback), uintptr(unsafe.Pointer(nil)))
	if errno != 0 {
		return monitors, fmt.Errorf("syscall.SyscallN() error occur when use getAllDisplayMonitors")
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors#return-value
	if r1 == 0 {
		return monitors, fmt.Errorf("EnumDisplayMonitors function fails")
	}
	return monitors, nil
}
