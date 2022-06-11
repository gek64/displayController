package displayController

import (
	"fmt"
	"testing"
)

func TestBrightnessTest(t *testing.T) {
	// 获取所有显示设备(包括虚拟设备)
	monitors, err := GetAllMonitors()
	if err != nil {
		t.Fatal(err)
	}

	for i, monitor := range monitors {
		// 获取物理显示器设备
		physicalMonitor, err := GetPhysicalMonitor(monitor.Handle)
		if err != nil {
			t.Fatal(err)
		}

		// 获取当前显示器的亮度值及亮度最大值
		currentValue, maximumValue, err := GetVCPFeatureAndVCPFeatureReply(physicalMonitor.Handle, Brightness)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("显示器 %d,驱动名称 %s,当前亮度为 %d,最大亮度为 %d\n", i, physicalMonitor.Description, currentValue, maximumValue)
		fmt.Println("将进行亮度测试,显示器亮度会从 0% 变化到 100%,之后又从 100% 变化到 0% ,测试完成后会恢复原有亮度")

		// 将当前显示器亮度设置为原有值
		err = SetVCPFeature(physicalMonitor.Handle, Brightness, currentValue)
		if err != nil {
			t.Fatal(err)
		}

		// 对当前显示器进行亮度测试
		err = BrightnessTest(physicalMonitor.Handle)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("亮度测试完成")
	}
}
