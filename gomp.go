package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"

#include "main.c"
#include "player.c"
#include "vehicle.c"
#include "event.c"
#include "core.c"
#include "all.c"

#endif
*/
import "C"
import (
	"sync"
	"unsafe"
)

var onReadyFunc *func()
var onResetFunc *func()
var onFreeFunc *func()

//export onReady
func onReady() {
	if onReadyFunc == nil {
		return
	}

	(*onReadyFunc)()
}

//export onReset
func onReset() {
	if onResetFunc == nil {
		return
	}

	(*onResetFunc)()
}

//export onFree
func onFree() {
	if onFreeFunc == nil {
		return
	}

	(*onFreeFunc)()
}

var apiMutex sync.Mutex

func SetComponentVersion(major, minor, patch uint8, prerel uint16) {
	C.setComponentVersion(C.uint8_t(major), C.uint8_t(minor), C.uint8_t(patch), C.uint16_t(prerel))
}

func SetComponentName(name string) {
	if name == "" {
		return
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.setComponentName(cName)
}

func OnReady(handler func()) {
	if onReadyFunc != nil {
		panic("OnReady handler already registered")
	}

	onReadyFunc = &handler
}

func OnReset(handler func()) {
	if onResetFunc != nil {
		panic("OnReset handler already registered")
	}

	onResetFunc = &handler
}

func OnFree(handler func()) {
	if onFreeFunc != nil {
		panic("OnFree handler already registered")
	}

	onFreeFunc = &handler
}
