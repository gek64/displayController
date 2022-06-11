package displayController

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
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type DisplayMonitorInfo struct {
	Handle        syscall.Handle
	DeviceContext syscall.Handle
	RectAngle     RECT
}

type PhysicalMonitorInfo struct {
	Handle      syscall.Handle
	Description string
}

// freeLibrary 释放库文件(仅模块内部使用,外部使用无效)
func freeLibrary() (err error) {
	err = syscall.FreeLibrary(user32)
	if err != nil {
		return err
	}
	return syscall.FreeLibrary(dxva2)
}
