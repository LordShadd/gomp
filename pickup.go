package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "ompcapi.h"
#include "pickup.h"

#endif
*/
import "C"
import (
	"unsafe"
)

type Pickup struct {
	ptr unsafe.Pointer
}

func pickupFromPointer(ptr unsafe.Pointer) *Pickup {
	if ptr == nil {
		return nil
	}

	return &Pickup{ptr}
}

func PickupFromID(pickupID int) *Pickup {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	ptr := C.Pickup_FromID(C.int(pickupID))

	if ptr == nil {
		return nil
	}

	return &Pickup{ptr}
}

func PickupCreate(model, pickupType int, x, y, z float32,
	virtualWorld int) *Pickup {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var pickupID C.int

	ptr := C.Pickup_Create(C.int(model), C.int(pickupType), C.float(x), C.float(y), C.float(z), C.int(virtualWorld), &pickupID)

	if ptr == nil {
		return nil
	}

	return &Pickup{
		ptr,
	}
}

func PickupAddStatic(model int, pickupType int, x, y, z float32,
	virtualWorld int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Pickup_AddStatic(C.int(model), C.int(pickupType), C.float(x), C.float(y), C.float(z), C.int(virtualWorld)))
}

func (p *Pickup) Destroy() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Pickup_Destroy(p.ptr))
}

func (p *Pickup) GetID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Pickup_GetID(p.ptr))
}

func (p *Pickup) IsValid() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Pickup_IsValid(p.ptr))
}

func (p *Pickup) IsStreamedIn(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if p == nil || p.ptr == nil {
		return false
	}

	return bool(C.Pickup_IsStreamedIn(player.ptr, p.ptr))
}

func (p *Pickup) GetPos() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var x, y, z C.float

	ret := C.Pickup_GetPos(p.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (p *Pickup) GetModel() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Pickup_GetModel(p.ptr))
}

func (p *Pickup) GetType() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Pickup_GetType(p.ptr))
}

func (p *Pickup) GetVirtualWorld() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Pickup_GetVirtualWorld(p.ptr))
}

func (p *Pickup) SetPos(x, y, z float32, update bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Pickup_SetPos(p.ptr, C.float(x), C.float(y), C.float(z), C.bool(update)))
}

func (p *Pickup) SetModel(model int, update bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Pickup_SetModel(p.ptr, C.int(model), C.bool(update)))
}

func (p *Pickup) SetType(pickupType int, update bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Pickup_SetType(p.ptr, C.int(pickupType), C.bool(update)))
}

func (p *Pickup) SetVirtualWorld(virtualworld int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Pickup_SetVirtualWorld(p.ptr, C.int(virtualworld)))
}

func (p *Pickup) ShowForPlayer(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Pickup_ShowForPlayer(player.ptr, p.ptr))
}

func (p *Pickup) HideForPlayer(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Pickup_HideForPlayer(player.ptr, p.ptr))
}

func (p *Pickup) IsHiddenForPlayer(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Pickup_IsHiddenForPlayer(player.ptr, p.ptr))
}
