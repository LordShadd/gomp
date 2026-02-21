package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "ompcapi.h"
#include "object.h"

#endif
*/
import "C"
import (
	"unsafe"
)

type Object struct {
	ptr unsafe.Pointer
}

func objectFromPointer(ptr unsafe.Pointer) *Object {
	if ptr == nil {
		return nil
	}

	return &Object{ptr}
}

func ObjectCreate(modelid int, x, y, z, rotationX,
	rotationY, rotationZ, drawDistance float32,
) *Object {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var id C.int

	ptr := C.Object_Create(C.int(modelid), C.float(x), C.float(y), C.float(z), C.float(rotationX), C.float(rotationY), C.float(rotationZ),
		C.float(drawDistance), &id)

	if ptr == nil {
		return nil
	}

	return &Object{ptr}
}

func ObjectFromID(objectID int) *Object {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	ptr := C.Object_FromID(C.int(objectID))

	if ptr == nil {
		return nil
	}

	return &Object{ptr}
}

func (o *Object) Destroy() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	return bool(C.Object_Destroy(o.ptr))
}

func (o *Object) GetID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0
	}

	return int(C.Object_GetID(o.ptr))
}

func (o *Object) AttachToVehicle(vehicle *Vehicle, offsetX, offsetY, offsetZ, rotationX, rotationY, rotationZ float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil || vehicle == nil || vehicle.ptr == nil {
		return false
	}

	return bool(C.Object_AttachToVehicle(o.ptr, vehicle.ptr, C.float(offsetX), C.float(offsetY), C.float(offsetZ), C.float(rotationX), C.float(rotationY), C.float(rotationZ)))
}

func (o *Object) AttachToObject(objAttachedTo *Object, offsetX, offsetY, offsetZ, rotationX, rotationY, rotationZ float32, syncRotation bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil || objAttachedTo == nil || objAttachedTo.ptr == nil {
		return false
	}

	return bool(C.Object_AttachToObject(o.ptr, objAttachedTo.ptr, C.float(offsetX), C.float(offsetY), C.float(offsetZ), C.float(rotationX), C.float(rotationY), C.float(rotationZ), C.bool(syncRotation)))
}

func (o *Object) AttachToPlayer(player *Player, offsetX, offsetY, offsetZ, rotationX, rotationY, rotationZ float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.Object_AttachToPlayer(o.ptr, player.ptr, C.float(offsetX), C.float(offsetY), C.float(offsetZ), C.float(rotationX), C.float(rotationY), C.float(rotationZ)))
}

func (o *Object) SetPos(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	return bool(C.Object_SetPos(o.ptr, C.float(x), C.float(y), C.float(z)))
}

func (o *Object) GetPos() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0, 0, 0, false
	}

	var x, y, z C.float
	res := C.Object_GetPos(o.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(res)
}

func (o *Object) SetRot(rotationX, rotationY, rotationZ float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	return bool(C.Object_SetRot(o.ptr, C.float(rotationX), C.float(rotationY), C.float(rotationZ)))
}

func (o *Object) GetRot() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0, 0, 0, false
	}

	var rotationX, rotationY, rotationZ C.float
	res := C.Object_GetRot(o.ptr, &rotationX, &rotationY, &rotationZ)

	return float32(rotationX), float32(rotationY), float32(rotationZ), bool(res)
}

func (o *Object) GetModel() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0
	}

	return int(C.Object_GetModel(o.ptr))
}

func (o *Object) SetNoCameraCollision() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	return bool(C.Object_SetNoCameraCollision(o.ptr))
}

func (o *Object) IsValid() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	return bool(C.Object_IsValid(o.ptr))
}

func (o *Object) Move(x, y, z, speed, rotationX, rotationY, rotationZ float32) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0
	}

	return int(C.Object_Move(o.ptr, C.float(x), C.float(y), C.float(z), C.float(speed), C.float(rotationX), C.float(rotationY), C.float(rotationZ)))
}

func (o *Object) Stop() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	return bool(C.Object_Stop(o.ptr))
}

func (o *Object) IsMoving() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	return bool(C.Object_IsMoving(o.ptr))
}

func (o *Object) BeginEditing(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.Object_BeginEditing(player.ptr, o.ptr))
}

func ObjectBeginSelecting(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if player == nil || player.ptr == nil {
		return false
	}

	return bool(C.Object_BeginSelecting(player.ptr))
}

func ObjectEndEditing(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if player == nil || player.ptr == nil {
		return false
	}

	return bool(C.Object_EndEditing(player.ptr))
}

func (o *Object) SetMaterial(materialIndex, modelId int, textureLibrary, textureName string, materialColor uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	cTextureLibrary := C.CString(textureLibrary)
	defer C.free(unsafe.Pointer(cTextureLibrary))

	cTextureName := C.CString(textureName)
	defer C.free(unsafe.Pointer(cTextureName))

	return bool(C.Object_SetMaterial(o.ptr, C.int(materialIndex), C.int(modelId), cTextureLibrary, cTextureName, C.uint32_t(materialColor)))
}

