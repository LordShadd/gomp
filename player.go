package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "player.h"
#include "vehicle.h"

#endif
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const InvalidPlayerID = 0xFFFF

type Player struct {
	ptr unsafe.Pointer
}

func playerFromPointer(ptr unsafe.Pointer) *Player {
	if ptr == nil {
		return nil
	}

	return &Player{ptr}
}

func PlayerFromID(playerID int) *Player {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	ptr := C.Player_FromID(C.int(playerID))

	if ptr == nil {
		return nil
	}

	return &Player{ptr}
}

func PlayerGetAnimationName(index int) (string, string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var lib, name C.struct_CAPIStringView

	ret := C.Player_GetAnimationName(
		C.int(index),
		&lib,
		&name,
	)

	if !bool(ret) {
		return "", "", false
	}

	sLib := C.GoStringN(lib.data, C.int(lib.len))
	sName := C.GoStringN(name.data, C.int(name.len))

	return sLib, sName, true
}

func (p *Player) isValid() bool {
	if p == nil || p.ptr == nil {
		return false
	}

	playerID := int(C.Player_GetID(p.ptr))

	if playerID == InvalidPlayerID || playerID == -1 {
		return false
	}

	return true
}

func (p *Player) GetID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if p == nil || p.ptr == nil {
		return InvalidPlayerID
	}

	return int(C.Player_GetID(p.ptr))
}

func (p *Player) GetName() (string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return "", false
	}

	var nameView C.struct_CAPIStringView
	C.Player_GetName(p.ptr, &nameView)

	if nameView.data == nil {
		return "", false
	}

	return C.GoStringN(nameView.data, C.int(nameView.len)), true
}

func (p *Player) GetCustomSkin() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return -1
	}

	return int(C.Player_GetCustomSkin(p.ptr))
}

func (p *Player) GetDialog() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return -1
	}

	return int(C.Player_GetDialog(p.ptr))
}

func (p *Player) SendClientMessage(color uint32, text string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return bool(C.Player_SendClientMessage(p.ptr, C.uint32_t(color), cText))
}

func (p *Player) SetCameraPos(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetCameraPos(p.ptr, C.float(x), C.float(y), C.float(z)))
}

func (p *Player) SetDrunkLevel(level int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetDrunkLevel(p.ptr, C.int(level)))
}

func (p *Player) SetInterior(interior int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetInterior(p.ptr, C.int(interior)))
}

func (p *Player) SetWantedLevel(level int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetWantedLevel(p.ptr, C.int(level)))
}

func (p *Player) SetWeather(weather int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetWeather(p.ptr, C.int(weather)))
}

func (p *Player) GetWeather() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetWeather(p.ptr))
}

func (p *Player) SetSkin(skin int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetSkin(p.ptr, C.int(skin)))
}

func (p *Player) SetShopName(name string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return bool(C.Player_SetShopName(p.ptr, cName))
}

func (p *Player) GiveMoney(amount int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_GiveMoney(p.ptr, C.int(amount)))
}

func (p *Player) SetCameraLookAt(x, y, z float32, cutType int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetCameraLookAt(p.ptr, C.float(x), C.float(y), C.float(z), C.int(cutType)))
}

func (p *Player) SetCameraBehind() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetCameraBehind(p.ptr))
}

func (p *Player) CreateExplosion(x, y, z float32, typeID int, radius float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_CreateExplosion(p.ptr, C.float(x), C.float(y), C.float(z), C.int(typeID), C.float(radius)))
}

func (p *Player) PlayAudioStream(url string, x, y, z, distance float32, usePos bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))

	return bool(C.Player_PlayAudioStream(p.ptr, cUrl, C.float(x), C.float(y), C.float(z), C.float(distance), C.bool(usePos)))
}

func (p *Player) StopAudioStream() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_StopAudioStream(p.ptr))
}

func (p *Player) ToggleWidescreen(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_ToggleWidescreen(p.ptr, C.bool(enable)))
}

func (p *Player) IsWidescreenToggled() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsWidescreenToggled(p.ptr))
}

func (p *Player) SetHealth(health float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetHealth(p.ptr, C.float(health)))
}

func (p *Player) GetHealth() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0.0
	}

	return float32(C.Player_GetHealth(p.ptr))
}

