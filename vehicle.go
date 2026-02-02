package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "vehicle.h"

#endif
*/
import "C"
import (
	"unsafe"
)

type Vehicle struct {
	ptr unsafe.Pointer
}

func vehicleFromPointer(ptr unsafe.Pointer) *Vehicle {
	if ptr == nil {
		return nil
	}

	return &Vehicle{ptr}
}

func VehicleFromID(vehicleID int) *Vehicle {
	ptr := C.Vehicle_FromID(C.int(vehicleID))

	if ptr == nil {
		return nil
	}

	return &Vehicle{ptr}
}

func VehicleCreate(modelID int, x, y, z, rotation float32, color1, color2, respawnDelay int, addSiren bool) *Vehicle {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var vehicleID C.int

	ptr := C.Vehicle_Create(
		C.int(modelID),
		C.float(x),
		C.float(y),
		C.float(z),
		C.float(rotation),
		C.int(color1),
		C.int(color2),
		C.int(respawnDelay),
		C.bool(addSiren),
		&vehicleID,
	)

	if ptr == nil {
		return nil
	}

	return &Vehicle{
		ptr,
	}
}

func VehicleAddStatic(modelid int, x, y, z, angle float32, color1, color2 int) *Vehicle {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var vehicleID C.int

	ptr := C.Vehicle_AddStatic(
		C.int(modelid),
		C.float(x),
		C.float(y),
		C.float(z),
		C.float(angle),
		C.int(color1),
		C.int(color2),
		&vehicleID,
	)

	if ptr == nil {
		return nil
	}

	return &Vehicle{ptr: ptr}
}

func VehicleAddStaticEx(modelid int, x, y, z, angle float32, color1, color2, respawnDelay int, addSiren bool) *Vehicle {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var vehicleID C.int

	ptr := C.Vehicle_AddStaticEx(
		C.int(modelid),
		C.float(x),
		C.float(y),
		C.float(z),
		C.float(angle),
		C.int(color1),
		C.int(color2),
		C.int(respawnDelay),
		C.bool(addSiren),
		&vehicleID,
	)

	if ptr == nil {
		return nil
	}

	return &Vehicle{ptr: ptr}
}

func VehicleGetModelInfo(vehicleModel, infoType int) (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var x, y, z C.float

	ret := C.Vehicle_GetModelInfo(
		C.int(vehicleModel),
		C.int(infoType),
		&x,
		&y,
		&z,
	)

	return float32(x), float32(y), float32(z), bool(ret)
}

func VehicleGetRandomColorPair(modelID int) (int, int, int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	var c1, c2, c3, c4 C.int

	ret := C.Vehicle_GetRandomColorPair(
		C.int(modelID),
		&c1,
		&c2,
		&c3,
		&c4,
	)

	return int(c1), int(c2), int(c3), int(c4), bool(ret)
}

func VehicleColorIndexToColor(colorIndex, alpha int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Vehicle_ColorIndexToColor(C.int(colorIndex), C.int(alpha)))
}

func VehicleGetModelsUsed() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Vehicle_GetModelsUsed())
}

func VehicleGetModelCount(modelID int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Vehicle_GetModelCount(C.int(modelID)))
}

func VehicleEnableFriendlyFire() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Vehicle_EnableFriendlyFire())
}

func (v *Vehicle) Destroy() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_Destroy(v.ptr))
}

func (v *Vehicle) GetID() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {

	}

	return int(C.Vehicle_GetID(v.ptr))
}

func (v *Vehicle) GetMaxPassengerSeats() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetMaxPassengerSeats(C.Vehicle_GetModel(v.ptr)))
}

func (v *Vehicle) SetParamsCarDoors(frontLeft, frontRight, rearLeft, rearRight int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetParamsCarDoors(
		v.ptr,
		C.int(frontLeft),
		C.int(frontRight),
		C.int(rearLeft),
		C.int(rearRight),
	))
}

func (v *Vehicle) GetParamsCarDoors() (int, int, int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0, 0, 0, 0, false
	}

	var fl, fr, rl, rr C.int

	ret := C.Vehicle_GetParamsCarDoors(
		v.ptr,
		&fl,
		&fr,
		&rl,
		&rr,
	)

	return int(fl), int(fr), int(rl), int(rr), bool(ret)
}

