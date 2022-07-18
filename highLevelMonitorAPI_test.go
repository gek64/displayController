package displayController

import (
	"fmt"
	"log"
)

func ExampleGetMonitorBrightness() {
	compositeMonitors, err := GetCompositeMonitors()
	if err != nil {
		log.Fatalln(err)
	}
	curr, min, max, err := GetMonitorBrightness(compositeMonitors[0].PhysicalInfo.Handle)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(curr, min, max)
}

func ExampleSetMonitorBrightness() {
	compositeMonitors, err := GetCompositeMonitors()
	if err != nil {
		log.Fatalln(err)
	}
	err = SetMonitorBrightness(compositeMonitors[0].PhysicalInfo.Handle, 100)
	if err != nil {
		log.Fatalln(err)
	}
}
