package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "ompcapi.h"
#include "all.h"

#endif
*/
import "C"
import (
	"unsafe"
)

func SendClientMessage(color uint32, text string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return bool(C.All_SendClientMessage(C.uint32_t(color), cText))
}

func CreateExplosion(x, y, z float32, explosionType int, radius float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.All_CreateExplosion(C.float(x), C.float(y), C.float(z), C.int(explosionType), C.float(radius)))
}

func SendDeathMessage(killer, killee Player, weapon int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !killer.isValid() || !killee.isValid() {
		return false
	}

	return bool(C.All_SendDeathMessage(killer.ptr, killee.ptr, C.int(weapon)))
}

func EnableStuntBonus(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.All_EnableStuntBonus(C.bool(enable)))
}
