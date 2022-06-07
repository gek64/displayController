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

	monitors, _ := getAllDisplayMonitors()

	for i, monitor := range monitors {
		fmt.Printf("Monitor:%d\n", i)
		fmt.Printf("Handle:%v\n", monitor.handle)
		fmt.Printf("DeviceContext:%v\n", monitor.deviceContext)
		fmt.Printf("Rectangle:%v\n", monitor.rectAngle)

		monitorDescription, err := getPhysicalMonitorDescription(monitor.handle)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Description:%s\n", monitorDescription)
	}
}