func (p *Player) SetArmor(armor float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetArmor(p.ptr, C.float(armor)))
}

func (p *Player) GetArmor() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0.0
	}

	return float32(C.Player_GetArmor(p.ptr))
}

func (p *Player) SetTeam(team int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetTeam(p.ptr, C.int(team)))
}

func (p *Player) GetTeam() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetTeam(p.ptr))
}

func (p *Player) SetScore(score int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetScore(p.ptr, C.int(score)))
}

func (p *Player) GetScore() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetScore(p.ptr))
}

func (p *Player) GetSkin() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetSkin(p.ptr))
}

func (p *Player) SetColor(color uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetColor(p.ptr, C.uint32_t(color)))
}

func (p *Player) GetColor() uint32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return uint32(C.Player_GetColor(p.ptr))
}

func (p *Player) GetDefaultColor() uint32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return uint32(C.Player_GetDefaultColor(p.ptr))
}

func (p *Player) GetDrunkLevel() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetDrunkLevel(p.ptr))
}

func (p *Player) GiveWeapon(weapon, ammo int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_GiveWeapon(p.ptr, C.int(weapon), C.int(ammo)))
}

func (p *Player) RemoveWeapon(weapon int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_RemoveWeapon(p.ptr, C.int(weapon)))
}

func (p *Player) GetMoney() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetMoney(p.ptr))
}

func (p *Player) ResetMoney() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_ResetMoney(p.ptr))
}

func (p *Player) SetName(name string) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int(C.Player_SetName(p.ptr, cName))
}

func (p *Player) GetState() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetState(p.ptr))
}

func (p *Player) GetPing() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetPing(p.ptr))
}

func (p *Player) GetWeapon() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetWeapon(p.ptr))
}

func (p *Player) SetTime(hour, minute int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetTime(p.ptr, C.int(hour), C.int(minute)))
}

func (p *Player) GetTime() (int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, false
	}

	var hour, minute C.int
	ret := C.Player_GetTime(p.ptr, &hour, &minute)

	return int(hour), int(minute), bool(ret)
}

func (p *Player) ToggleClock(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_ToggleClock(p.ptr, C.bool(enable)))
}

func (p *Player) HasClock() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_HasClock(p.ptr))
}

func (p *Player) ForceClassSelection() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_ForceClassSelection(p.ptr))
}

func (p *Player) GetWantedLevel() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetWantedLevel(p.ptr))
}

func (p *Player) SetFightingStyle(style int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetFightingStyle(p.ptr, C.int(style)))
}

func (p *Player) GetFightingStyle() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetFightingStyle(p.ptr))
}

func (p *Player) SetVelocity(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetVelocity(p.ptr, C.float(x), C.float(y), C.float(z)))
}

func (p *Player) GetVelocity() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, 0, false
	}

	var x, y, z C.float
	ret := C.Player_GetVelocity(p.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (p *Player) GetCameraPos() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, 0, false
	}

	var x, y, z C.float
	ret := C.Player_GetCameraPos(p.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (p *Player) GetDistanceFromPoint(x, y, z float32) float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0.0
	}

	return float32(C.Player_GetDistanceFromPoint(p.ptr, C.float(x), C.float(y), C.float(z)))
}

func (p *Player) GetInterior() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetInterior(p.ptr))
}

func (p *Player) SetPos(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetPos(p.ptr, C.float(x), C.float(y), C.float(z)))
}

func (p *Player) GetPos() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, 0, false
	}

	var x, y, z C.float
	ret := C.Player_GetPos(p.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (p *Player) GetVirtualWorld() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetVirtualWorld(p.ptr))
}

func (p *Player) IsNPC() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsNPC(p.ptr))
}

func (p *Player) IsStreamedIn(other *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() || other == nil || other.ptr == nil || other.GetID() == -1 {
		return false
	}

	return bool(C.Player_IsStreamedIn(p.ptr, other.ptr))
}

func (p *Player) PlayGameSound(sound int, x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_PlayGameSound(p.ptr, C.int(sound), C.float(x), C.float(y), C.float(z)))
}

func (p *Player) SpectatePlayer(target *Player, mode int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() || target == nil || target.ptr == nil || target.GetID() == -1 {
		return false
	}

	return bool(C.Player_SpectatePlayer(p.ptr, target.ptr, C.int(mode)))
}

