package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func getAllDisplayMonitors() (monitors []DisplayMonitorInfo, err error) {
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-monitorenumproc
	var fnCallback = func(hMonitor syscall.Handle, hdc syscall.Handle, rect *Rect, lParam uintptr) int {
		monitors = append(monitors, DisplayMonitorInfo{handle: hMonitor, deviceContext: hdc, rectAngle: *rect})
		// 继续枚举下一个显示器,1代表true
		return 1
	}

	// 两者结果一致 uintptr(unsafe.Pointer(nil)) or uintptr(syscall.Handle(0))
	r1, _, callErr := syscall.SyscallN(procEnumDisplayMonitors, uintptr(unsafe.Pointer(nil)), uintptr(unsafe.Pointer(nil)), syscall.NewCallback(fnCallback), uintptr(unsafe.Pointer(nil)))
	if callErr != 0 {
		return monitors, fmt.Errorf("call getPhysicalMonitor %s", callErr.Error())
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors#return-value
	if r1 == 0 {
		return monitors, fmt.Errorf("EnumDisplayMonitors function fails")
	}
	return monitors, nil
}

func getPhysicalMonitorNumber(hMonitor syscall.Handle) (number int32, err error) {
	r1, _, callErr := syscall.SyscallN(procGetNumberOfPhysicalMonitorsFromHMONITOR, uintptr(hMonitor), uintptr(unsafe.Pointer(&number)))
	if callErr != 0 {
		return 0, fmt.Errorf("call getPhysicalMonitor %s", callErr.Error())
	}

	if r1 != 1 {
		return 0, fmt.Errorf("GetNumberOfPhysicalMonitorsFromHMONITOR function fails")
	}

	return number, nil
}

func getPhysicalMonitorDescription(hMonitor syscall.Handle) (monitorDescription string, err error) {
	bytes := make([]byte, 256)
	r1, _, callErr := syscall.SyscallN(procGetPhysicalMonitorsFromHMONITOR, uintptr(hMonitor), uintptr(1), uintptr(unsafe.Pointer(&bytes[0])))
	if callErr != 0 {
		return "", fmt.Errorf("call getPhysicalMonitor %s", callErr.Error())
	}

	if r1 != 1 {
		return "", fmt.Errorf("GetPhysicalMonitorsFromHMONITOR function fails")
	}

	// 第8位以后才是显示器描述信息
	// 每个字母用0隔开,需要重新整理
	var newBytes []byte
	for _, b := range bytes[8:] {
		if b != 0 {
			newBytes = append(newBytes, b)
		}
	}

	return string(newBytes), nil
}
