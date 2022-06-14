```
██████╗ ██╗███████╗██████╗ ██╗      █████╗ ██╗   ██╗ ██████╗ ██████╗ ███╗   ██╗████████╗██████╗  ██████╗ ██╗     ██╗     ███████╗██████╗ 
██╔══██╗██║██╔════╝██╔══██╗██║     ██╔══██╗╚██╗ ██╔╝██╔════╝██╔═══██╗████╗  ██║╚══██╔══╝██╔══██╗██╔═══██╗██║     ██║     ██╔════╝██╔══██╗
██║  ██║██║███████╗██████╔╝██║     ███████║ ╚████╔╝ ██║     ██║   ██║██╔██╗ ██║   ██║   ██████╔╝██║   ██║██║     ██║     █████╗  ██████╔╝
██║  ██║██║╚════██║██╔═══╝ ██║     ██╔══██║  ╚██╔╝  ██║     ██║   ██║██║╚██╗██║   ██║   ██╔══██╗██║   ██║██║     ██║     ██╔══╝  ██╔══██╗
██████╔╝██║███████║██║     ███████╗██║  ██║   ██║   ╚██████╗╚██████╔╝██║ ╚████║   ██║   ██║  ██║╚██████╔╝███████╗███████╗███████╗██║  ██║
╚═════╝ ╚═╝╚══════╝╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝
```
- 调用系统底层库来访问显示器`DDC/CI`通道与接口
- 获取显示器驱动信息，例如显示器驱动名称、当前显示器位置
- 获取显示器当前参数的值及参数的取值范围，例如亮度、锐利度、对比度、红、绿、蓝及其他可自定义查询参数
- 设置显示器参数的值，例如亮度、锐利度、对比度、红、绿、蓝及其他可自定义设置参数
- 与显示器进行交互，例如设置显示的输入源，控制显示的电源模式

## 安装
```shell
go get github.com/gek64/displayController
```

## go doc
- https://pkg.go.dev/github.com/gek64/displayController

## 例子
```go
package main

import (
	"fmt"
	"github.com/gek64/displayController"
)

func main() {
	// 获取系统显示设备
	monitors, _ := displayController.GetAllMonitors()

	// 在系统所有显示设备中逐个遍历
	for i, monitor := range monitors {
		// 获取物理显示器
		physicalMonitor, _ := displayController.GetPhysicalMonitor(monitor.Handle)

		// 获取物理显示器的亮度参数的当前值及最大值
		currentValue, maximumValue, _ := displayController.GetVCPFeatureAndVCPFeatureReply(physicalMonitor.Handle, displayController.Brightness)

		// 将当前显示器的亮度设置为100
		_ = displayController.SetVCPFeature(physicalMonitor.Handle, displayController.Brightness, 100)

		fmt.Printf("Display Monitor: %d,Driver Name: %s,Current Brightness: %d,Maximum Brightness: %d\n", i, physicalMonitor.Description, currentValue, maximumValue)
	}
}
```

## 常见问题
### 这个模块支持哪些系统?
- 这个模块当前只支持 windows, 未来会考虑支持macOS、linux内核系统、freebsd等系统

### 可以正常获取显示驱动程序参数，但是无法获得和控制显示显示监视器参数。
- 本程序使用`vesa`在`2004`年定义的`DDC/CI`显示器通讯标准协议与显示器进行数据交换，绝大部分的现代显示器都默认支持并启用了这一项功能，但部分显示器的制造商可能因为某些特定因素的考量而默认关闭了这个选项，请确认显示器`OSD`菜单中是否已经开启了`DDC/CI`功能选项，或与您的显示器制造商联系获取更多有关的信息

### 能自定义查询、设置的参数除了已在库文件中定义的还有哪些？
- 请参考以下两篇文章来获取更多的自定义选项
- https://www.ddcutil.com/vcpinfo_output/
- https://www.hattelandtechnology.com/hubfs/pdf/misc/doc101681-1_8_and_13inch_dis_ddc_control.pdf

### 如何查找我自己的显示监视器支持的参数？
- 如果监视器不支持某个参数，则在调用命令时将返回错误，您可以使用错误信息来确定监视器是否支持某个参数
- 可以使用这个工具来检查你的显示支持的参数[ControlMyMonitor](https://www.nirsoft.net/utils/control_my_monitor.html)

## 许可证
- **GPL-3.0 License**
- 查看 `LICENSE` 获取详细内容

## 致谢
- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