func (p *Player) SpectateVehicle(targetVehicle unsafe.Pointer, mode int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SpectateVehicle(p.ptr, targetVehicle, C.int(mode)))
}

func (p *Player) SetVirtualWorld(vw int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetVirtualWorld(p.ptr, C.int(vw)))
}

func (p *Player) SetWorldBounds(xMax, xMin, yMax, yMin float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetWorldBounds(p.ptr, C.float(xMax), C.float(xMin), C.float(yMax), C.float(yMin)))
}

func (p *Player) ClearWorldBounds() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_ClearWorldBounds(p.ptr))
}

func (p *Player) GetWorldBounds() (float32, float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, 0, 0, false
	}

	var xMax, xMin, yMax, yMin C.float
	ret := C.Player_GetWorldBounds(p.ptr, &xMax, &xMin, &yMax, &yMin)

	return float32(xMax), float32(xMin), float32(yMax), float32(yMin), bool(ret)
}

func (p *Player) ClearAnimations(syncType int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_ClearAnimations(p.ptr, C.int(syncType)))
}

func (p *Player) GetLastShotVectors() (float32, float32, float32, float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, 0, 0, 0, 0, false
	}

	var ox, oy, oz, hx, hy, hz C.float
	ret := C.Player_GetLastShotVectors(p.ptr, &ox, &oy, &oz, &hx, &hy, &hz)

	return float32(ox), float32(oy), float32(oz), float32(hx), float32(hy), float32(hz), bool(ret)
}

func (p *Player) GetCameraTargetPlayer() *Player {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return playerFromPointer(C.Player_GetCameraTargetPlayer(p.ptr))
}

func (p *Player) GetCameraTargetActor() unsafe.Pointer {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return C.Player_GetCameraTargetActor(p.ptr)
}

func (p *Player) GetCameraTargetObject() unsafe.Pointer {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return C.Player_GetCameraTargetObject(p.ptr)
}

func (p *Player) GetCameraTargetVehicle() unsafe.Pointer {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return C.Player_GetCameraTargetVehicle(p.ptr)
}

func (p *Player) PutInVehicle(vehicle *Vehicle, seat int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_PutInVehicle(p.ptr, vehicle.ptr, C.int(seat)))
}

func (p *Player) RemoveBuilding(model int, x, y, z, radius float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_RemoveBuilding(p.ptr, C.int(model), C.float(x), C.float(y), C.float(z), C.float(radius)))
}

func (p *Player) GetBuildingsRemoved() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetBuildingsRemoved(p.ptr))
}

func (p *Player) RemoveFromVehicle(force bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_RemoveFromVehicle(p.ptr, C.bool(force)))
}

func (p *Player) RemoveMapIcon(icon int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_RemoveMapIcon(p.ptr, C.int(icon)))
}

func (p *Player) SetMapIcon(iconID int, x, y, z float32, iconType int, color uint32, style int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetMapIcon(p.ptr, C.int(iconID), C.float(x), C.float(y), C.float(z), C.int(iconType), C.uint32_t(color), C.int(style)))
}

func (p *Player) ResetWeapons() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_ResetWeapons(p.ptr))
}

func (p *Player) SetAmmo(id uint8, ammo uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetAmmo(p.ptr, C.uint8_t(id), C.uint32_t(ammo)))
}

func (p *Player) SetArmedWeapon(weapon uint8) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetArmedWeapon(p.ptr, C.uint8_t(weapon)))
}

func (p *Player) SetChatBubble(text string, color uint32, drawDistance float32, expireTime int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return bool(C.Player_SetChatBubble(p.ptr, cText, C.uint32_t(color), C.float(drawDistance), C.int(expireTime)))
}

func (p *Player) SetPosFindZ(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetPosFindZ(p.ptr, C.float(x), C.float(y), C.float(z)))
}

func (p *Player) SetSkillLevel(weapon uint8, level int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetSkillLevel(p.ptr, C.uint8_t(weapon), C.int(level)))
}

func (p *Player) SetSpecialAction(action uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetSpecialAction(p.ptr, C.uint32_t(action)))
}

func (p *Player) ShowNameTagForPlayer(other *Player, enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() || other == nil || other.ptr == nil || other.GetID() == -1 {
		return false
	}

	return bool(C.Player_ShowNameTagForPlayer(p.ptr, other.ptr, C.bool(enable)))
}