func (v *Vehicle) SetParamsCarWindows(frontLeft, frontRight, rearLeft, rearRight int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetParamsCarWindows(
		v.ptr,
		C.int(frontLeft),
		C.int(frontRight),
		C.int(rearLeft),
		C.int(rearRight),
	))
}

func (v *Vehicle) GetParamsCarWindows() (int, int, int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0, 0, 0, 0, false
	}

	var fl, fr, rl, rr C.int

	ret := C.Vehicle_GetParamsCarWindows(
		v.ptr,
		&fl,
		&fr,
		&rl,
		&rr,
	)

	return int(fl), int(fr), int(rl), int(rr), bool(ret)
}

func (v *Vehicle) GetMatrix() (float32, float32, float32, float32, float32, float32, float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, 0, false
	}

	var rx, ry, rz, ux, uy, uz, ax, ay, az C.float

	ret := C.Vehicle_GetMatrix(
		v.ptr,
		&rx, &ry, &rz,
		&ux, &uy, &uz,
		&ax, &ay, &az,
	)

	return float32(rx), float32(ry), float32(rz),
		float32(ux), float32(uy), float32(uz),
		float32(ax), float32(ay), float32(az),
		bool(ret)
}

func (v *Vehicle) GetSpawnInfo() (float32, float32, float32, float32, int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0, 0, 0, 0, 0, 0, false
	}

	var x, y, z, rot C.float
	var c1, c2 C.int

	ret := C.Vehicle_GetSpawnInfo(
		v.ptr,
		&x,
		&y,
		&z,
		&rot,
		&c1,
		&c2,
	)

	return float32(x), float32(y), float32(z), float32(rot), int(c1), int(c2), bool(ret)
}

func (v *Vehicle) SetSpawnInfo(modelID int, x, y, z, rotation float32, color1, color2, respawnTime, interior int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetSpawnInfo(
		v.ptr,
		C.int(modelID),
		C.float(x),
		C.float(y),
		C.float(z),
		C.float(rotation),
		C.int(color1),
		C.int(color2),
		C.int(respawnTime),
		C.int(interior),
	))
}

func (v *Vehicle) GetComponentType(componentID int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return int(C.Vehicle_GetComponentType(C.int(componentID)))
}

func (v *Vehicle) CanHaveComponent(modelID, componentID int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Vehicle_CanHaveComponent(C.int(modelID), C.int(componentID)))
}

func (v *Vehicle) SetToRespawn() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetToRespawn(v.ptr))
}

func (v *Vehicle) LinkToInterior(interiorID int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_LinkToInterior(v.ptr, C.int(interiorID)))
}

func (v *Vehicle) SetNumberPlate(plate string) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	cPlate := C.CString(plate)
	defer C.free(unsafe.Pointer(cPlate))

	return bool(C.Vehicle_SetNumberPlate(v.ptr, cPlate))
}

func (v *Vehicle) GetNumberPlate() (string, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return "", false
	}

	var plateView C.struct_CAPIStringView
	ret := C.Vehicle_GetNumberPlate(v.ptr, &plateView)

	if !ret || plateView.data == nil {
		return "", false
	}

	return C.GoStringN(plateView.data, C.int(plateView.len)), true
}

func (v *Vehicle) SetRespawnDelay(delay int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetRespawnDelay(v.ptr, C.int(delay)))
}

func (v *Vehicle) GetRespawnDelay() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetRespawnDelay(v.ptr))
}

func (v *Vehicle) GetRespawnTick() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetRespawnTick(v.ptr))
}

func (v *Vehicle) GetOccupiedTick() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetOccupiedTick(v.ptr))
}

func (v *Vehicle) HasBeenOccupied() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_HasBeenOccupied(v.ptr))
}

func (v *Vehicle) IsOccupied() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_IsOccupied(v.ptr))
}

func (v *Vehicle) IsDead() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return true
	}

	return bool(C.Vehicle_IsDead(v.ptr))
}

