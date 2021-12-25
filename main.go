package main

import (
	//#include <stdlib.h>
	"C"
	"time"
	"unsafe"
)

func main() {
	read(737, 473, 307, 196)
	for result == "" {
		time.Sleep(100 * time.Millisecond)
	}
	print(result)
}

var result = ""
var cResult *C.char = nil

//export read
func read(x, y, w, h int) {
	img := shot(x, y, w, h)
	file := convert(img)
	go recogniseAndSet(file)
}

//export lastResult
func lastResult() *C.char {
	return cResult
}

func setResult(res string) {
	result = res
	C.free(unsafe.Pointer(cResult))
	cResult = C.CString(res)
}
