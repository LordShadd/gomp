package colmap

/*
#include <stdlib.h>
#include "lib/bullet_capi.h"
*/
import "C"
import "unsafe"

func Init(cadbPath string) bool {
	worldMu.Lock()
	defer worldMu.Unlock()
	if world != nil {
		return true
	}
	world = C.cm_world_create()
	if world == nil {
		return false
	}
	if !loadCADB(cadbPath) {
		C.cm_world_destroy(world)
		world = nil
		return false
	}
	initMap()
	return true
}

func Destroy() {
	worldMu.Lock()
	defer worldMu.Unlock()
	if world == nil {
		return
	}
	shapeMu.Lock()
	for _, s := range modelShapes {
		C.cm_shape_destroy(s)
	}
	modelShapes = make(map[int32]C.CMShape)
	convexShapes = make(map[int32]C.CMShape)
	shapeMu.Unlock()
	C.cm_world_destroy(world)
	world = nil
}

func withWorld(fn func(C.CMWorld)) bool {
	worldMu.Lock()
	defer worldMu.Unlock()
	if world == nil {
		return false
	}
	fn(world)
	return true
}

func withWorldRet(fn func(C.CMWorld) int) (int, bool) {
	worldMu.Lock()
	defer worldMu.Unlock()
	if world == nil {
		return 0, false
	}
	return fn(world), true
}

func cStr(s string) *C.char {
	return C.CString(s)
}

func freeStr(s *C.char) {
	C.free(unsafe.Pointer(s))
}
