package main

import (
	"fmt"
	"github.com/gek64/displayController"
	"log"
)

func main() {
	compositeMonitors, err := displayController.GetCompositeMonitors()
	if err != nil {
		log.Fatalln(err)
	}
	for _, monitor := range compositeMonitors {
		fmt.Println(displayController.RestoreRestoreMonitorFactoryDefaults(monitor.PhysicalInfo.Handle))
	}
}