func (p *Player) ToggleControllable(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_ToggleControllable(p.ptr, C.bool(enable)))
}

func (p *Player) ToggleSpectating(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_ToggleSpectating(p.ptr, C.bool(enable)))
}

func (p *Player) ApplyAnimation(animLib, animName string, delta float32, loop, lockX, lockY, freeze bool, time int, sync int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	cAnimLib := C.CString(animLib)
	defer C.free(unsafe.Pointer(cAnimLib))

	cAnimName := C.CString(animName)
	defer C.free(unsafe.Pointer(cAnimName))

	return bool(C.Player_ApplyAnimation(p.ptr, cAnimLib, cAnimName, C.float(delta), C.bool(loop), C.bool(lockX), C.bool(lockY), C.bool(freeze), C.uint32_t(time), C.int(sync)))
}

func (p *Player) EditAttachedObject(index int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_EditAttachedObject(p.ptr, C.int(index)))
}

func (p *Player) EnableCameraTarget(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_EnableCameraTarget(p.ptr, C.bool(enable)))
}

func (p *Player) EnableStuntBonus(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_EnableStuntBonus(p.ptr, C.bool(enable)))
}

func (p *Player) GetPlayerAmmo() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetPlayerAmmo(p.ptr))
}

func (p *Player) GetAnimationIndex() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetAnimationIndex(p.ptr))
}

func (p *Player) GetFacingAngle() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0.0
	}

	return float32(C.Player_GetFacingAngle(p.ptr))
}

func (p *Player) GetSpecialAction() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetSpecialAction(p.ptr))
}

func (p *Player) GetVehicleID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetVehicleID(p.ptr))
}

func (p *Player) GetVehicleSeat() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetVehicleSeat(p.ptr))
}

func (p *Player) GetWeaponData(slot int) (int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, false
	}

	var weapon, ammo C.int
	ret := C.Player_GetWeaponData(p.ptr, C.int(slot), &weapon, &ammo)

	return int(weapon), int(ammo), bool(ret)
}

func (p *Player) GetWeaponState() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetWeaponState(p.ptr))
}

func (p *Player) InterpolateCameraPos(fromX, fromY, fromZ, toX, toY, toZ float32, time, cut int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_InterpolateCameraPos(p.ptr, C.float(fromX), C.float(fromY), C.float(fromZ), C.float(toX), C.float(toY), C.float(toZ), C.int(time), C.int(cut)))
}

func (p *Player) InterpolateCameraLookAt(fromX, fromY, fromZ, toX, toY, toZ float32, time, cut int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_InterpolateCameraLookAt(p.ptr, C.float(fromX), C.float(fromY), C.float(fromZ), C.float(toX), C.float(toY), C.float(toZ), C.int(time), C.int(cut)))
}

func (p *Player) IsPlayerAttachedObjectSlotUsed(index int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsPlayerAttachedObjectSlotUsed(p.ptr, C.int(index)))
}

func (p *Player) AttachCameraToObject(object unsafe.Pointer) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_AttachCameraToObject(p.ptr, object))
}

func (p *Player) AttachCameraToPlayerObject(object unsafe.Pointer) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_AttachCameraToPlayerObject(p.ptr, object))
}

func (p *Player) GetCameraAspectRatio() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0.0
	}

	return float32(C.Player_GetCameraAspectRatio(p.ptr))
}

func (p *Player) GetCameraFrontVector() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, 0, false
	}

	var x, y, z C.float
	ret := C.Player_GetCameraFrontVector(p.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (p *Player) GetCameraMode() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetCameraMode(p.ptr))
}

func (p *Player) GetKeys() (int, int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, 0, false
	}

	var keys, updown, leftright C.int
	ret := C.Player_GetKeys(p.ptr, &keys, &updown, &leftright)

	return int(keys), int(updown), int(leftright), bool(ret)
}

func (p *Player) GetSurfingVehicle() unsafe.Pointer {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return C.Player_GetSurfingVehicle(p.ptr)
}

func (p *Player) GetSurfingObject() unsafe.Pointer {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return C.Player_GetSurfingObject(p.ptr)
}

func (p *Player) GetTargetPlayer() *Player {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return playerFromPointer(C.Player_GetTargetPlayer(p.ptr))
}

