package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "npc.h"

#endif
*/
import "C"
import "unsafe"

type NPC struct {
	ptr unsafe.Pointer
}

func npcFromPointer(ptr unsafe.Pointer) *NPC {
	if ptr == nil {
		return nil
	}

	return &NPC{ptr}
}

func NPCFromID(npcid int) *NPC {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	ptr := C.NPC_FromID(C.int(npcid))

	if ptr == nil {
		return nil
	}

	return &NPC{ptr}
}

func NPCCreate(name string, id *int) *NPC {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var cID C.int
	ptr := C.NPC_Create(cName, &cID)

	if id != nil {
		*id = int(cID)
	}

	return npcFromPointer(ptr)
}

func NPCConnect(name, script string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cScript := C.CString(script)
	defer C.free(unsafe.Pointer(cScript))

	return bool(C.NPC_Connect(cName, cScript))
}

func NPCGetAll(maxNPCs int) []int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	arr := make([]C.int, maxNPCs)
	count := int(C.NPC_GetAll(&arr[0], C.int(maxNPCs)))

	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = int(arr[i])
	}

	return result
}

func NPCCreatePath() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.NPC_CreatePath())
}

func NPCDestroyPath(pathId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_DestroyPath(C.int(pathId)))
}

func NPCDestroyAllPath() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_DestroyAllPath())
}

func NPCGetPathCount() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.NPC_GetPathCount())
}

func NPCAddPointToPath(pathId int, x, y, z, stopRange float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_AddPointToPath(C.int(pathId), C.float(x), C.float(y), C.float(z), C.float(stopRange)))
}

func NPCRemovePointFromPath(pathId, pointIndex int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_RemovePointFromPath(C.int(pathId), C.int(pointIndex)))
}

func NPCClearPath(pathId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_ClearPath(C.int(pathId)))
}

func NPCGetPathPointCount(pathId int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.NPC_GetPathPointCount(C.int(pathId)))
}

func NPCGetPathPoint(pathId, pointIndex int) (float32, float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var x, y, z, stopRange C.float
	ret := C.NPC_GetPathPoint(C.int(pathId), C.int(pointIndex), &x, &y, &z, &stopRange)

	return float32(x), float32(y), float32(z), float32(stopRange), bool(ret)
}

func NPCIsValidPath(pathId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_IsValidPath(C.int(pathId)))
}

func NPCHasPathPointInRange(pathId int, x, y, z, radius float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_HasPathPointInRange(C.int(pathId), C.float(x), C.float(y), C.float(z), C.float(radius)))
}

func NPCLoadRecord(filePath string) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	cPath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cPath))

	return int(C.NPC_LoadRecord(cPath))
}

func NPCUnloadRecord(recordId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_UnloadRecord(C.int(recordId)))
}

func NPCIsValidRecord(recordId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_IsValidRecord(C.int(recordId)))
}

func NPCGetRecordCount() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.NPC_GetRecordCount())
}

func NPCUnloadAllRecords() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_UnloadAllRecords())
}

func NPCOpenNode(nodeId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_OpenNode(C.int(nodeId)))
}

func NPCCloseNode(nodeId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_CloseNode(C.int(nodeId)))
}

func NPCIsNodeOpen(nodeId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_IsNodeOpen(C.int(nodeId)))
}

func NPCGetNodeType(nodeId int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.NPC_GetNodeType(C.int(nodeId)))
}

func NPCSetNodePoint(nodeId, pointId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.NPC_SetNodePoint(C.int(nodeId), C.int(pointId)))
}