func (v *Vehicle) ToggleSirenEnabled(status bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_ToggleSirenEnabled(v.ptr, C.bool(status)))
}

func (v *Vehicle) IsSirenEnabled() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_IsSirenEnabled(v.ptr))
}

func (v *Vehicle) SetParamsSirenState(enabled bool) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetParamsSirenState(v.ptr, C.bool(enabled)))
}

func (v *Vehicle) GetParamsSirenState() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetParamsSirenState(v.ptr))
}

func (v *Vehicle) GetSirenState() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetSirenState(v.ptr))
}

func (v *Vehicle) GetHydraReactorAngle() uint32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return uint32(C.Vehicle_GetHydraReactorAngle(v.ptr))
}

func (v *Vehicle) GetLandingGearState() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetLandingGearState(v.ptr))
}

func (v *Vehicle) GetTrainSpeed() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0.0
	}

	return float32(C.Vehicle_GetTrainSpeed(v.ptr))
}

func (v *Vehicle) GetLastDriver() *Player {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return nil
	}

	return playerFromPointer(C.Vehicle_GetLastDriver(v.ptr))
}

func (v *Vehicle) GetDriver() *Player {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return nil
	}

	return playerFromPointer(C.Vehicle_GetDriver(v.ptr))
}

func (v *Vehicle) GetOccupant(seat int) *Player {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return nil
	}

	return playerFromPointer(C.Vehicle_GetOccupant(v.ptr, C.int(seat)))
}

func (v *Vehicle) CountOccupants() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_CountOccupants(v.ptr))
}

func (v *Vehicle) GetCab() *Vehicle {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return nil
	}

	ptr := C.Vehicle_GetCab(v.ptr)

	if ptr == nil {
		return nil
	}

	return &Vehicle{ptr: ptr}
}

func (v *Vehicle) GetTower() *Vehicle {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return nil
	}

	ptr := C.Vehicle_GetTower(v.ptr)

	if ptr == nil {
		return nil
	}

	return &Vehicle{ptr: ptr}
}

func (v *Vehicle) GetTrailer() *Vehicle {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return nil
	}

	ptr := C.Vehicle_GetTrailer(v.ptr)

	if ptr == nil {
		return nil
	}

	return &Vehicle{ptr: ptr}
}

func (v *Vehicle) AttachTrailer(trailer *Vehicle) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil || trailer == nil || trailer.ptr == nil {
		return false
	}

	return bool(C.Vehicle_AttachTrailer(trailer.ptr, v.ptr))
}

func (v *Vehicle) DetachTrailer() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_DetachTrailer(v.ptr))
}

func (v *Vehicle) IsTrailerAttached() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_IsTrailerAttached(v.ptr))
}

func (v *Vehicle) Repair() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_Repair(v.ptr))
}

func (v *Vehicle) AddComponent(componentID int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_AddComponent(v.ptr, C.int(componentID)))
}

func (v *Vehicle) RemoveComponent(componentID int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_RemoveComponent(v.ptr, C.int(componentID)))
}

func (v *Vehicle) GetComponentInSlot(slot int) int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetComponentInSlot(v.ptr, C.int(slot)))
}

func (v *Vehicle) ChangePaintjob(paintjobID int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_ChangePaintjob(v.ptr, C.int(paintjobID)))
}

func (v *Vehicle) GetPaintjob() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetPaintjob(v.ptr))
}

func (v *Vehicle) ChangeColor(color1, color2 int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_ChangeColor(v.ptr, C.int(color1), C.int(color2)))
}

func (v *Vehicle) GetColor() (int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0, 0, false
	}

	var c1, c2 C.int

	ret := C.Vehicle_GetColor(v.ptr, &c1, &c2)

	return int(c1), int(c2), bool(ret)
}

func (v *Vehicle) UpdateDamageStatus(panels, doors, lights, tires int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_UpdateDamageStatus(
		v.ptr,
		C.int(panels),
		C.int(doors),
		C.int(lights),
		C.int(tires),
	))
}

func (v *Vehicle) GetDamageStatus() (int, int, int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0, 0, 0, 0, false
	}

	var panels, doors, lights, tires C.int

	ret := C.Vehicle_GetDamageStatus(
		v.ptr,
		&panels,
		&doors,
		&lights,
		&tires,
	)

	return int(panels), int(doors), int(lights), int(tires), bool(ret)
}

