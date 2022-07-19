package displayController

import (
	"fmt"
	"testing"
)

func TestGetMonitorBrightness(t *testing.T) {
	compositeMonitors, err := GetCompositeMonitors()
	if err != nil {
		t.Fatal(err)
	}
	curr, min, max, err := GetMonitorBrightness(compositeMonitors[0].PhysicalInfo.Handle)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(curr, min, max)
}

func TestSetMonitorBrightness(t *testing.T) {
	compositeMonitors, err := GetCompositeMonitors()
	if err != nil {
		t.Fatal(err)
	}

	curr, _, _, err := GetMonitorBrightness(compositeMonitors[0].PhysicalInfo.Handle)
	if err != nil {
		t.Fatal(err)
	}

	err = SetMonitorBrightness(compositeMonitors[0].PhysicalInfo.Handle, curr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMonitorContrast(t *testing.T) {
	compositeMonitors, err := GetCompositeMonitors()
	if err != nil {
		t.Fatal(err)
	}
	curr, min, max, err := GetMonitorContrast(compositeMonitors[0].PhysicalInfo.Handle)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(curr, min, max)
}

func TestSetMonitorContrast(t *testing.T) {
	compositeMonitors, err := GetCompositeMonitors()
	if err != nil {
		t.Fatal(err)
	}

	curr, _, _, err := GetMonitorContrast(compositeMonitors[0].PhysicalInfo.Handle)
	if err != nil {
		t.Fatal(err)
	}

	err = SetMonitorContrast(compositeMonitors[0].PhysicalInfo.Handle, curr)
	if err != nil {
		t.Fatal(err)
	}
}