func NPCGetNodePointPosition(nodeId int) (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var x, y, z C.float
	ret := C.NPC_GetNodePointPosition(C.int(nodeId), &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func NPCGetNodePointCount(nodeId int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.NPC_GetNodePointCount(C.int(nodeId)))
}

func NPCGetNodeInfo(nodeId int) (uint32, uint32, uint32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var vehicleNodes, pedNodes, naviNodes C.uint32_t
	ret := C.NPC_GetNodeInfo(C.int(nodeId), &vehicleNodes, &pedNodes, &naviNodes)

	return uint32(vehicleNodes), uint32(pedNodes), uint32(naviNodes), bool(ret)
}

func (n *NPC) isValid() bool {
	if n == nil || n.ptr == nil {
		return false
	}

	return bool(C.NPC_IsValid(n.ptr))
}

func (n *NPC) GetID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if n == nil || n.ptr == nil {
		return -1
	}

	return int(C.NPC_GetID(n.ptr))
}

func (n *NPC) IsValid() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return n.isValid()
}

func (n *NPC) Destroy() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_Destroy(n.ptr))
}

func (n *NPC) GetPlayer() *Player {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return nil
	}

	return playerFromPointer(C.NPC_GetPlayer(n.ptr))
}

func (n *NPC) Spawn() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_Spawn(n.ptr))
}

func (n *NPC) Respawn() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_Respawn(n.ptr))
}

func (n *NPC) SetPos(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetPos(n.ptr, C.float(x), C.float(y), C.float(z)))
}

func (n *NPC) GetPos() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0, 0, 0, false
	}

	var x, y, z C.float
	ret := C.NPC_GetPos(n.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (n *NPC) SetRot(rx, ry, rz float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetRot(n.ptr, C.float(rx), C.float(ry), C.float(rz)))
}

func (n *NPC) GetRot() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0, 0, 0, false
	}

	var rx, ry, rz C.float
	ret := C.NPC_GetRot(n.ptr, &rx, &ry, &rz)

	return float32(rx), float32(ry), float32(rz), bool(ret)
}

func (n *NPC) SetFacingAngle(angle float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetFacingAngle(n.ptr, C.float(angle)))
}

func (n *NPC) GetFacingAngle() (float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0, false
	}

	var angle C.float
	ret := C.NPC_GetFacingAngle(n.ptr, &angle)

	return float32(angle), bool(ret)
}

func (n *NPC) SetVirtualWorld(virtualWorld int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetVirtualWorld(n.ptr, C.int(virtualWorld)))
}

func (n *NPC) GetVirtualWorld() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetVirtualWorld(n.ptr))
}

func (n *NPC) SetInterior(interior int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetInterior(n.ptr, C.int(interior)))
}

func (n *NPC) GetInterior() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetInterior(n.ptr))
}

func (n *NPC) Move(x, y, z float32, moveType int, moveSpeed, stopRange float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_Move(n.ptr, C.float(x), C.float(y), C.float(z), C.int(moveType), C.float(moveSpeed), C.float(stopRange)))
}

func (n *NPC) MoveToPlayer(player *Player, moveType int, moveSpeed, stopRange float32, posCheckUpdateDelay int, autoRestart bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.NPC_MoveToPlayer(n.ptr, player.ptr, C.int(moveType), C.float(moveSpeed), C.float(stopRange), C.int(posCheckUpdateDelay), C.bool(autoRestart)))
}

func (n *NPC) StopMove() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_StopMove(n.ptr))
}

func (n *NPC) IsMoving() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsMoving(n.ptr))
}

func (n *NPC) SetSkin(model int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetSkin(n.ptr, C.int(model)))
}

func (n *NPC) IsStreamedIn(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.NPC_IsStreamedIn(n.ptr, player.ptr))
}

func (n *NPC) IsAnyStreamedIn() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsAnyStreamedIn(n.ptr))
}

func (n *NPC) SetHealth(health float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetHealth(n.ptr, C.float(health)))
}

func (n *NPC) GetHealth() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0.0
	}

	return float32(C.NPC_GetHealth(n.ptr))
}

func (n *NPC) SetArmour(armour float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetArmour(n.ptr, C.float(armour)))
}

func (n *NPC) GetArmour() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0.0
	}

	return float32(C.NPC_GetArmour(n.ptr))
}