func (v *Vehicle) UseManualEngineAndLights() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	return bool(C.Vehicle_UseManualEngineAndLights())
}

func (v *Vehicle) IsValid() bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_IsValid(v.ptr))
}

func (v *Vehicle) GetModel() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetModel(v.ptr))
}

func (v *Vehicle) IsStreamedIn(player *Player) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.Vehicle_IsStreamedIn(v.ptr, player.ptr))
}

func (v *Vehicle) GetPos() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0, 0, 0, false
	}

	var x, y, z C.float

	ret := C.Vehicle_GetPos(v.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (v *Vehicle) SetPos(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetPos(v.ptr, C.float(x), C.float(y), C.float(z)))
}

func (v *Vehicle) GetZAngle() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0.0
	}

	return float32(C.Vehicle_GetZAngle(v.ptr))
}

func (v *Vehicle) SetZAngle(angle float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetZAngle(v.ptr, C.float(angle)))
}

func (v *Vehicle) GetRotationQuat() (float32, float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0, 0, 0, 0, false
	}

	var w, x, y, z C.float

	ret := C.Vehicle_GetRotationQuat(v.ptr, &w, &x, &y, &z)

	return float32(w), float32(x), float32(y), float32(z), bool(ret)
}

func (v *Vehicle) GetDistanceFromPoint(x, y, z float32) float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0.0
	}

	return float32(C.Vehicle_GetDistanceFromPoint(v.ptr, C.float(x), C.float(y), C.float(z)))
}

func (v *Vehicle) SetVelocity(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetVelocity(v.ptr, C.float(x), C.float(y), C.float(z)))
}

func (v *Vehicle) GetVelocity() (float32, float32, float32, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0, 0, 0, false
	}

	var x, y, z C.float

	ret := C.Vehicle_GetVelocity(v.ptr, &x, &y, &z)

	return float32(x), float32(y), float32(z), bool(ret)
}

func (v *Vehicle) SetAngularVelocity(x, y, z float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetAngularVelocity(v.ptr, C.float(x), C.float(y), C.float(z)))
}

func (v *Vehicle) SetParamsEx(engine, lights, alarm, doors, bonnet, boot, objective int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetParamsEx(
		v.ptr,
		C.int(engine),
		C.int(lights),
		C.int(alarm),
		C.int(doors),
		C.int(bonnet),
		C.int(boot),
		C.int(objective),
	))
}

func (v *Vehicle) GetParamsEx() (int, int, int, int, int, int, int, bool) {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0, 0, 0, 0, 0, 0, 0, false
	}

	var engine, lights, alarm, doors, bonnet, boot, objective C.int

	ret := C.Vehicle_GetParamsEx(
		v.ptr,
		&engine,
		&lights,
		&alarm,
		&doors,
		&bonnet,
		&boot,
		&objective,
	)

	return int(engine), int(lights), int(alarm), int(doors), int(bonnet), int(boot), int(objective), bool(ret)
}

func (v *Vehicle) SetParamsForPlayer(player *Player, objective, doors int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil || player == nil || player.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetParamsForPlayer(
		v.ptr,
		player.ptr,
		C.int(objective),
		C.int(doors),
	))
}

func (v *Vehicle) SetHealth(health float32) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetHealth(v.ptr, C.float(health)))
}

func (v *Vehicle) GetHealth() float32 {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0.0
	}

	return float32(C.Vehicle_GetHealth(v.ptr))
}

func (v *Vehicle) SetVirtualWorld(worldID int) bool {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return false
	}

	return bool(C.Vehicle_SetVirtualWorld(v.ptr, C.int(worldID)))
}

func (v *Vehicle) GetVirtualWorld() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetVirtualWorld(v.ptr))
}

func (v *Vehicle) GetInterior() int {
	apiMutex.Lock()
	defer apiMutex.Unlock()

	if v == nil || v.ptr == nil {
		return 0
	}

	return int(C.Vehicle_GetInterior(v.ptr))
}
