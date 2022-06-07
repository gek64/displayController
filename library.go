package main

import "syscall"

var (
	user32, _                                      = syscall.LoadLibrary("User32.dll")
	dxva2, _                                       = syscall.LoadLibrary("dxva2.dll")
	procEnumDisplayMonitors, _                     = syscall.GetProcAddress(user32, "EnumDisplayMonitors")
	procGetNumberOfPhysicalMonitorsFromHMONITOR, _ = syscall.GetProcAddress(dxva2, "GetNumberOfPhysicalMonitorsFromHMONITOR")
	procGetPhysicalMonitorsFromHMONITOR, _         = syscall.GetProcAddress(dxva2, "GetPhysicalMonitorsFromHMONITOR")
)

type Rect struct {
	left   int32
	top    int32
	right  int32
	bottom int32
}

type DisplayMonitorInfo struct {
	handle        syscall.Handle
	deviceContext syscall.Handle
	rectAngle     Rect
}

type PhysicalMonitor struct {
	hPhysicalMonitor             syscall.Handle
	szPhysicalMonitorDescription string
}