func (n *NPC) IsDead() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsDead(n.ptr))
}

func (n *NPC) SetInvulnerable(toggle bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetInvulnerable(n.ptr, C.bool(toggle)))
}

func (n *NPC) IsInvulnerable() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsInvulnerable(n.ptr))
}

func (n *NPC) SetWeapon(weapon uint8) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetWeapon(n.ptr, C.uint8_t(weapon)))
}

func (n *NPC) GetWeapon() uint8 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return uint8(C.NPC_GetWeapon(n.ptr))
}

func (n *NPC) SetAmmo(ammo int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetAmmo(n.ptr, C.int(ammo)))
}

func (n *NPC) GetAmmo() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetAmmo(n.ptr))
}

func (n *NPC) SetAmmoInClip(ammo int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetAmmoInClip(n.ptr, C.int(ammo)))
}

func (n *NPC) GetAmmoInClip() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetAmmoInClip(n.ptr))
}

func (n *NPC) EnableReloading(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_EnableReloading(n.ptr, C.bool(enable)))
}

func (n *NPC) IsReloadEnabled() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsReloadEnabled(n.ptr))
}

func (n *NPC) IsReloading() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsReloading(n.ptr))
}

func (n *NPC) EnableInfiniteAmmo(enable bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_EnableInfiniteAmmo(n.ptr, C.bool(enable)))
}

func (n *NPC) IsInfiniteAmmoEnabled() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsInfiniteAmmoEnabled(n.ptr))
}

func (n *NPC) GetWeaponState() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetWeaponState(n.ptr))
}

func (n *NPC) SetKeys(upAndDown, leftAndRight, keys uint16) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetKeys(n.ptr, C.uint16_t(upAndDown), C.uint16_t(leftAndRight), C.uint16_t(keys)))
}

func (n *NPC) GetKeys() (uint16, uint16, uint16, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0, 0, 0, false
	}

	var upAndDown, leftAndRight, keys C.uint16_t
	ret := C.NPC_GetKeys(n.ptr, &upAndDown, &leftAndRight, &keys)

	return uint16(upAndDown), uint16(leftAndRight), uint16(keys), bool(ret)
}

func (n *NPC) SetWeaponSkillLevel(skill uint8, level int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetWeaponSkillLevel(n.ptr, C.uint8_t(skill), C.int(level)))
}

func (n *NPC) GetWeaponSkillLevel(skill int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetWeaponSkillLevel(n.ptr, C.int(skill)))
}

func (n *NPC) MeleeAttack(time int, secondaryAttack bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_MeleeAttack(n.ptr, C.int(time), C.bool(secondaryAttack)))
}

func (n *NPC) StopMeleeAttack() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_StopMeleeAttack(n.ptr))
}

func (n *NPC) IsMeleeAttacking() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsMeleeAttacking(n.ptr))
}

func (n *NPC) SetFightingStyle(style int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetFightingStyle(n.ptr, C.int(style)))
}

func (n *NPC) GetFightingStyle() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetFightingStyle(n.ptr))
}

func (n *NPC) Shoot(weapon uint8, hitId, hitType int, endX, endY, endZ, offsetX, offsetY, offsetZ float32, isHit bool, checkInBetweenFlags uint8) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_Shoot(n.ptr, C.uint8_t(weapon), C.int(hitId), C.int(hitType), C.float(endX), C.float(endY), C.float(endZ), C.float(offsetX), C.float(offsetY), C.float(offsetZ), C.bool(isHit), C.uint8_t(checkInBetweenFlags)))
}

func (n *NPC) IsShooting() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsShooting(n.ptr))
}

func (n *NPC) AimAt(x, y, z float32, shoot bool, shootDelay int, updateAngle bool, offsetFromX, offsetFromY, offsetFromZ float32, checkInBetweenFlags uint8) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_AimAt(n.ptr, C.float(x), C.float(y), C.float(z), C.bool(shoot), C.int(shootDelay), C.bool(updateAngle), C.float(offsetFromX), C.float(offsetFromY), C.float(offsetFromZ), C.uint8_t(checkInBetweenFlags)))
}

