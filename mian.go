package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	defer func(handle syscall.Handle) {
		err := syscall.FreeLibrary(handle)
		if err != nil {
			log.Panicln(err)
		}
	}(user32)
	defer func(handle syscall.Handle) {
		err := syscall.FreeLibrary(handle)
		if err != nil {
			log.Panicln(err)
		}
	}(dxva2)

	monitors, err := getAllMonitors()
	if err != nil {
		log.Fatalln(err)
	}

	for i, monitor := range monitors {
		fmt.Println("monitor", i, monitor)

		physicalMonitorInfo, err := getPhysicalMonitorInfo(monitor.handle)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(getVCPFeatureAndVCPFeatureReply(physicalMonitorInfo.handle, 0x10))
	}
}
