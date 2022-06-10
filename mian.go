package main

import (
	"log"
)

func main() {
	defer FreeLibrary()

	monitors, err := GetAllMonitors()
	if err != nil {
		log.Fatalln(err)
	}

	physicalMonitorInfo, err := GetPhysicalMonitor(monitors[0].handle)
	if err != nil {
		log.Fatalln(err)
	}

	err = BrightnessTest(physicalMonitorInfo.handle)
	if err != nil {
		log.Fatalln(err)
	}
}
