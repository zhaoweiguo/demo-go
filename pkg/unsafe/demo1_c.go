package main

/*
#include <stdio.h>
*/

import "C"
import "unsafe"

func main() {
	cstr := C.CString("hello world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
