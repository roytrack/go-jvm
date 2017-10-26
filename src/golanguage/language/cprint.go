package main

import (
	"unsafe"
	"C"
)

func main() {

	cstr := C.CString("hello world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))

}
