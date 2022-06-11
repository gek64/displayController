package monitor_win

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

// https://www.hattelandtechnology.com/hubfs/pdf/misc/doc101681-1_8_and_13inch_dis_ddc_control.pdf
// https://www.ddcutil.com/vcpinfo_output/
var (
	Brightness            byte = 0x10
	Contrast              byte = 0x12
	Red                   byte = 0x16
	Green                 byte = 0x18
	Blue                  byte = 0x1a
	InputSource           byte = 0x60
	Volume                byte = 0x62
	Sharpness             byte = 0x87
	ColorSaturation       byte = 0x8a
	MuteORScreenBlank     byte = 0x8d
	HorizontalFrequency   byte = 0xac
	VerticalFrequency     byte = 0xae
	DisplayTechnologyType byte = 0xb6
	DisplayUsageTime      byte = 0xc0
	PowerMode             byte = 0xd6
)

// GetVCPFeatureAndVCPFeatureReply 获取显示器VCP参数(需要使用 GetPhysicalMonitorInfo 获取到的物理显示器 Handle)
func GetVCPFeatureAndVCPFeatureReply(hPhysicalMonitor syscall.Handle, bVCPCode byte) (currentValue int, maximumValue int, err error) {
	var pvct = 0
	var pdwCurrentValue = 0
	var pdwMaximumValue = 0
	_, _, callErr := syscall.SyscallN(procGetVCPFeatureAndVCPFeatureReply,
		uintptr(hPhysicalMonitor),
		uintptr(bVCPCode),
		uintptr(unsafe.Pointer(&pvct)),
		uintptr(unsafe.Pointer(&pdwCurrentValue)),
		uintptr(unsafe.Pointer(&pdwMaximumValue)),
	)
	if callErr != 0 {
		return pdwCurrentValue, pdwMaximumValue, fmt.Errorf(callErr.Error())
	}

	return pdwCurrentValue, pdwMaximumValue, nil
}

// SetVCPFeature 设置显示器VCP参数
func SetVCPFeature(hPhysicalMonitor syscall.Handle, bVCPCode byte, value int) (err error) {
	_, _, callErr := syscall.SyscallN(procSetVCPFeature,
		uintptr(hPhysicalMonitor),
		uintptr(bVCPCode),
		uintptr(value),
	)
	if callErr != 0 {
		return fmt.Errorf(callErr.Error())
	}
	return nil
}

// BrightnessTest 亮度循环测试
func BrightnessTest(hPhysicalMonitor syscall.Handle) (err error) {

	currentBrightness, _, err := GetVCPFeatureAndVCPFeatureReply(hPhysicalMonitor, Brightness)
	if err != nil {
		return err
	}

	defer func(hPhysicalMonitor syscall.Handle, bVCPCode byte, value int) {
		err := SetVCPFeature(hPhysicalMonitor, bVCPCode, value)
		if err != nil {
			log.Panicln(err)
		}
	}(hPhysicalMonitor, Brightness, currentBrightness)

	for i := 0; i <= 100; i++ {
		fmt.Printf("set display brightness to %d%%\n", i)
		err = SetVCPFeature(hPhysicalMonitor, Brightness, i)
		if err != nil {
			return err
		}
	}

	for i := 100; i >= 0; i-- {
		fmt.Printf("set display brightness to %d%%\n", i)
		err = SetVCPFeature(hPhysicalMonitor, Brightness, i)
		if err != nil {
			return err
		}
	}

	fmt.Println("test complete")
	return nil
}
