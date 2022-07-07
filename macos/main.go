package main

import (
	"fmt"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(syscall.SYS___MAC_GET_PID, 0, 0, 0)

	fmt.Println(pid)
}
