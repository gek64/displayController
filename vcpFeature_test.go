package displayController

import (
	"fmt"
	"testing"
)

func TestBrightnessTest(t *testing.T) {
	// 获取所有显示设备(包括虚拟设备)
	monitors, err := GetSystemMonitors()
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

		fmt.Printf("Display Monitor: %d,Driver Name: %s,Current Brightness: %d,Maximum Brightness: %d\n", i, physicalMonitor.Description, currentValue, maximumValue)
		fmt.Println("The brightness test will be performed, the brightness of the display will change from 0% to 100%, then from 100% to 0%\nAfter the test is completed, the original brightness will be restored")

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

		fmt.Println("Brightness test is completed")
	}
}
