package main

import "fmt"

func main() {
	monitors, _ := getAllDisplayMonitors()

	for i, monitor := range monitors {
		fmt.Printf("Monitor:%d\n", i)
		fmt.Printf("Handle:%v\n", monitor.handle)
		fmt.Printf("deviceContext:%v\n", monitor.deviceContext)
		fmt.Printf("Rectangle:%v\n", monitor.rectAngle)
	}
}
