package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "actor.h"

#endif
*/
import "C"
import "unsafe"

type Actor struct {
	ptr unsafe.Pointer
}

func actorFromPointer(ptr unsafe.Pointer) *Actor {
	if ptr == nil {
		return nil
	}

	return &Actor{ptr}
}

func ActorCreate(model int, x, y, z, rot float32) *Actor {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var id C.int

	ptr := C.Actor_Create(C.int(model), C.float(x), C.float(y), C.float(z), C.float(rot), &id)

	return actorFromPointer(ptr)
}

func ActorFromID(actorid int) *Actor {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	ptr := C.Actor_FromID(C.int(actorid))

	return actorFromPointer(ptr)
}

func (a *Actor) isValid() bool {
	return a != nil && a.ptr != nil
}

func (a *Actor) Destroy() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	return bool(C.Actor_Destroy(a.ptr))
}

func (a *Actor) GetID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return -1
	}

	return int(C.Actor_GetID(a.ptr))
}

func (a *Actor) IsStreamedInFor(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.Actor_IsStreamedInFor(a.ptr, player.ptr))
}

func (a *Actor) SetVirtualWorld(vw int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	return bool(C.Actor_SetVirtualWorld(a.ptr, C.int(vw)))
}

func (a *Actor) GetVirtualWorld() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return -1
	}

	return int(C.Actor_GetVirtualWorld(a.ptr))
}

func (a *Actor) ApplyAnimation(name, library string, delta float32, loop, lockX, lockY, freeze bool, time int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cLibrary := C.CString(library)
	defer C.free(unsafe.Pointer(cLibrary))

	return bool(C.Actor_ApplyAnimation(a.ptr, cName, cLibrary, C.float(delta), C.bool(loop), C.bool(lockX), C.bool(lockY), C.bool(freeze), C.int(time)))
}

func (a *Actor) ClearAnimations() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	return bool(C.Actor_ClearAnimations(a.ptr))
}

func (a *Actor) SetPos(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	return bool(C.Actor_SetPos(a.ptr, C.float(x), C.float(y), C.float(z)))
}

func (a *Actor) GetPos() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return 0, 0, 0, false
	}

	var x, y, z C.float
	ret := C.Actor_GetPos(a.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (a *Actor) SetFacingAngle(angle float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	return bool(C.Actor_SetFacingAngle(a.ptr, C.float(angle)))
}

func (a *Actor) GetFacingAngle() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return 0.0
	}

	return float32(C.Actor_GetFacingAngle(a.ptr))
}

func (a *Actor) SetHealth(hp float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	return bool(C.Actor_SetHealth(a.ptr, C.float(hp)))
}

func (a *Actor) GetHealth() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return 0.0
	}

	return float32(C.Actor_GetHealth(a.ptr))
}

func (a *Actor) SetInvulnerable(toggle bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	return bool(C.Actor_SetInvulnerable(a.ptr, C.bool(toggle)))
}

func (a *Actor) IsInvulnerable() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	return bool(C.Actor_IsInvulnerable(a.ptr))
}

func (a *Actor) IsValid() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	return bool(C.Actor_IsValid(a.ptr))
}

func (a *Actor) SetSkin(skin int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return false
	}

	return bool(C.Actor_SetSkin(a.ptr, C.int(skin)))
}

func (a *Actor) GetSkin() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return -1
	}

	return int(C.Actor_GetSkin(a.ptr))
}

func (a *Actor) GetAnimation() (string, string, float32, bool, bool, bool, bool, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return "", "", 0.0, false, false, false, false, 0, false
	}

	var libraryView C.struct_CAPIStringView
	var nameView C.struct_CAPIStringView
	var delta C.float
	var loop, lockX, lockY, freeze C.bool
	var time C.int

	ret := C.Actor_GetAnimation(a.ptr, &libraryView, &nameView, &delta, &loop, &lockX, &lockY, &freeze, &time)

	library := C.GoStringN(libraryView.data, C.int(libraryView.len))
	name := C.GoStringN(nameView.data, C.int(nameView.len))

	return library, name, float32(delta), bool(loop), bool(lockX), bool(lockY), bool(freeze), int(time), bool(ret)
}

func (a *Actor) GetSpawnInfo() (float32, float32, float32, float32, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !a.isValid() {
		return 0, 0, 0, 0, 0, false
	}

	var x, y, z, angle C.float
	var skin C.int

	ret := C.Actor_GetSpawnInfo(a.ptr, &x, &y, &z, &angle, &skin)

	return float32(x), float32(y), float32(z), float32(angle), int(skin), bool(ret)
}
