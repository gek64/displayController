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
	curr, min, max, err := displayController.GetMonitorBrightness(compositeMonitors[0].PhysicalInfo.Handle)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(curr, min, max)
}
