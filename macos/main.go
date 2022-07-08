package main

/*
#include <IOKit/IOKitLib.h>

*/
import "C"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
