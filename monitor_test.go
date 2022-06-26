package displayController

import (
	"fmt"
	"testing"
)

func TestGetCompositeMonitors(t *testing.T) {
	compositeMonitors, err := GetCompositeMonitors()
	if err != nil {
		t.Fatal(err)
	}

	for i, compositeMonitor := range compositeMonitors {
		fmt.Printf("Monitor No.%d\n", i)
		fmt.Printf("PhysicalInfo:%v\n", compositeMonitor.PhysicalInfo)
		fmt.Printf("SysInfo:%v\n", compositeMonitor.SysInfo)

		currentValue, _, err := GetVCPFeatureAndVCPFeatureReply(compositeMonitor.PhysicalInfo.Handle, Brightness)
		if err != nil {
			t.Fatal(err)
		}

		err = SetVCPFeature(compositeMonitor.PhysicalInfo.Handle, Brightness, currentValue)
		if err != nil {
			t.Fatal(err)
		}
	}

	fmt.Println("GetCompositeMonitors test complete")
}