func (o *Object) SetMaterialText(text string, materialIndex, materialSize int, fontface string, fontsize int, bold bool, fontColor, backgroundColor uint32, textalignment int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	cFontface := C.CString(fontface)
	defer C.free(unsafe.Pointer(cFontface))

	return bool(C.Object_SetMaterialText(o.ptr, cText, C.int(materialIndex), C.int(materialSize), cFontface, C.int(fontsize), C.bool(bold), C.uint32_t(fontColor), C.uint32_t(backgroundColor), C.int(textalignment)))
}

func ObjectSetDefaultCameraCollision(disable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Object_SetDefaultCameraCollision(C.bool(disable)))
}

func (o *Object) GetDrawDistance() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0
	}

	return float32(C.Object_GetDrawDistance(o.ptr))
}

func (o *Object) GetMoveSpeed() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0
	}

	return float32(C.Object_GetMoveSpeed(o.ptr))
}

func (o *Object) GetMovingTargetPos() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0, 0, 0, false
	}

	var targetX, targetY, targetZ C.float
	res := C.Object_GetMovingTargetPos(o.ptr, &targetX, &targetY, &targetZ)

	return float32(targetX), float32(targetY), float32(targetZ), bool(res)
}

func (o *Object) GetMovingTargetRot() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0, 0, 0, false
	}

	var rotationX, rotationY, rotationZ C.float
	res := C.Object_GetMovingTargetRot(o.ptr, &rotationX, &rotationY, &rotationZ)

	return float32(rotationX), float32(rotationY), float32(rotationZ), bool(res)
}

func (o *Object) GetAttachedData() (int, int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0, 0, 0, false
	}

	var parentVehicle, parentObject, parentPlayer C.int
	res := C.Object_GetAttachedData(o.ptr, &parentVehicle, &parentObject, &parentPlayer)

	return int(parentVehicle), int(parentObject), int(parentPlayer), bool(res)
}

func (o *Object) GetAttachedOffset() (float32, float32, float32, float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0, 0, 0, 0, 0, 0, false
	}

	var offsetX, offsetY, offsetZ, rotationX, rotationY, rotationZ C.float
	res := C.Object_GetAttachedOffset(o.ptr, &offsetX, &offsetY, &offsetZ, &rotationX, &rotationY, &rotationZ)

	return float32(offsetX), float32(offsetY), float32(offsetZ), float32(rotationX), float32(rotationY), float32(rotationZ), bool(res)
}

func (o *Object) GetSyncRotation() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	return bool(C.Object_GetSyncRotation(o.ptr))
}

func (o *Object) IsMaterialSlotUsed(materialIndex int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	return bool(C.Object_IsMaterialSlotUsed(o.ptr, C.int(materialIndex)))
}

func (o *Object) GetMaterial(materialIndex int) (int, string, string, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return 0, "", "", 0, false
	}

	var modelid C.int
	var textureLibrary C.struct_CAPIStringView
	var textureName C.struct_CAPIStringView
	var materialColor C.int

	C.Object_GetMaterial(o.ptr, C.int(materialIndex), &modelid, &textureLibrary, &textureName, &materialColor)

	if textureLibrary.data == nil || textureName.data == nil {
		return 0, "", "", 0, false
	}

	return int(modelid), C.GoStringN(textureLibrary.data, C.int(textureLibrary.len)), C.GoStringN(textureName.data, C.int(textureName.len)), int(materialColor), true
}

func (o *Object) GetMaterialText(materialIndex int) (string, int, string, int, bool, uint32, uint32, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return "", 0, "", 0, false, 0, 0, 0, false
	}

	var text C.struct_CAPIStringView
	var materialSize C.int
	var fontFace C.struct_CAPIStringView
	var fontSize C.int
	var bold C.bool
	var fontColor C.int
	var backgroundColor C.int
	var textAlignment C.int

	C.Object_GetMaterialText(o.ptr, C.int(materialIndex), &text, &materialSize, &fontFace, &fontSize, &bold, &fontColor, &backgroundColor, &textAlignment)

	if text.data == nil || fontFace.data == nil {
		return "", 0, "", 0, false, 0, 0, 0, false
	}

	return C.GoStringN(text.data, C.int(text.len)), int(materialSize), C.GoStringN(fontFace.data, C.int(fontFace.len)), int(fontSize), bool(bold), uint32(fontColor), uint32(backgroundColor), int(textAlignment), true
}

func (o *Object) IsObjectNoCameraCollision() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if o == nil || o.ptr == nil {
		return false
	}

	return bool(C.Object_IsObjectNoCameraCollision(o.ptr))
}

func ObjectGetType(player *Player, objectID int) uint8 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if player == nil || player.ptr == nil {
		return 0
	}

	return uint8(C.Object_GetType(player.ptr, C.int(objectID)))
}