func (p *Player) GetTargetActor() unsafe.Pointer {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return C.Player_GetTargetActor(p.ptr)
}

func (p *Player) IsInVehicle(vehicle unsafe.Pointer) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsInVehicle(p.ptr, vehicle))
}

func (p *Player) IsInAnyVehicle() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsInAnyVehicle(p.ptr))
}

func (p *Player) IsInRangeOfPoint(rangeVal, x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsInRangeOfPoint(p.ptr, C.float(rangeVal), C.float(x), C.float(y), C.float(z)))
}

func (p *Player) PlayCrimeReport(suspect *Player, crime int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() || suspect == nil || suspect.ptr == nil || suspect.GetID() == -1 {
		return false
	}

	return bool(C.Player_PlayCrimeReport(p.ptr, suspect.ptr, C.int(crime)))
}

func (p *Player) RemoveAttachedObject(index int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_RemoveAttachedObject(p.ptr, C.int(index)))
}

func (p *Player) SetAttachedObject(index, modelid, bone int, offX, offY, offZ, rotX, rotY, rotZ, sclX, sclY, sclZ float32, color1, color2 uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetAttachedObject(p.ptr, C.int(index), C.int(modelid), C.int(bone), C.float(offX), C.float(offY), C.float(offZ), C.float(rotX), C.float(rotY), C.float(rotZ), C.float(sclX), C.float(sclY), C.float(sclZ), C.uint32_t(color1), C.uint32_t(color2)))
}

func (p *Player) SetMarkerForPlayer(other *Player, color uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() || other == nil || other.ptr == nil || other.GetID() == -1 {
		return false
	}

	return bool(C.Player_SetMarkerForPlayer(p.ptr, other.ptr, C.uint32_t(color)))
}

func (p *Player) GetMarkerForPlayer(other *Player) uint32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() || other == nil || other.ptr == nil || other.GetID() == -1 {
		return 0
	}

	return uint32(C.Player_GetMarkerForPlayer(p.ptr, other.ptr))
}

func (p *Player) AllowTeleport(allow bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_AllowTeleport(p.ptr, C.bool(allow)))
}

func (p *Player) IsTeleportAllowed() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsTeleportAllowed(p.ptr))
}

func (p *Player) DisableRemoteVehicleCollisions(disable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_DisableRemoteVehicleCollisions(p.ptr, C.bool(disable)))
}

func (p *Player) GetCameraZoom() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0.0
	}

	return float32(C.Player_GetCameraZoom(p.ptr))
}

func (p *Player) SelectTextDraw(hoverColor uint32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SelectTextDraw(p.ptr, C.uint32_t(hoverColor)))
}

func (p *Player) CancelSelectTextDraw() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_CancelSelectTextDraw(p.ptr))
}

func (p *Player) SendClientCheck(actionType, address, offset, count int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SendClientCheck(p.ptr, C.int(actionType), C.int(address), C.int(offset), C.int(count)))
}

func (p *Player) ShowGameText(text string, time, style int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return bool(C.Player_ShowGameText(p.ptr, cText, C.int(time), C.int(style)))
}

func (p *Player) HideGameText(style int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_HideGameText(p.ptr, C.int(style)))
}

func (p *Player) HasGameText(style int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_HasGameText(p.ptr, C.int(style)))
}

func (p *Player) GetGameText(style int) (string, int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return "", 0, 0, false
	}

	var messageView C.struct_CAPIStringView
	var time, remaining C.int

	ret := C.Player_GetGameText(
		p.ptr,
		C.int(style),
		&messageView,
		&time,
		&remaining,
	)

	if !bool(ret) {
		return "", 0, 0, false
	}

	var message string

	if messageView.data != nil {
		message = C.GoStringN(messageView.data, C.int(messageView.len))
	}

	return message, int(time), int(remaining), true
}

func (p *Player) SendDeathMessage(killer, killee *Player, weapon int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var kPtr, kePtr unsafe.Pointer

	if killer != nil && killer.ptr != nil && killer.GetID() != -1 {
		kPtr = killer.ptr
	}

	if killee != nil && killee.ptr != nil && killee.GetID() != -1 {
		kePtr = killee.ptr
	}

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SendDeathMessage(p.ptr, kPtr, kePtr, C.int(weapon)))
}

