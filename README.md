```
        ▄▄   ▄▄                     ▄▄                                                                           ▄▄    ▄▄
      ▀███   ██                   ▀███                      ▄▄█▀▀▀█▄█                     ██                    ▀███  ▀███
        ██                          ██                    ▄██▀     ▀█                     ██                      ██    ██
   ▄█▀▀███ ▀███  ▄██▀███████████▄   ██  ▄█▀██▄ ▀██▀   ▀██ ▀█▀       ▀ ▄██▀██▄▀████████▄ ██████▀███▄███  ▄██▀██▄   ██    ██   ▄▄█▀██▀███▄███
 ▄██    ██   ██  ██   ▀▀ ██   ▀██   ██ ██   ██   ██   ▄█  ██         ██▀   ▀██ ██    ██   ██    ██▀ ▀▀ ██▀   ▀██  ██    ██  ▄█▀   ██ ██▀ ▀▀
 ███    ██   ██  ▀█████▄ ██    ██   ██  ▄█████    ██ ▄█   ██▄        ██     ██ ██    ██   ██    ██     ██     ██  ██    ██  ██▀▀▀▀▀▀ ██
 ▀██    ██   ██  █▄   ██ ██   ▄██   ██ ██   ██     ███    ▀██▄     ▄▀██▄   ▄██ ██    ██   ██    ██     ██▄   ▄██  ██    ██  ██▄    ▄ ██
  ▀████▀███▄████▄██████▀ ██████▀  ▄████▄████▀██▄   ▄█       ▀▀█████▀  ▀█████▀▄████  ████▄ ▀████████▄    ▀█████▀ ▄████▄▄████▄ ▀█████▀████▄
                         ██                      ▄█
                       ▄████▄                  ██▀
```
[中文说明](#)
- 1
- 2
- 3

## Usage
```go
package main

import (
	"fmt"
	"github.com/gek64/displayController"
)

func main() {
	// get all monitors
	monitors, _ := displayController.GetAllMonitors()

	for i, monitor := range monitors {
		// get physical monitor
		physicalMonitor, _ := displayController.GetPhysicalMonitor(monitor.Handle)

		// get physical monitor current brightness value, maximum brightness value
		currentValue, maximumValue, _ := displayController.GetVCPFeatureAndVCPFeatureReply(physicalMonitor.Handle, displayController.Brightness)

		// set monitor brightness to 100
		_ = displayController.SetVCPFeature(physicalMonitor.Handle, displayController.Brightness, 100)

		fmt.Printf("Display Monitor: %d,Driver Name: %s,Current Brightness: %d,Maximum Brightness: %d\n", i, physicalMonitor.Description, currentValue, maximumValue)
	}
}
```

## Install
```shell
go get github.com/gek64/displayController
```

## FAQ
### 1
- 1

### 2
- 2

### 3
- 3

## License
- **GPL-3.0 License**
- See `LICENSE` for details

## Credits
- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
