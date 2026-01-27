package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "player.h"

#endif
*/
import "C"
import "unsafe"

const InvalidPlayerID = 0xFFFF

type Player struct {
	ptr unsafe.Pointer
}

func PlayerFromID(playerID int) *Player {
	ptr := C.Player_FromID(C.int(playerID))

	if ptr == nil {
		return nil
	}

	return &Player{ptr}
}

func PlayerFromPointer(ptr unsafe.Pointer) *Player {
	if ptr == nil {
		return nil
	}

	return &Player{ptr}
}

func (p *Player) GetID() int {
	if p == nil || p.ptr == nil {
		return InvalidPlayerID
	}

	return int(C.Player_GetID(p.ptr))
}

func (p *Player) GetName() (string, bool) {
	if p == nil || p.ptr == nil {
		return "", false
	}

	var nameView C.struct_CAPIStringView
	C.Player_GetName(p.ptr, &nameView)

	if nameView.data == nil {
		return "", false
	}

	return C.GoStringN(nameView.data, C.int(nameView.len)), true
}
