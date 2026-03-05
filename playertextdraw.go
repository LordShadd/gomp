package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "ompcapi.h"
#include "playertextdraw.h"

#endif
*/
import "C"
import (
	"unsafe"
)

type PlayerTextDraw struct {
	ptr    unsafe.Pointer
	player *Player
}

func playerTextDrawFromPointer(player *Player, ptr unsafe.Pointer) *PlayerTextDraw {
	if ptr == nil || player == nil {
		return nil
	}

	return &PlayerTextDraw{ptr, player}
}

func PlayerTextDrawCreate(player *Player, x, y float32, text string) *PlayerTextDraw {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if player == nil || player.ptr == nil {
		return nil
	}

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	var id C.int

	ptr := C.PlayerTextDraw_Create(player.ptr, C.float(x), C.float(y), cText, &id)

	if ptr == nil {
		return nil
	}

	return &PlayerTextDraw{ptr, player}
}

func PlayerTextDrawFromID(player *Player, textdrawID int) *PlayerTextDraw {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if player == nil || player.ptr == nil {
		return nil
	}

	ptr := C.PlayerTextDraw_FromID(player.ptr, C.int(textdrawID))

	if ptr == nil {
		return nil
	}

	return &PlayerTextDraw{ptr, player}
}

func (ptd *PlayerTextDraw) Destroy() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_Destroy(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) GetID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0
	}

	return int(C.PlayerTextDraw_GetID(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) IsValid() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_IsValid(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) IsVisible() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_IsVisible(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) SetLetterSize(x, y float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetLetterSize(ptd.player.ptr, ptd.ptr, C.float(x), C.float(y)))
}

func (ptd *PlayerTextDraw) SetTextSize(x, y float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetTextSize(ptd.player.ptr, ptd.ptr, C.float(x), C.float(y)))
}

func (ptd *PlayerTextDraw) SetAlignment(alignment int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetAlignment(ptd.player.ptr, ptd.ptr, C.int(alignment)))
}

func (ptd *PlayerTextDraw) SetColor(color uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetColor(ptd.player.ptr, ptd.ptr, C.uint32_t(color)))
}

func (ptd *PlayerTextDraw) UseBox(use bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_UseBox(ptd.player.ptr, ptd.ptr, C.bool(use)))
}

func (ptd *PlayerTextDraw) SetBoxColor(color uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetBoxColor(ptd.player.ptr, ptd.ptr, C.uint32_t(color)))
}

func (ptd *PlayerTextDraw) SetShadow(size int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetShadow(ptd.player.ptr, ptd.ptr, C.int(size)))
}

func (ptd *PlayerTextDraw) SetOutline(size int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetOutline(ptd.player.ptr, ptd.ptr, C.int(size)))
}

func (ptd *PlayerTextDraw) SetBackgroundColor(color uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetBackgroundColor(ptd.player.ptr, ptd.ptr, C.uint32_t(color)))
}

func (ptd *PlayerTextDraw) SetFont(font int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetFont(ptd.player.ptr, ptd.ptr, C.int(font)))
}

func (ptd *PlayerTextDraw) SetProportional(set bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetProportional(ptd.player.ptr, ptd.ptr, C.bool(set)))
}

func (ptd *PlayerTextDraw) SetSelectable(set bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetSelectable(ptd.player.ptr, ptd.ptr, C.bool(set)))
}

func (ptd *PlayerTextDraw) Show() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_Show(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) Hide() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_Hide(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) SetString(text string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return bool(C.PlayerTextDraw_SetString(ptd.player.ptr, ptd.ptr, cText))
}

func (ptd *PlayerTextDraw) SetPreviewModel(model int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetPreviewModel(ptd.player.ptr, ptd.ptr, C.int(model)))
}

func (ptd *PlayerTextDraw) SetPreviewRot(rx, ry, rz, zoom float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetPreviewRot(ptd.player.ptr, ptd.ptr, C.float(rx), C.float(ry), C.float(rz), C.float(zoom)))
}

func (ptd *PlayerTextDraw) SetPreviewVehCol(color1, color2 int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetPreviewVehCol(ptd.player.ptr, ptd.ptr, C.int(color1), C.int(color2)))
}

func (ptd *PlayerTextDraw) SetPos(x, y float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_SetPos(ptd.player.ptr, ptd.ptr, C.float(x), C.float(y)))
}

func (ptd *PlayerTextDraw) GetString() (string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return "", false
	}

	var text C.struct_CAPIStringView
	res := C.PlayerTextDraw_GetString(ptd.player.ptr, ptd.ptr, &text)

	if !res || text.data == nil {
		return "", false
	}

	return C.GoStringN(text.data, C.int(text.len)), true
}

func (ptd *PlayerTextDraw) GetLetterSize() (float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0, 0, false
	}

	var x, y C.float
	res := C.PlayerTextDraw_GetLetterSize(ptd.player.ptr, ptd.ptr, &x, &y)

	return float32(x), float32(y), bool(res)
}

func (ptd *PlayerTextDraw) GetTextSize() (float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0, 0, false
	}

	var x, y C.float
	res := C.PlayerTextDraw_GetTextSize(ptd.player.ptr, ptd.ptr, &x, &y)

	return float32(x), float32(y), bool(res)
}

func (ptd *PlayerTextDraw) GetPos() (float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0, 0, false
	}

	var x, y C.float
	res := C.PlayerTextDraw_GetPos(ptd.player.ptr, ptd.ptr, &x, &y)

	return float32(x), float32(y), bool(res)
}

func (ptd *PlayerTextDraw) GetColor() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0
	}

	return int(C.PlayerTextDraw_GetColor(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) GetBoxColor() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0
	}

	return int(C.PlayerTextDraw_GetBoxColor(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) GetBackgroundColor() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0
	}

	return int(C.PlayerTextDraw_GetBackgroundColor(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) GetShadow() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0
	}

	return int(C.PlayerTextDraw_GetShadow(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) GetOutline() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0
	}

	return int(C.PlayerTextDraw_GetOutline(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) GetFont() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0
	}

	return int(C.PlayerTextDraw_GetFont(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) IsBox() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_IsBox(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) IsProportional() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_IsProportional(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) IsSelectable() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return false
	}

	return bool(C.PlayerTextDraw_IsSelectable(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) GetAlignment() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0
	}

	return int(C.PlayerTextDraw_GetAlignment(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) GetPreviewModel() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0
	}

	return int(C.PlayerTextDraw_GetPreviewModel(ptd.player.ptr, ptd.ptr))
}

func (ptd *PlayerTextDraw) GetPreviewRot() (float32, float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0, 0, 0, 0, false
	}

	var rx, ry, rz, zoom C.float
	res := C.PlayerTextDraw_GetPreviewRot(ptd.player.ptr, ptd.ptr, &rx, &ry, &rz, &zoom)

	return float32(rx), float32(ry), float32(rz), float32(zoom), bool(res)
}

func (ptd *PlayerTextDraw) GetPreviewVehColor() (int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if ptd == nil || ptd.ptr == nil || ptd.player == nil || ptd.player.ptr == nil {
		return 0, 0, false
	}

	var color1, color2 C.int
	res := C.PlayerTextDraw_GetPreviewVehColor(ptd.player.ptr, ptd.ptr, &color1, &color2)

	return int(color1), int(color2), bool(res)
}
