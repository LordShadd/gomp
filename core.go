package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "ompcapi.h"
#include "core.h"

#endif
*/
import "C"
import (
	"unsafe"
)

func TickCount() uint32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return uint32(C.Core_TickCount())
}

func MaxPlayers() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Core_MaxPlayers())
}

func Log(text string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return bool(C.Core_Log(cText))
}

func IsAdminTeleportAllowed() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_IsAdminTeleportAllowed())
}

func AllowAdminTeleport(allow bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_AllowAdminTeleport(C.bool(allow)))
}

func AreAllAnimationsEnabled() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_AreAllAnimationsEnabled())
}

func EnableAllAnimations(allow bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_EnableAllAnimations(C.bool(allow)))
}

func IsAnimationLibraryValid(name string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return bool(C.Core_IsAnimationLibraryValid(cName))
}

func AreInteriorWeaponsAllowed() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_AreInteriorWeaponsAllowed())
}

func AllowInteriorWeapons(allow bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_AllowInteriorWeapons(C.bool(allow)))
}

func BlockIpAddress(ip string, timeMS int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cIP := C.CString(ip)
	defer C.free(unsafe.Pointer(cIP))

	return bool(C.Core_BlockIpAddress(cIP, C.int(timeMS)))
}

func UnBlockIpAddress(ip string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cIP := C.CString(ip)
	defer C.free(unsafe.Pointer(cIP))

	return bool(C.Core_UnBlockIpAddress(cIP))
}

func DisableEntryExitMarkers() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_DisableEntryExitMarkers())
}

func DisableNameTagsLOS() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_DisableNameTagsLOS())
}

func EnableZoneNames(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_EnableZoneNames(C.bool(enable)))
}

func ShowGameTextForAll(msg string, time, style int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))

	return bool(C.Core_ShowGameTextForAll(cMsg, C.int(time), C.int(style)))
}

func HideGameTextForAll(style int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_HideGameTextForAll(C.int(style)))
}

func NetworkStats() (string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	buf := make([]byte, 4096)
	var strBuf C.struct_CAPIStringBuffer
	strBuf.capacity = C.uint(len(buf))
	strBuf.data = (*C.char)(unsafe.Pointer(&buf[0]))

	ret := C.Core_NetworkStats(&strBuf)

	if ret <= 0 {
		return "", false
	}

	return string(buf[:strBuf.len]), true
}

func ServerTickRate() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Core_ServerTickRate())
}

func GetWeaponName(weaponID int) (string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var nameView C.struct_CAPIStringView
	ret := C.Core_GetWeaponName(C.int(weaponID), &nameView)

	if !ret || nameView.data == nil {
		return "", false
	}

	return C.GoStringN(nameView.data, C.int(nameView.len)), true
}

func SetChatRadius(radius float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_SetChatRadius(C.float(radius)))
}

func SetMarkerRadius(radius float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_SetMarkerRadius(C.float(radius)))
}

func SendRconCommand(cmd string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cCmd := C.CString(cmd)
	defer C.free(unsafe.Pointer(cCmd))

	return bool(C.Core_SendRconCommand(cCmd))
}

func SetDeathDropAmount(amount int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_SetDeathDropAmount(C.int(amount)))
}

func GameModeSetText(text string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return bool(C.Core_GameMode_SetText(cText))
}

func SetGravity(gravity float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_SetGravity(C.float(gravity)))
}

func GetGravity() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return float32(C.Core_GetGravity())
}

func SetNameTagsDrawDistance(dist float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_SetNameTagsDrawDistance(C.float(dist)))
}

func SetWeather(weatherID int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_SetWeather(C.int(weatherID)))
}

func SetWorldTime(hour int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_SetWorldTime(C.int(hour)))
}

func ShowNameTags(show bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_ShowNameTags(C.bool(show)))
}

func ShowPlayerMarkers(mode int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_ShowPlayerMarkers(C.int(mode)))
}

func UsePedAnims() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_UsePedAnims())
}

func GetWeather() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Core_GetWeather())
}

func GetWorldTime() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Core_GetWorldTime())
}

func ToggleChatTextReplacement(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_ToggleChatTextReplacement(C.bool(enable)))
}

func IsChatTextReplacementToggled() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_IsChatTextReplacementToggled())
}

func IsNickNameValid(name string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return bool(C.Core_IsNickNameValid(cName))
}

func AllowNickNameCharacter(char int, allow bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_AllowNickNameCharacter(C.int(char), C.bool(allow)))
}

func IsNickNameCharacterAllowed(char int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_IsNickNameCharacterAllowed(C.int(char)))
}

func ClearBanList() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Core_ClearBanList())
}

func IsIpAddressBanned(ip string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cIP := C.CString(ip)
	defer C.free(unsafe.Pointer(cIP))

	return bool(C.Core_IsIpAddressBanned(cIP))
}

func GetWeaponSlot(weapon uint8) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Core_GetWeaponSlot(C.uint8_t(weapon)))
}

func AddRule(name, value string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	return bool(C.Core_AddRule(cName, cValue))
}

func IsValidRule(name string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return bool(C.Core_IsValidRule(cName))
}

func RemoveRule(name string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return bool(C.Core_RemoveRule(cName))
}
