package monitor_win

import (
	"syscall"
)

var (
	user32, _                                      = syscall.LoadLibrary("User32.dll")
	dxva2, _                                       = syscall.LoadLibrary("dxva2.dll")
	procEnumDisplayMonitors, _                     = syscall.GetProcAddress(user32, "EnumDisplayMonitors")
	procGetNumberOfPhysicalMonitorsFromHMONITOR, _ = syscall.GetProcAddress(dxva2, "GetNumberOfPhysicalMonitorsFromHMONITOR")
	procGetPhysicalMonitorsFromHMONITOR, _         = syscall.GetProcAddress(dxva2, "GetPhysicalMonitorsFromHMONITOR")
	procGetVCPFeatureAndVCPFeatureReply, _         = syscall.GetProcAddress(dxva2, "GetVCPFeatureAndVCPFeatureReply")
	procSetVCPFeature, _                           = syscall.GetProcAddress(dxva2, "SetVCPFeature")
)

type RECT struct {
	left   int32
	top    int32
	right  int32
	bottom int32
}

type DisplayMonitorInfo struct {
	handle        syscall.Handle
	deviceContext syscall.Handle
	rectAngle     RECT
}

type PhysicalMonitorInfo struct {
	handle      syscall.Handle
	description string
}

// FreeLibrary 释放库文件
func FreeLibrary() (err error) {
	err = syscall.FreeLibrary(user32)
	if err != nil {
		return err
	}
	return syscall.FreeLibrary(dxva2)
}
