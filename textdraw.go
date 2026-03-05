package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "ompcapi.h"
#include "textdraw.h"

#endif
*/
import "C"
import (
	"unsafe"
)

type TextDraw struct {
	ptr unsafe.Pointer
}

func textDrawFromPointer(ptr unsafe.Pointer) *TextDraw {
	if ptr == nil {
		return nil
	}

	return &TextDraw{ptr}
}

func TextDrawCreate(x, y float32, text string) *TextDraw {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	var id C.int

	ptr := C.TextDraw_Create(C.float(x), C.float(y), cText, &id)

	if ptr == nil {
		return nil
	}

	return &TextDraw{ptr}
}

func TextDrawFromID(textdrawID int) *TextDraw {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	ptr := C.TextDraw_FromID(C.int(textdrawID))

	if ptr == nil {
		return nil
	}

	return &TextDraw{ptr}
}

func (td *TextDraw) Destroy() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_Destroy(td.ptr))
}

func (td *TextDraw) GetID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0
	}

	return int(C.TextDraw_GetID(td.ptr))
}

func (td *TextDraw) IsValid() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_IsValid(td.ptr))
}

func (td *TextDraw) IsVisibleForPlayer(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.TextDraw_IsVisibleForPlayer(player.ptr, td.ptr))
}

func (td *TextDraw) SetLetterSize(sizeX, sizeY float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetLetterSize(td.ptr, C.float(sizeX), C.float(sizeY)))
}

func (td *TextDraw) SetTextSize(sizeX, sizeY float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetTextSize(td.ptr, C.float(sizeX), C.float(sizeY)))
}

func (td *TextDraw) SetAlignment(alignment int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetAlignment(td.ptr, C.int(alignment)))
}

func (td *TextDraw) SetColor(color uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetColor(td.ptr, C.uint32_t(color)))
}

func (td *TextDraw) SetUseBox(use bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetUseBox(td.ptr, C.bool(use)))
}

func (td *TextDraw) SetBoxColor(color uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetBoxColor(td.ptr, C.uint32_t(color)))
}

func (td *TextDraw) SetShadow(size int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetShadow(td.ptr, C.int(size)))
}

func (td *TextDraw) SetOutline(size int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetOutline(td.ptr, C.int(size)))
}

func (td *TextDraw) SetBackgroundColor(color uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetBackgroundColor(td.ptr, C.uint32_t(color)))
}

func (td *TextDraw) SetFont(font int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetFont(td.ptr, C.int(font)))
}

func (td *TextDraw) SetProportional(set bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetProportional(td.ptr, C.bool(set)))
}

func (td *TextDraw) SetSelectable(set bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetSelectable(td.ptr, C.bool(set)))
}

func (td *TextDraw) ShowForPlayer(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.TextDraw_ShowForPlayer(player.ptr, td.ptr))
}

func (td *TextDraw) HideForPlayer(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.TextDraw_HideForPlayer(player.ptr, td.ptr))
}

func (td *TextDraw) ShowForAll() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_ShowForAll(td.ptr))
}

func (td *TextDraw) HideForAll() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_HideForAll(td.ptr))
}

func (td *TextDraw) SetString(text string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return bool(C.TextDraw_SetString(td.ptr, cText))
}

func (td *TextDraw) SetPreviewModel(model int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetPreviewModel(td.ptr, C.int(model)))
}

func (td *TextDraw) SetPreviewRot(rotationX, rotationY, rotationZ, zoom float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetPreviewRot(td.ptr, C.float(rotationX), C.float(rotationY), C.float(rotationZ), C.float(zoom)))
}

func (td *TextDraw) SetPreviewVehCol(color1, color2 int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetPreviewVehCol(td.ptr, C.int(color1), C.int(color2)))
}

func (td *TextDraw) SetPos(x, y float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_SetPos(td.ptr, C.float(x), C.float(y)))
}

func (td *TextDraw) GetString() (string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return "", false
	}

	var text C.struct_CAPIStringView
	res := C.TextDraw_GetString(td.ptr, &text)

	if !res || text.data == nil {
		return "", false
	}

	return C.GoStringN(text.data, C.int(text.len)), true
}

func (td *TextDraw) GetLetterSize() (float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0, 0, false
	}

	var sizeX, sizeY C.float
	res := C.TextDraw_GetLetterSize(td.ptr, &sizeX, &sizeY)

	return float32(sizeX), float32(sizeY), bool(res)
}

func (td *TextDraw) GetTextSize() (float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0, 0, false
	}

	var sizeX, sizeY C.float
	res := C.TextDraw_GetTextSize(td.ptr, &sizeX, &sizeY)

	return float32(sizeX), float32(sizeY), bool(res)
}

func (td *TextDraw) GetPos() (float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0, 0, false
	}

	var x, y C.float
	res := C.TextDraw_GetPos(td.ptr, &x, &y)

	return float32(x), float32(y), bool(res)
}

func (td *TextDraw) GetColor() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0
	}

	return int(C.TextDraw_GetColor(td.ptr))
}

func (td *TextDraw) GetBoxColor() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0
	}

	return int(C.TextDraw_GetBoxColor(td.ptr))
}

func (td *TextDraw) GetBackgroundColor() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0
	}

	return int(C.TextDraw_GetBackgroundColor(td.ptr))
}

func (td *TextDraw) GetShadow() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0
	}

	return int(C.TextDraw_GetShadow(td.ptr))
}

func (td *TextDraw) GetOutline() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0
	}

	return int(C.TextDraw_GetOutline(td.ptr))
}

func (td *TextDraw) GetFont() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0
	}

	return int(C.TextDraw_GetFont(td.ptr))
}

func (td *TextDraw) IsBox() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_IsBox(td.ptr))
}

func (td *TextDraw) IsProportional() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_IsProportional(td.ptr))
}

func (td *TextDraw) IsSelectable() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return false
	}

	return bool(C.TextDraw_IsSelectable(td.ptr))
}

func (td *TextDraw) GetAlignment() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0
	}

	return int(C.TextDraw_GetAlignment(td.ptr))
}

func (td *TextDraw) GetPreviewModel() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0
	}

	return int(C.TextDraw_GetPreviewModel(td.ptr))
}

func (td *TextDraw) GetPreviewRot() (float32, float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0, 0, 0, 0, false
	}

	var x, y, z, zoom C.float
	res := C.TextDraw_GetPreviewRot(td.ptr, &x, &y, &z, &zoom)

	return float32(x), float32(y), float32(z), float32(zoom), bool(res)
}

func (td *TextDraw) GetPreviewVehColor() (int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil {
		return 0, 0, false
	}

	var color1, color2 C.int
	res := C.TextDraw_GetPreviewVehColor(td.ptr, &color1, &color2)

	return int(color1), int(color2), bool(res)
}

func (td *TextDraw) SetStringForPlayer(player *Player, text string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if td == nil || td.ptr == nil || player == nil || player.ptr == nil {
		return false
	}

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return bool(C.TextDraw_SetStringForPlayer(td.ptr, player.ptr, cText))
}
