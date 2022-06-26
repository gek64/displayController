```
██████╗ ██╗███████╗██████╗ ██╗      █████╗ ██╗   ██╗ ██████╗ ██████╗ ███╗   ██╗████████╗██████╗  ██████╗ ██╗     ██╗     ███████╗██████╗ 
██╔══██╗██║██╔════╝██╔══██╗██║     ██╔══██╗╚██╗ ██╔╝██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝██╔══██╗██╔═══██╗██║     ██║     ██╔════╝██╔══██╗
██║  ██║██║███████╗██████╔╝██║     ███████║ ╚████╔╝ ██║     ██║   ██║██╔██╗ ██║   ██║   ██████╔╝██║   ██║██║     ██║     █████╗  ██████╔╝
██║  ██║██║╚════██║██╔═══╝ ██║     ██╔══██║  ╚██╔╝  ██║     ██║   ██║██║╚██╗██║   ██║   ██╔══██╗██║   ██║██║     ██║     ██╔══╝  ██╔══██╗
██████╔╝██║███████║██║     ███████╗██║  ██║   ██║   ╚██████╗╚██████╔╝██║ ╚████║   ██║   ██║  ██║╚██████╔╝███████╗███████╗███████╗██║  ██║
╚═════╝ ╚═╝╚══════╝╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝
```
[![Go Reference](https://pkg.go.dev/badge/github.com/gek64/displayController.svg)](https://pkg.go.dev/github.com/gek64/displayController)

[中文说明](https://github.com/gek64/displayController/blob/main/README_chs.md)
- Call the low-level library of the system to access the display monitor `DDC/CI` channel and interface
- Get the display driver information, such as display driver name, current display location
- Get the value of the current parameter and the range of the current parameters, such as brightness, sharpness, contrast, red, green, blue, and other custom query parameters
- Set the value of the display parameter, such as brightness, sharpness, contrast, red, green, blue, and other custom settings parameters
- Interacting with the display, such as setting the displayed input source, controlling the displayed power mode

## Install
```shell
go get github.com/gek64/displayController
```

## go doc
- https://pkg.go.dev/github.com/gek64/displayController

## Example
```go
package main

import (
	"fmt"
	"github.com/gek64/displayController"
)

func main() {
	// Get the system display devices
	compositeMonitors, err := GetCompositeMonitors()
	if err != nil {
		log.Fatal(err)
	}

	// Travel in all display devices one by one
	for i, compositeMonitor := range compositeMonitors {
		fmt.Printf("Monitor No.%d\n", i)
		fmt.Printf("PhysicalInfo:%v\n", compositeMonitor.PhysicalInfo)
		fmt.Printf("SysInfo:%v\n", compositeMonitor.SysInfo)

		// Get the current and maximum value of the brightness parameters of the physical display
		currentValue, _, err := GetVCPFeatureAndVCPFeatureReply(compositeMonitor.PhysicalInfo.Handvaluele, Brightness)
		if err != nil {
			t.Fatal(err)
		}

		// Set the brightness of the current display to current value
		err = SetVCPFeature(compositeMonitor.PhysicalInfo.Handle, Brightness, currentValue)
		if err != nil {
			t.Fatal(err)
		}
	}
}
```

## FAQ
### What operating system does this module support?
- It only supports Windows now, and support systems such as macOS, Linux kernel system and freeBSD will be considered in the future.

### Get the display driver parameter normally, but the display monitor parameter cannot be obtained and controlled.
- This program uses the `VESA` `DDC/CI` Display communication standard protocol which release in 1998 to exchange data with physical display monitors. Most of the modern display supports and enables this feature by default, If you encounter this problem, please confirm whether the `DDC/CI` function has been opened in OSD menu, or contact your display manufacturer to get more relevant information

### What other parameters can be customized?
- Please refer to the following articles to get more custom parameters
- https://www.ddcutil.com/vcpinfo_output/
- https://www.hattelandtechnology.com/hubfs/pdf/misc/doc101681-1_8_and_13inch_dis_ddc_control.pdf

### How to find parameters supported by my own display monitor？
- If the monitor does not support a certain parameter, the error will be returned when calling the command. You can use the error information to determine whether the monitor supports a certain parameter
- You can use this tool to check which parameters that your monitor supported [ControlMyMonitor](https://www.nirsoft.net/utils/control_my_monitor.html)

## License
- **GPL-3.0 License**
- See `LICENSE` for details

## Credits
- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
