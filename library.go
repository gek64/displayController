package displayController

import (
	"syscall"
)

var (
	// load library
	user32, _ = syscall.LoadLibrary("User32.dll")
	dxva2, _  = syscall.LoadLibrary("dxva2.dll")

	// https://docs.microsoft.com/en-us/windows/win32/api/_monitor/
	// physical monitor api
	procEnumDisplayMonitors, _                     = syscall.GetProcAddress(user32, "EnumDisplayMonitors")
	procGetNumberOfPhysicalMonitorsFromHMONITOR, _ = syscall.GetProcAddress(dxva2, "GetNumberOfPhysicalMonitorsFromHMONITOR")

	// low level monitor api
	procGetPhysicalMonitorsFromHMONITOR, _ = syscall.GetProcAddress(dxva2, "GetPhysicalMonitorsFromHMONITOR")
	procGetVCPFeatureAndVCPFeatureReply, _ = syscall.GetProcAddress(dxva2, "GetVCPFeatureAndVCPFeatureReply")
	procSetVCPFeature, _                   = syscall.GetProcAddress(dxva2, "SetVCPFeature")

	// high level monitor api
	// get
	procGetMonitorBrightness, _          = syscall.GetProcAddress(dxva2, "GetMonitorBrightness")
	procGetMonitorCapabilities, _        = syscall.GetProcAddress(dxva2, "GetMonitorCapabilities")
	procGetMonitorColorTemperature, _    = syscall.GetProcAddress(dxva2, "GetMonitorColorTemperature")
	procGetMonitorContrast, _            = syscall.GetProcAddress(dxva2, "GetMonitorContrast")
	procGetMonitorDisplayAreaPosition, _ = syscall.GetProcAddress(dxva2, "GetMonitorDisplayAreaPosition")
	procGetMonitorDisplayAreaSize, _     = syscall.GetProcAddress(dxva2, "GetMonitorDisplayAreaSize")
	procGetMonitorRedGreenOrBlueDrive, _ = syscall.GetProcAddress(dxva2, "GetMonitorRedGreenOrBlueDrive")
	procGetMonitorRedGreenOrBlueGain, _  = syscall.GetProcAddress(dxva2, "GetMonitorRedGreenOrBlueGain")
	procGetMonitorTechnologyType, _      = syscall.GetProcAddress(dxva2, "GetMonitorTechnologyType")
	// set
	procSetMonitorBrightness, _          = syscall.GetProcAddress(dxva2, "SetMonitorBrightness")
	procSetMonitorColorTemperature, _    = syscall.GetProcAddress(dxva2, "SetMonitorColorTemperature")
	procSetMonitorContrast, _            = syscall.GetProcAddress(dxva2, "SetMonitorContrast")
	procSetMonitorDisplayAreaPosition, _ = syscall.GetProcAddress(dxva2, "SetMonitorDisplayAreaPosition")
	procSetMonitorDisplayAreaSize, _     = syscall.GetProcAddress(dxva2, "SetMonitorDisplayAreaSize")
	procSetMonitorRedGreenOrBlueDrive, _ = syscall.GetProcAddress(dxva2, "SetMonitorRedGreenOrBlueDrive")
	procSetMonitorRedGreenOrBlueGain, _  = syscall.GetProcAddress(dxva2, "SetMonitorRedGreenOrBlueGain")
	// other
	procRestoreMonitorFactoryColorDefaults, _ = syscall.GetProcAddress(dxva2, "RestoreMonitorFactoryColorDefaults")
	procRestoreMonitorFactoryDefaults, _      = syscall.GetProcAddress(dxva2, "RestoreMonitorFactoryDefaults")
	procSaveCurrentMonitorSettings, _         = syscall.GetProcAddress(dxva2, "SaveCurrentMonitorSettings")
	procDegaussMonitor, _                     = syscall.GetProcAddress(dxva2, "DegaussMonitor")
)

// loadLibraryFunc
func loadLibraryFunc(library string, fun string) (lib syscall.Handle, proc uintptr, err error) {
	// load library
	lib, err = syscall.LoadLibrary(library)
	if err != nil {
		return 0, 0, err
	}

	// load func
	procFunc, err := syscall.GetProcAddress(lib, fun)
	if err != nil {
		return 0, 0, err
	}

	return lib, procFunc, err
}

// freeLibrary
func freeLibrary(library syscall.Handle) (err error) {
	return syscall.FreeLibrary(library)
}
