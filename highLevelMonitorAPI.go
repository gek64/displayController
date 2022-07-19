package displayController

import (
	"fmt"
	"syscall"
	"unsafe"
)

// GetMonitorBrightness 获取显示器亮度
func GetMonitorBrightness(hPhysicalMonitor syscall.Handle) (currentValue int, minimumValue int, maximumValue int, err error) {
	_, _, callErr := syscall.SyscallN(procGetMonitorBrightness, uintptr(hPhysicalMonitor), uintptr(unsafe.Pointer(&minimumValue)), uintptr(unsafe.Pointer(&currentValue)), uintptr(unsafe.Pointer(&maximumValue)))
	if callErr != 0 {
		return currentValue, minimumValue, maximumValue, fmt.Errorf(callErr.Error())
	}
	return currentValue, minimumValue, maximumValue, nil
}

// SetMonitorBrightness 设置显示器亮度
func SetMonitorBrightness(hPhysicalMonitor syscall.Handle, value int) (err error) {
	_, _, callErr := syscall.SyscallN(procSetMonitorBrightness, uintptr(hPhysicalMonitor), uintptr(value))
	if callErr != 0 {
		return fmt.Errorf(callErr.Error())
	}
	return nil
}

// GetMonitorContrast 获取显示器对比度
func GetMonitorContrast(hPhysicalMonitor syscall.Handle) (currentValue int, minimumValue int, maximumValue int, err error) {
	_, _, callErr := syscall.SyscallN(procGetMonitorContrast, uintptr(hPhysicalMonitor), uintptr(unsafe.Pointer(&minimumValue)), uintptr(unsafe.Pointer(&currentValue)), uintptr(unsafe.Pointer(&maximumValue)))
	if callErr != 0 {
		return currentValue, minimumValue, maximumValue, fmt.Errorf(callErr.Error())
	}
	return currentValue, minimumValue, maximumValue, nil
}

// SetMonitorContrast 设置显示器对比度
func SetMonitorContrast(hPhysicalMonitor syscall.Handle, value int) (err error) {
	_, _, callErr := syscall.SyscallN(procSetMonitorContrast, uintptr(hPhysicalMonitor), uintptr(value))
	if callErr != 0 {
		return fmt.Errorf(callErr.Error())
	}
	return nil
}

// RestoreMonitorFactoryColorDefaults 还原显示器出场默认颜色设置
func RestoreMonitorFactoryColorDefaults(hPhysicalMonitor syscall.Handle) (err error) {
	_, _, callErr := syscall.SyscallN(procRestoreMonitorFactoryColorDefaults, uintptr(hPhysicalMonitor))
	if callErr != 0 {
		return fmt.Errorf(callErr.Error())
	}
	return nil
}

// RestoreRestoreMonitorFactoryDefaults 还原显示器出场默认设置
func RestoreRestoreMonitorFactoryDefaults(hPhysicalMonitor syscall.Handle) (err error) {
	_, _, callErr := syscall.SyscallN(procRestoreMonitorFactoryDefaults, uintptr(hPhysicalMonitor))
	if callErr != 0 {
		return fmt.Errorf(callErr.Error())
	}
	return nil
}