func (n *NPC) AimAtPlayer(atPlayer *Player, shoot bool, shootDelay int, updateAngle bool, offsetX, offsetY, offsetZ, offsetFromX, offsetFromY, offsetFromZ float32, checkInBetweenFlags uint8) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() || atPlayer == nil || atPlayer.ptr == nil {
		return false
	}

	return bool(C.NPC_AimAtPlayer(n.ptr, atPlayer.ptr, C.bool(shoot), C.int(shootDelay), C.bool(updateAngle), C.float(offsetX), C.float(offsetY), C.float(offsetZ), C.float(offsetFromX), C.float(offsetFromY), C.float(offsetFromZ), C.uint8_t(checkInBetweenFlags)))
}

func (n *NPC) StopAim() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_StopAim(n.ptr))
}

func (n *NPC) IsAiming() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsAiming(n.ptr))
}

func (n *NPC) IsAimingAtPlayer(atPlayer *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() || atPlayer == nil || atPlayer.ptr == nil {
		return false
	}

	return bool(C.NPC_IsAimingAtPlayer(n.ptr, atPlayer.ptr))
}

func (n *NPC) SetWeaponAccuracy(weapon int, accuracy float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetWeaponAccuracy(n.ptr, C.int(weapon), C.float(accuracy)))
}

func (n *NPC) GetWeaponAccuracy(weapon int) float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0.0
	}

	return float32(C.NPC_GetWeaponAccuracy(n.ptr, C.int(weapon)))
}

func (n *NPC) SetWeaponReloadTime(weapon, time int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetWeaponReloadTime(n.ptr, C.int(weapon), C.int(time)))
}

func (n *NPC) GetWeaponReloadTime(weapon int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetWeaponReloadTime(n.ptr, C.int(weapon)))
}

func (n *NPC) GetWeaponActualReloadTime(weapon int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetWeaponActualReloadTime(n.ptr, C.int(weapon)))
}

func (n *NPC) SetWeaponShootTime(weapon, time int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetWeaponShootTime(n.ptr, C.int(weapon), C.int(time)))
}

func (n *NPC) GetWeaponShootTime(weapon int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetWeaponShootTime(n.ptr, C.int(weapon)))
}

func (n *NPC) SetWeaponClipSize(weapon, size int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetWeaponClipSize(n.ptr, C.int(weapon), C.int(size)))
}

func (n *NPC) GetWeaponClipSize(weapon int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetWeaponClipSize(n.ptr, C.int(weapon)))
}

func (n *NPC) GetWeaponActualClipSize(weapon int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetWeaponActualClipSize(n.ptr, C.int(weapon)))
}

func (n *NPC) EnterVehicle(vehicle *Vehicle, seatId, moveType int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() || vehicle == nil || vehicle.ptr == nil {
		return false
	}

	return bool(C.NPC_EnterVehicle(n.ptr, vehicle.ptr, C.int(seatId), C.int(moveType)))
}

func (n *NPC) ExitVehicle() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_ExitVehicle(n.ptr))
}

func (n *NPC) PutInVehicle(vehicle *Vehicle, seatId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() || vehicle == nil || vehicle.ptr == nil {
		return false
	}

	return bool(C.NPC_PutInVehicle(n.ptr, vehicle.ptr, C.int(seatId)))
}

func (n *NPC) RemoveFromVehicle() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_RemoveFromVehicle(n.ptr))
}

func (n *NPC) GetVehicle() *Vehicle {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return nil
	}

	return vehicleFromPointer(C.NPC_GetVehicle(n.ptr))
}

func (n *NPC) GetVehicleID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return -1
	}

	return int(C.NPC_GetVehicleID(n.ptr))
}

func (n *NPC) GetEnteringVehicle() *Vehicle {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return nil
	}

	return vehicleFromPointer(C.NPC_GetEnteringVehicle(n.ptr))
}

