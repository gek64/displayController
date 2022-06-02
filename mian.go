package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	user32, err := syscall.LoadLibrary("user32.dll")
	if err != nil {
		log.Fatalln(err)
	}

	defer func(handle syscall.Handle) {
		err := syscall.FreeLibrary(handle)
		if err != nil {
			log.Fatalln(err)
		}
	}(user32)

	//https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfrompoint
	procAddress, err := syscall.GetProcAddress(user32, "MonitorFromPoint")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(procAddress)

}
