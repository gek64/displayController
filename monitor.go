package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

// 获取所有屏幕设备信息
func getAllMonitors() (monitors []DisplayMonitorInfo, err error) {
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-monitorenumproc
	var fnCallback = func(hMonitor syscall.Handle, hdc syscall.Handle, rect *RECT, lParam uintptr) int {
		monitors = append(monitors, DisplayMonitorInfo{handle: hMonitor, deviceContext: hdc, rectAngle: *rect})
		// 继续枚举下一个显示器,1代表true
		return 1
	}

	// 两者结果一致 uintptr(unsafe.Pointer(nil)) or uintptr(syscall.Handle(0))
	_, _, callErr := syscall.SyscallN(procEnumDisplayMonitors,
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(nil)),
		syscall.NewCallback(fnCallback),
		uintptr(unsafe.Pointer(nil)),
	)
	if callErr != 0 {
		return monitors, fmt.Errorf(callErr.Error())
	}

	return monitors, nil
}

// 获取显示器句柄下的显示器数量
func getMonitorNumberFromHandle(hMonitor syscall.Handle) (number int32, err error) {
	_, _, callErr := syscall.SyscallN(procGetNumberOfPhysicalMonitorsFromHMONITOR,
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(&number)),
	)
	if callErr != 0 {
		return 0, fmt.Errorf(callErr.Error())
	}

	return number, nil
}

// 获取物理显示器信息
func getPhysicalMonitorInfo(hMonitor syscall.Handle) (info PhysicalMonitorInfo, err error) {
	bytes := make([]byte, 256)
	_, _, callErr := syscall.SyscallN(procGetPhysicalMonitorsFromHMONITOR,
		uintptr(hMonitor),
		uintptr(1),
		uintptr(unsafe.Pointer(&bytes[0])),
	)
	if callErr != 0 {
		return PhysicalMonitorInfo{}, fmt.Errorf(callErr.Error())
	}

	// 第8位以后才是显示器描述信息
	// 每个字母用0隔开,需要重新整理
	var newBytes []byte
	for _, b := range bytes[8:] {
		if b != 0 {
			newBytes = append(newBytes, b)
		}
	}

	return PhysicalMonitorInfo{handle: syscall.Handle(bytes[0]), description: string(newBytes)}, nil
}

// 获取显示器VCP参数(需要使用 getPhysicalMonitorInfo 获取到的物理显示器 handle)
func getVCPFeatureAndVCPFeatureReply(hPhysicalMonitor syscall.Handle, bVCPCode int32) (pvct int, pdwCurrentValue int, pdwMaximumValue int, err error) {
	_, _, callErr := syscall.SyscallN(procGetVCPFeatureAndVCPFeatureReply,
		uintptr(hPhysicalMonitor),
		uintptr(bVCPCode),
		uintptr(unsafe.Pointer(&pvct)),
		uintptr(unsafe.Pointer(&pdwCurrentValue)),
		uintptr(unsafe.Pointer(&pdwMaximumValue)),
	)
	if callErr != 0 {
		return 0, 0, 0, fmt.Errorf(callErr.Error())
	}

	return pvct, pdwCurrentValue, pdwMaximumValue, nil
}
