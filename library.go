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
