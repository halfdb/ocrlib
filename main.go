package main

import (
	//#include <stdlib.h>
	"C"
	"sync"
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

var mutex = sync.Mutex{}
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
	mutex.Lock()
	defer mutex.Unlock()
	return cResult
}

func setResult(res string) {
	mutex.Lock()
	defer mutex.Unlock()
	result = res
	C.free(unsafe.Pointer(cResult))
	cResult = C.CString(res)
}