func (n *NPC) GetEnteringVehicleID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return -1
	}

	return int(C.NPC_GetEnteringVehicleID(n.ptr))
}

func (n *NPC) GetVehicleSeat() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return -1
	}

	return int(C.NPC_GetVehicleSeat(n.ptr))
}

func (n *NPC) GetEnteringVehicleSeat() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return -1
	}

	return int(C.NPC_GetEnteringVehicleSeat(n.ptr))
}

func (n *NPC) IsEnteringVehicle() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsEnteringVehicle(n.ptr))
}

func (n *NPC) UseVehicleSiren(use bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_UseVehicleSiren(n.ptr, C.bool(use)))
}

func (n *NPC) IsVehicleSirenUsed() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsVehicleSirenUsed(n.ptr))
}

func (n *NPC) SetVehicleHealth(health float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetVehicleHealth(n.ptr, C.float(health)))
}

func (n *NPC) GetVehicleHealth() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0.0
	}

	return float32(C.NPC_GetVehicleHealth(n.ptr))
}

func (n *NPC) SetVehicleHydraThrusters(direction int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetVehicleHydraThrusters(n.ptr, C.int(direction)))
}

func (n *NPC) GetVehicleHydraThrusters() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetVehicleHydraThrusters(n.ptr))
}

func (n *NPC) SetVehicleGearState(gearState int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetVehicleGearState(n.ptr, C.int(gearState)))
}

func (n *NPC) GetVehicleGearState() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetVehicleGearState(n.ptr))
}

func (n *NPC) SetVehicleTrainSpeed(speed float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetVehicleTrainSpeed(n.ptr, C.float(speed)))
}

func (n *NPC) GetVehicleTrainSpeed() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0.0
	}

	return float32(C.NPC_GetVehicleTrainSpeed(n.ptr))
}

func (n *NPC) GetCurrentPathPointIndex() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return -1
	}

	return int(C.NPC_GetCurrentPathPointIndex(n.ptr))
}

func (n *NPC) MoveByPath(pathId, moveType int, moveSpeed float32, reverse bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_MoveByPath(n.ptr, C.int(pathId), C.int(moveType), C.float(moveSpeed), C.bool(reverse)))
}

func (n *NPC) ResetAnimation() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_ResetAnimation(n.ptr))
}

func (n *NPC) SetAnimation(animationId int, delta float32, loop, lockX, lockY, freeze bool, time int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetAnimation(n.ptr, C.int(animationId), C.float(delta), C.bool(loop), C.bool(lockX), C.bool(lockY), C.bool(freeze), C.int(time)))
}

func (n *NPC) GetAnimation() (int, float32, bool, bool, bool, bool, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0, 0, false, false, false, false, 0, false
	}

	var animationId, time C.int
	var delta C.float
	var loop, lockX, lockY, freeze C.bool
	ret := C.NPC_GetAnimation(n.ptr, &animationId, &delta, &loop, &lockX, &lockY, &freeze, &time)

	return int(animationId), float32(delta), bool(loop), bool(lockX), bool(lockY), bool(freeze), int(time), bool(ret)
}

func (n *NPC) ApplyAnimation(animlib, animname string, delta float32, loop, lockX, lockY, freeze bool, time int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	cAnimlib := C.CString(animlib)
	defer C.free(unsafe.Pointer(cAnimlib))

	cAnimname := C.CString(animname)
	defer C.free(unsafe.Pointer(cAnimname))

	return bool(C.NPC_ApplyAnimation(n.ptr, cAnimlib, cAnimname, C.float(delta), C.bool(loop), C.bool(lockX), C.bool(lockY), C.bool(freeze), C.int(time)))
}

func (n *NPC) ClearAnimations() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_ClearAnimations(n.ptr))
}

func (n *NPC) SetSpecialAction(action int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetSpecialAction(n.ptr, C.int(action)))
}

