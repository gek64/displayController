package displayController

import (
	"fmt"
	"syscall"
	"unsafe"
)

type CompositeMonitorInfo struct {
	PhysicalInfo PhysicalMonitorInfo
	SysInfo      SystemMonitorInfo
}

// GetCompositeMonitors 获取复合显示器信息
func GetCompositeMonitors() (monitors []CompositeMonitorInfo, err error) {
	// 获取系统显示器信息
	systemMonitors, err := GetSystemMonitors()
	if err != nil {
		return nil, err
	}

	for _, sysMonitor := range systemMonitors {
		// 获取物理显示器信息
		physicalMonitor, err := GetPhysicalMonitor(sysMonitor.Handle)
		if err != nil {
			continue
		}
		// 拼接复合显示器信息
		monitors = append(monitors, CompositeMonitorInfo{
			PhysicalInfo: physicalMonitor,
			SysInfo:      sysMonitor,
		})
	}

	return monitors, nil
}

// GetSystemMonitors 获取所有屏幕设备信息
func GetSystemMonitors() (info []SystemMonitorInfo, err error) {
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-monitorenumproc
	var fnCallback = func(hMonitor syscall.Handle, hdc syscall.Handle, rect *RECT, lParam uintptr) int {
		info = append(info, SystemMonitorInfo{Handle: hMonitor, DeviceContext: hdc, RectAngle: *rect})
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
		return info, fmt.Errorf(callErr.Error())
	}

	return info, nil
}

// GetMonitorNumberFromHandle 获取显示器句柄下的显示器数量
func GetMonitorNumberFromHandle(hMonitor syscall.Handle) (number int32, err error) {
	_, _, callErr := syscall.SyscallN(procGetNumberOfPhysicalMonitorsFromHMONITOR,
		uintptr(hMonitor),
		uintptr(number),
	)
	if callErr != 0 {
		return 0, fmt.Errorf(callErr.Error())
	}

	return number, nil
}

// GetPhysicalMonitor 获取物理显示器信息
func GetPhysicalMonitor(hMonitor syscall.Handle) (info PhysicalMonitorInfo, err error) {
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

	return PhysicalMonitorInfo{Handle: syscall.Handle(bytes[0]), Description: string(newBytes)}, nil
}