func (p *Player) SendMessageToPlayer(sender *Player, message string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() || sender == nil || sender.ptr == nil || sender.GetID() == -1 {
		return false
	}

	cMsg := C.CString(message)
	defer C.free(unsafe.Pointer(cMsg))

	return bool(C.Player_SendMessageToPlayer(p.ptr, sender.ptr, cMsg))
}

func (p *Player) GetSkillLevel(skill int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetSkillLevel(p.ptr, C.int(skill)))
}

func (p *Player) GetZAim() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0.0
	}

	return float32(C.Player_GetZAim(p.ptr))
}

func (p *Player) GetSurfingOffsets() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, 0, false
	}

	var x, y, z C.float
	ret := C.Player_GetSurfingOffsets(p.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (p *Player) GetRotationQuat() (float32, float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, 0, 0, false
	}

	var x, y, z, w C.float
	ret := C.Player_GetRotationQuat(p.ptr, &x, &y, &z, &w)

	return float32(x), float32(y), float32(z), float32(w), bool(ret)
}

func (p *Player) GetPlayerSpectateID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetPlayerSpectateID(p.ptr))
}

func (p *Player) GetSpectateType() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetSpectateType(p.ptr))
}

func (p *Player) GetRawIp() uint32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return uint32(C.Player_GetRawIp(p.ptr))
}

func (p *Player) SetGravity(gravity float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetGravity(p.ptr, C.float(gravity)))
}

func (p *Player) GetGravity() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0.0
	}

	return float32(C.Player_GetGravity(p.ptr))
}

func (p *Player) IsAdmin() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsAdmin(p.ptr))
}

func (p *Player) SetAdmin(set bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetAdmin(p.ptr, C.bool(set)))
}

func (p *Player) Kick() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	fmt.Println("PLAYER ID?", p.GetID())

	if !p.isValid() {
		return false
	}

	return bool(C.Player_Kick(p.ptr))
}

func (p *Player) Ban() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_Ban(p.ptr))
}

func (p *Player) BanEx(reason string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	cReason := C.CString(reason)
	defer C.free(unsafe.Pointer(cReason))

	return bool(C.Player_BanEx(p.ptr, cReason))
}

func (p *Player) Spawn() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_Spawn(p.ptr))
}

func (p *Player) IsSpawned() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsSpawned(p.ptr))
}

func (p *Player) IsControllable() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsControllable(p.ptr))
}

func (p *Player) IsCameraTargetEnabled() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsCameraTargetEnabled(p.ptr))
}

func (p *Player) ToggleGhostMode(toggle bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_ToggleGhostMode(p.ptr, C.bool(toggle)))
}

func (p *Player) GetGhostMode() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_GetGhostMode(p.ptr))
}

func (p *Player) AllowWeapons(allow bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_AllowWeapons(p.ptr, C.bool(allow)))
}

func (p *Player) AreWeaponsAllowed() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_AreWeaponsAllowed(p.ptr))
}

func (p *Player) IsPlayerUsingOfficialClient() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsPlayerUsingOfficialClient(p.ptr))
}

func (p *Player) GetAnimationFlags() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetAnimationFlags(p.ptr))
}

func (p *Player) IsInDriveByMode() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsInDriveByMode(p.ptr))
}

func (p *Player) IsCuffed() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsCuffed(p.ptr))
}

func (p *Player) IsUsingOmp() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsUsingOmp(p.ptr))
}

func (p *Player) IsInModShop() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_IsInModShop(p.ptr))
}

func (p *Player) GetSirenState() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetSirenState(p.ptr))
}

func (p *Player) GetLandingGearState() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_GetLandingGearState(p.ptr))
}

func (p *Player) GetHydraReactorAngle() uint32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return uint32(C.Player_GetHydraReactorAngle(p.ptr))
}

func (p *Player) GetTrainSpeed() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0.0
	}

	return float32(C.Player_GetTrainSpeed(p.ptr))
}

func (p *Player) NetStatsBytesReceived() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_NetStatsBytesReceived(p.ptr))
}

func (p *Player) NetStatsBytesSent() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_NetStatsBytesSent(p.ptr))
}

func (p *Player) NetStatsConnectionStatus() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_NetStatsConnectionStatus(p.ptr))
}