func (n *NPC) GetSpecialAction() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0
	}

	return int(C.NPC_GetSpecialAction(n.ptr))
}

func (n *NPC) StartPlayback(recordName string, autoUnload bool, startPosX, startPosY, startPosZ, startRotX, startRotY, startRotZ float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	cName := C.CString(recordName)
	defer C.free(unsafe.Pointer(cName))

	return bool(C.NPC_StartPlayback(n.ptr, cName, C.bool(autoUnload), C.float(startPosX), C.float(startPosY), C.float(startPosZ), C.float(startRotX), C.float(startRotY), C.float(startRotZ)))
}

func (n *NPC) StartPlaybackEx(recordId int, autoUnload bool, startPosX, startPosY, startPosZ, startRotX, startRotY, startRotZ float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_StartPlaybackEx(n.ptr, C.int(recordId), C.bool(autoUnload), C.float(startPosX), C.float(startPosY), C.float(startPosZ), C.float(startRotX), C.float(startRotY), C.float(startRotZ)))
}

func (n *NPC) StopPlayback() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_StopPlayback(n.ptr))
}

func (n *NPC) PausePlayback(paused bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_PausePlayback(n.ptr, C.bool(paused)))
}

func (n *NPC) IsPlayingPlayback() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsPlayingPlayback(n.ptr))
}

func (n *NPC) IsPlaybackPaused() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsPlaybackPaused(n.ptr))
}

func (n *NPC) PlayNode(nodeId, moveType int, moveSpeed, radius float32, setAngle bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_PlayNode(n.ptr, C.int(nodeId), C.int(moveType), C.float(moveSpeed), C.float(radius), C.bool(setAngle)))
}

func (n *NPC) StopPlayingNode() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_StopPlayingNode(n.ptr))
}

func (n *NPC) PausePlayingNode() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_PausePlayingNode(n.ptr))
}

func (n *NPC) ResumePlayingNode() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_ResumePlayingNode(n.ptr))
}

func (n *NPC) IsPlayingNodePaused() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsPlayingNodePaused(n.ptr))
}

func (n *NPC) IsPlayingNode() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_IsPlayingNode(n.ptr))
}

func (n *NPC) ChangeNode(nodeId, linkId int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return -1
	}

	return int(C.NPC_ChangeNode(n.ptr, C.int(nodeId), C.int(linkId)))
}

func (n *NPC) UpdateNodePoint(pointId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_UpdateNodePoint(n.ptr, C.int(pointId)))
}

func (n *NPC) SetSurfingOffset(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetSurfingOffset(n.ptr, C.float(x), C.float(y), C.float(z)))
}

func (n *NPC) GetSurfingOffset() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return 0, 0, 0, false
	}

	var x, y, z C.float
	ret := C.NPC_GetSurfingOffset(n.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (n *NPC) SetSurfingVehicle(vehicle *Vehicle) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() || vehicle == nil || vehicle.ptr == nil {
		return false
	}

	return bool(C.NPC_SetSurfingVehicle(n.ptr, vehicle.ptr))
}

func (n *NPC) GetSurfingVehicle() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return -1
	}

	return int(C.NPC_GetSurfingVehicle(n.ptr))
}

func (n *NPC) SetSurfingObject(object unsafe.Pointer) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_SetSurfingObject(n.ptr, object))
}

func (n *NPC) GetSurfingObject() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return -1
	}

	return int(C.NPC_GetSurfingObject(n.ptr))
}

func (n *NPC) SetSurfingPlayerObject(player *Player, objectId int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.NPC_SetSurfingPlayerObject(n.ptr, player.ptr, C.int(objectId)))
}

func (n *NPC) GetSurfingPlayerObject() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return -1
	}

	return int(C.NPC_GetSurfingPlayerObject(n.ptr))
}

func (n *NPC) ResetSurfingData() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if !n.isValid() {
		return false
	}

	return bool(C.NPC_ResetSurfingData(n.ptr))
}
