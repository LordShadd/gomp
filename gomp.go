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

#endif
*/
import "C"
import (
	"sync"
	"unsafe"
)

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