func (p *Player) NetStatsGetConnectedTime() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_NetStatsGetConnectedTime(p.ptr))
}

func (p *Player) NetStatsPacketLossPercent() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0.0
	}

	return float32(C.Player_NetStatsPacketLossPercent(p.ptr))
}

func (p *Player) NetStatsMessagesReceived() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_NetStatsMessagesReceived(p.ptr))
}

func (p *Player) NetStatsMessagesRecvPerSecond() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_NetStatsMessagesRecvPerSecond(p.ptr))
}

func (p *Player) NetStatsMessagesSent() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0
	}

	return int(C.Player_NetStatsMessagesSent(p.ptr))
}

func (p *Player) NetStatsGetIpPort() (string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return "", false
	}

	var ipPortBuffer C.struct_CAPIStringBuffer
	ret := C.Player_NetStatsGetIpPort(p.ptr, &ipPortBuffer)

	if !bool(ret) || ipPortBuffer.data == nil {
		return "", false
	}

	return C.GoStringN(ipPortBuffer.data, C.int(ipPortBuffer.len)), true
}

func (p *Player) SetFacingAngle(angle float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return false
	}

	return bool(C.Player_SetFacingAngle(p.ptr, C.float(angle)))
}

func (p *Player) GetIp() (string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return "", false
	}

	var ipBuffer C.struct_CAPIStringBuffer
	ret := C.Player_GetIp(p.ptr, &ipBuffer)

	if ret == 0 || ipBuffer.data == nil {
		return "", false
	}

	return C.GoStringN(ipBuffer.data, C.int(ipBuffer.len)), true
}

func (p *Player) GPCI() (string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return "", false
	}

	var gpciView C.struct_CAPIStringView
	ret := C.Player_GPCI(p.ptr, &gpciView)

	if !bool(ret) || gpciView.data == nil {
		return "", false
	}

	return C.GoStringN(gpciView.data, C.int(gpciView.len)), true
}

func (p *Player) GetVersion() (string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return "", false
	}

	var versionView C.struct_CAPIStringView
	ret := C.Player_GetVersion(p.ptr, &versionView)

	if ret == 0 || versionView.data == nil {
		return "", false
	}

	return C.GoStringN(versionView.data, C.int(versionView.len)), true
}

func (p *Player) GetMenu() unsafe.Pointer {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return C.Player_GetMenu(p.ptr)
}

func (p *Player) GetSurfingPlayerObject() unsafe.Pointer {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return C.Player_GetSurfingPlayerObject(p.ptr)
}

func (p *Player) GetCameraTargetPlayerObject() unsafe.Pointer {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return nil
	}

	return C.Player_GetCameraTargetPlayerObject(p.ptr)
}

func (p *Player) GetAttachedObject(index int) (int, int, float32, float32, float32, float32, float32, float32, float32, float32, float32, int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, false
	}

	var model, bone C.int
	var ox, oy, oz, rx, ry, rz, sx, sy, sz C.float
	var c1, c2 C.int

	ret := C.Player_GetAttachedObject(
		p.ptr,
		C.int(index),
		&model,
		&bone,
		&ox, &oy, &oz,
		&rx, &ry, &rz,
		&sx, &sy, &sz,
		&c1,
		&c2,
	)

	return int(model), int(bone),
		float32(ox), float32(oy), float32(oz),
		float32(rx), float32(ry), float32(rz),
		float32(sx), float32(sy), float32(sz),
		int(c1), int(c2), bool(ret)
}

func (p *Player) GetDialogData() (int, int, string, string, string, string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !p.isValid() {
		return 0, 0, "", "", "", "", false
	}

	var dialogID, style C.int
	var titleView, bodyView, b1View, b2View C.struct_CAPIStringView

	ret := C.Player_GetDialogData(
		p.ptr,
		&dialogID,
		&style,
		&titleView,
		&bodyView,
		&b1View,
		&b2View,
	)

	if !bool(ret) {
		return 0, 0, "", "", "", "", false
	}

	title := C.GoStringN(titleView.data, C.int(titleView.len))
	body := C.GoStringN(bodyView.data, C.int(bodyView.len))
	b1 := C.GoStringN(b1View.data, C.int(b1View.len))
	b2 := C.GoStringN(b2View.data, C.int(b2View.len))

	return int(dialogID), int(style), title, body, b1, b2, true
}
