package gomp

/*
#cgo linux CFLAGS: -I./lib -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOMP
#define GOMP

#include "main.h"
#include "event.h"

#endif
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type eventPriority int

const (
	eventPriorityHighest eventPriority = iota
	eventPriorityFairlyHigh
	eventPriorityDefault
	eventPriorityFairlyLow
	eventPriorityLowest
)

var eventHandlers map[reflect.Type][]reflect.Value = make(map[reflect.Type][]reflect.Value)

type EventPlayerConnect struct {
	Player *Player
}

type EventPlayerDisconnect struct {
	Player *Player
	Reason int
}

type EventPlayerCommandText struct {
	Player  *Player
	Command string
}

type EventPlayerRequestClass struct {
	Player  *Player
	ClassID int
}

type EventPlayerRequestSpawn struct {
	Player *Player
}

type EventPlayerSpawn struct {
	Player *Player
}

type EventPlayerDeath struct {
	Player *Player
	Killer *Player
	Reason int
}

type EventPlayerUpdate struct {
	Player *Player
}

type EventPlayerText struct {
	Player *Player
	Text   string
}

type EventPlayerInteriorChange struct {
	Player      *Player
	NewInterior int
	OldInterior int
}

type EventPlayerStateChange struct {
	Player   *Player
	NewState int
	OldState int
}

type EventPlayerKeyStateChange struct {
	Player  *Player
	NewKeys int
	OldKeys int
}

type EventPlayerEnterVehicle struct {
	Player    *Player
	Vehicle   *Vehicle
	Passenger bool
}

type EventPlayerExitVehicle struct {
	Player  *Player
	Vehicle *Vehicle
}

type EventPlayerEnterCheckpoint struct {
	Player *Player
}

type EventPlayerLeaveCheckpoint struct {
	Player *Player
}

type EventPlayerEnterRaceCheckpoint struct {
	Player *Player
}

type EventPlayerLeaveRaceCheckpoint struct {
	Player *Player
}

type EventPlayerGiveDamage struct {
	Player   *Player
	To       *Player
	Amount   float32
	Weapon   int
	BodyPart int
}

type EventPlayerTakeDamage struct {
	Player   *Player
	From     *Player
	Amount   float32
	Weapon   int
	BodyPart int
}

type EventPlayerGiveDamageActor struct {
	Player   *Player
	Actor    unsafe.Pointer
	Amount   float32
	Weapon   int
	BodyPart int
}

type EventPlayerShotMissed struct {
	Player *Player
	Weapon int
	X      float32
	Y      float32
	Z      float32
}

type EventPlayerShotPlayer struct {
	Player *Player
	Target *Player
	Weapon int
	X      float32
	Y      float32
	Z      float32
}

type EventPlayerShotVehicle struct {
	Player *Player
	Target *Vehicle
	Weapon int
	X      float32
	Y      float32
	Z      float32
}

type EventPlayerShotObject struct {
	Player *Player
	Target unsafe.Pointer
	Weapon int
	X      float32
	Y      float32
	Z      float32
}

type EventPlayerShotPlayerObject struct {
	Player *Player
	Target unsafe.Pointer
	Weapon int
	X      float32
	Y      float32
	Z      float32
}

type EventPlayerPickUpPickup struct {
	Player *Player
	Pickup unsafe.Pointer
}

type EventPlayerObjectMove struct {
	Player *Player
	Object unsafe.Pointer
}

type EventPlayerEditObject struct {
	Player    *Player
	Object    unsafe.Pointer
	Response  int
	OffsetX   float32
	OffsetY   float32
	OffsetZ   float32
	RotationX float32
	RotationY float32
	RotationZ float32
}

type EventPlayerEditAttachedObject struct {
	Player    *Player
	Saved     bool
	Index     int
	Model     int
	Bone      int
	OffsetX   float32
	OffsetY   float32
	OffsetZ   float32
	RotationX float32
	RotationY float32
	RotationZ float32
	ScaleX    float32
	ScaleY    float32
	ScaleZ    float32
}

type EventPlayerSelectObject struct {
	Player *Player
	Object unsafe.Pointer
	Model  int
	X      float32
	Y      float32
	Z      float32
}

type EventPlayerClickMap struct {
	Player *Player
	X      float32
	Y      float32
	Z      float32
}

type EventPlayerClickTextDraw struct {
	Player   *Player
	TextDraw unsafe.Pointer
}

type EventPlayerClickPlayerTextDraw struct {
	Player         *Player
	PlayerTextDraw unsafe.Pointer
}

type EventPlayerClickPlayer struct {
	Player  *Player
	Clicked *Player
	Source  int
}

type EventPlayerStreamIn struct {
	Player    *Player
	ForPlayer *Player
}

type EventPlayerStreamOut struct {
	Player    *Player
	ForPlayer *Player
}

type EventPlayerExitedMenu struct {
	Player *Player
}

type EventPlayerSelectedMenuRow struct {
	Player *Player
	Row    int
}

type EventPlayerRequestDownload struct {
	Player   *Player
	Type     int
	Checksum int
}

type EventTick struct{}

type EventIncomingConnection struct {
	Player    *Player
	IPAddress string
	Port      int
}

type EventRconLoginAttempt struct {
	Address  string
	Password string
	Success  bool
}

type EventConsoleText struct {
	Command    string
	Parameters string
}

type EventDialogResponse struct {
	Player    *Player
	DialogID  int
	Response  int
	ListItem  int
	InputText string
}

type EventVehicleSpawn struct {
	Vehicle *Vehicle
}

type EventVehicleDeath struct {
	Vehicle *Vehicle
	Player  *Player
}

type EventVehicleMod struct {
	Player    *Player
	Vehicle   *Vehicle
	Component int
}

type EventVehiclePaintJob struct {
	Player   *Player
	Vehicle  *Vehicle
	PaintJob int
}

type EventVehicleRespray struct {
	Player  *Player
	Vehicle *Vehicle
	Color1  int
	Color2  int
}

type EventVehicleDamageStatusUpdate struct {
	Vehicle *Vehicle
	Player  *Player
}

type EventVehicleSirenStateChange struct {
	Player     *Player
	Vehicle    *Vehicle
	SirenState int
}

type EventVehicleStreamIn struct {
	Vehicle *Vehicle
	Player  *Player
}

type EventVehicleStreamOut struct {
	Vehicle *Vehicle
	Player  *Player
}

type EventUnoccupiedVehicleUpdate struct {
	Vehicle   *Vehicle
	Player    *Player
	Seat      int
	PosX      float32
	PosY      float32
	PosZ      float32
	VelocityX float32
	VelocityY float32
	VelocityZ float32
}

type EventTrailerUpdate struct {
	Player  *Player
	Trailer *Vehicle
}

type EventActorStreamIn struct {
	Actor     unsafe.Pointer
	ForPlayer *Player
}

type EventActorStreamOut struct {
	Actor     unsafe.Pointer
	ForPlayer *Player
}

type EventObjectMove struct {
	Object unsafe.Pointer
}

type EventEnterExitModShop struct {
	Player     *Player
	EnterExit  int
	InteriorID int
}

type EventHandler func(data any)

func OnEvent(handler any) {
	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)

	if handlerType.Kind() != reflect.Func && handlerType.NumIn() != 1 {
		panic("Event handler should be a function with event data argument")
	}

	argType := handlerType.In(0)

	if eventHandlers[argType] == nil {
		eventHandlers[argType] = []reflect.Value{}
	}

	eventHandlers[argType] = append(eventHandlers[argType], handlerValue)
}

func addHandler(name string, priority eventPriority, fnPtr unsafe.Pointer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.Event_AddHandler(cName, C.int(priority), fnPtr)
}

func emmitEvent(data any) {
	handlers, ok := eventHandlers[reflect.TypeOf(data)]

	if !ok {
		return
	}

	args := []reflect.Value{reflect.ValueOf(data)}
	for _, handler := range handlers {
		handler.Call(args)
	}
}

//export entryPoint
func entryPoint() {
	addHandler("onPlayerConnect", eventPriorityDefault, C.onPlayerConnect)
	addHandler("onPlayerDisconnect", eventPriorityDefault, C.onPlayerDisconnect)
	addHandler("onPlayerCommandText", eventPriorityDefault, C.onPlayerCommandText)
	addHandler("onPlayerRequestClass", eventPriorityDefault, C.onPlayerRequestClass)
	addHandler("onPlayerRequestSpawn", eventPriorityDefault, C.onPlayerRequestSpawn)
	addHandler("onPlayerSpawn", eventPriorityDefault, C.onPlayerSpawn)
	addHandler("onPlayerDeath", eventPriorityDefault, C.onPlayerDeath)
	addHandler("onPlayerUpdate", eventPriorityDefault, C.onPlayerUpdate)
	addHandler("onPlayerText", eventPriorityDefault, C.onPlayerText)
	addHandler("onPlayerInteriorChange", eventPriorityDefault, C.onPlayerInteriorChange)
	addHandler("onPlayerStateChange", eventPriorityDefault, C.onPlayerStateChange)
	addHandler("onPlayerKeyStateChange", eventPriorityDefault, C.onPlayerKeyStateChange)
	addHandler("onPlayerEnterVehicle", eventPriorityDefault, C.onPlayerEnterVehicle)
	addHandler("onPlayerExitVehicle", eventPriorityDefault, C.onPlayerExitVehicle)
	addHandler("onPlayerEnterCheckpoint", eventPriorityDefault, C.onPlayerEnterCheckpoint)
	addHandler("onPlayerLeaveCheckpoint", eventPriorityDefault, C.onPlayerLeaveCheckpoint)
	addHandler("onPlayerEnterRaceCheckpoint", eventPriorityDefault, C.onPlayerEnterRaceCheckpoint)
	addHandler("onPlayerLeaveRaceCheckpoint", eventPriorityDefault, C.onPlayerLeaveRaceCheckpoint)
	addHandler("onPlayerGiveDamage", eventPriorityDefault, C.onPlayerGiveDamage)
	addHandler("onPlayerTakeDamage", eventPriorityDefault, C.onPlayerTakeDamage)
	addHandler("onPlayerGiveDamageActor", eventPriorityDefault, C.onPlayerGiveDamageActor)
	addHandler("onPlayerShotMissed", eventPriorityDefault, C.onPlayerShotMissed)
	addHandler("onPlayerShotPlayer", eventPriorityDefault, C.onPlayerShotPlayer)
	addHandler("onPlayerShotVehicle", eventPriorityDefault, C.onPlayerShotVehicle)
	addHandler("onPlayerShotObject", eventPriorityDefault, C.onPlayerShotObject)
	addHandler("onPlayerShotPlayerObject", eventPriorityDefault, C.onPlayerShotPlayerObject)
	addHandler("onPlayerPickUpPickup", eventPriorityDefault, C.onPlayerPickUpPickup)
	addHandler("onPlayerObjectMove", eventPriorityDefault, C.onPlayerObjectMove)
	addHandler("onPlayerEditObject", eventPriorityDefault, C.onPlayerEditObject)
	addHandler("onPlayerEditAttachedObject", eventPriorityDefault, C.onPlayerEditAttachedObject)
	addHandler("onPlayerSelectObject", eventPriorityDefault, C.onPlayerSelectObject)
	addHandler("onPlayerClickMap", eventPriorityDefault, C.onPlayerClickMap)
	addHandler("onPlayerClickTextDraw", eventPriorityDefault, C.onPlayerClickTextDraw)
	addHandler("onPlayerClickPlayerTextDraw", eventPriorityDefault, C.onPlayerClickPlayerTextDraw)
	addHandler("onPlayerClickPlayer", eventPriorityDefault, C.onPlayerClickPlayer)
	addHandler("onPlayerStreamIn", eventPriorityDefault, C.onPlayerStreamIn)
	addHandler("onPlayerStreamOut", eventPriorityDefault, C.onPlayerStreamOut)
	addHandler("onPlayerExitedMenu", eventPriorityDefault, C.onPlayerExitedMenu)
	addHandler("onPlayerSelectedMenuRow", eventPriorityDefault, C.onPlayerSelectedMenuRow)
	addHandler("onPlayerRequestDownload", eventPriorityDefault, C.onPlayerRequestDownload)
	addHandler("onTick", eventPriorityDefault, C.onTick)
	addHandler("onIncomingConnection", eventPriorityDefault, C.onIncomingConnection)
	addHandler("onRconLoginAttempt", eventPriorityDefault, C.onRconLoginAttempt)
	addHandler("onConsoleText", eventPriorityDefault, C.onConsoleText)
	addHandler("onDialogResponse", eventPriorityDefault, C.onDialogResponse)
	addHandler("onVehicleSpawn", eventPriorityDefault, C.onVehicleSpawn)
	addHandler("onVehicleDeath", eventPriorityDefault, C.onVehicleDeath)
	addHandler("onVehicleMod", eventPriorityDefault, C.onVehicleMod)
	addHandler("onVehiclePaintJob", eventPriorityDefault, C.onVehiclePaintJob)
	addHandler("onVehicleRespray", eventPriorityDefault, C.onVehicleRespray)
	addHandler("onVehicleDamageStatusUpdate", eventPriorityDefault, C.onVehicleDamageStatusUpdate)
	addHandler("onVehicleSirenStateChange", eventPriorityDefault, C.onVehicleSirenStateChange)
	addHandler("onVehicleStreamIn", eventPriorityDefault, C.onVehicleStreamIn)
	addHandler("onVehicleStreamOut", eventPriorityDefault, C.onVehicleStreamOut)
	addHandler("onUnoccupiedVehicleUpdate", eventPriorityDefault, C.onUnoccupiedVehicleUpdate)
	addHandler("onTrailerUpdate", eventPriorityDefault, C.onTrailerUpdate)
	addHandler("onActorStreamIn", eventPriorityDefault, C.onActorStreamIn)
	addHandler("onActorStreamOut", eventPriorityDefault, C.onActorStreamOut)
	addHandler("onObjectMove", eventPriorityDefault, C.onObjectMove)
	addHandler("onEnterExitModShop", eventPriorityDefault, C.onEnterExitModShop)
}

//export onPlayerConnect
func onPlayerConnect(args *C.struct_EventArgs_onPlayerConnect) C.bool {
	data := EventPlayerConnect{
		Player: playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerDisconnect
func onPlayerDisconnect(args *C.struct_EventArgs_onPlayerDisconnect) C.bool {
	data := EventPlayerDisconnect{
		Player: playerFromPointer(*args.list.player),
		Reason: int(*args.list.reason),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerCommandText
func onPlayerCommandText(args *C.struct_EventArgs_onPlayerCommandText) C.bool {
	data := EventPlayerCommandText{
		Player:  playerFromPointer(*args.list.player),
		Command: C.GoStringN(args.list.command.data, C.int(args.list.command.len)),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerRequestClass
func onPlayerRequestClass(args *C.struct_EventArgs_onPlayerRequestClass) C.bool {
	data := EventPlayerRequestClass{
		Player:  playerFromPointer(*args.list.player),
		ClassID: int(*args.list.classId),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerRequestSpawn
func onPlayerRequestSpawn(args *C.struct_EventArgs_onPlayerRequestSpawn) C.bool {
	data := EventPlayerRequestSpawn{
		Player: playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerSpawn
func onPlayerSpawn(args *C.struct_EventArgs_onPlayerSpawn) C.bool {
	data := EventPlayerSpawn{
		Player: playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerDeath
func onPlayerDeath(args *C.struct_EventArgs_onPlayerDeath) C.bool {
	data := EventPlayerDeath{
		Player: playerFromPointer(*args.list.player),
		Killer: playerFromPointer(*args.list.killer),
		Reason: int(*args.list.reason),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerUpdate
func onPlayerUpdate(args *C.struct_EventArgs_onPlayerUpdate) C.bool {
	data := EventPlayerUpdate{
		Player: playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerText
func onPlayerText(args *C.struct_EventArgs_onPlayerText) C.bool {
	data := EventPlayerText{
		Player: playerFromPointer(*args.list.player),
		Text:   C.GoStringN(args.list.text.data, C.int(args.list.text.len)),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerInteriorChange
func onPlayerInteriorChange(args *C.struct_EventArgs_onPlayerInteriorChange) C.bool {
	data := EventPlayerInteriorChange{
		Player:      playerFromPointer(*args.list.player),
		NewInterior: int(*args.list.newInterior),
		OldInterior: int(*args.list.oldInterior),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerStateChange
func onPlayerStateChange(args *C.struct_EventArgs_onPlayerStateChange) C.bool {
	data := EventPlayerStateChange{
		Player:   playerFromPointer(*args.list.player),
		NewState: int(*args.list.newState),
		OldState: int(*args.list.oldState),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(args *C.struct_EventArgs_onPlayerKeyStateChange) C.bool {
	data := EventPlayerKeyStateChange{
		Player:  playerFromPointer(*args.list.player),
		NewKeys: int(*args.list.newKeys),
		OldKeys: int(*args.list.oldKeys),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(args *C.struct_EventArgs_onPlayerEnterVehicle) C.bool {
	data := EventPlayerEnterVehicle{
		Player:    playerFromPointer(*args.list.player),
		Vehicle:   &Vehicle{ptr: *args.list.vehicle},
		Passenger: bool(*args.list.passenger),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(args *C.struct_EventArgs_onPlayerExitVehicle) C.bool {
	data := EventPlayerExitVehicle{
		Player:  playerFromPointer(*args.list.player),
		Vehicle: &Vehicle{ptr: *args.list.vehicle},
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerEnterCheckpoint
func onPlayerEnterCheckpoint(args *C.struct_EventArgs_onPlayerEnterCheckpoint) C.bool {
	data := EventPlayerEnterCheckpoint{
		Player: playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerLeaveCheckpoint
func onPlayerLeaveCheckpoint(args *C.struct_EventArgs_onPlayerLeaveCheckpoint) C.bool {
	data := EventPlayerLeaveCheckpoint{
		Player: playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerEnterRaceCheckpoint
func onPlayerEnterRaceCheckpoint(args *C.struct_EventArgs_onPlayerEnterRaceCheckpoint) C.bool {
	data := EventPlayerEnterRaceCheckpoint{
		Player: playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerLeaveRaceCheckpoint
func onPlayerLeaveRaceCheckpoint(args *C.struct_EventArgs_onPlayerLeaveRaceCheckpoint) C.bool {
	data := EventPlayerLeaveRaceCheckpoint{
		Player: playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(args *C.struct_EventArgs_onPlayerGiveDamage) C.bool {
	data := EventPlayerGiveDamage{
		Player:   playerFromPointer(*args.list.player),
		To:       playerFromPointer(*args.list.to),
		Amount:   float32(*args.list.amount),
		Weapon:   int(*args.list.weapon),
		BodyPart: int(*args.list.bodypart),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(args *C.struct_EventArgs_onPlayerTakeDamage) C.bool {
	data := EventPlayerTakeDamage{
		Player:   playerFromPointer(*args.list.player),
		From:     playerFromPointer(*args.list.from),
		Amount:   float32(*args.list.amount),
		Weapon:   int(*args.list.weapon),
		BodyPart: int(*args.list.bodypart),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerGiveDamageActor
func onPlayerGiveDamageActor(args *C.struct_EventArgs_onPlayerGiveDamageActor) C.bool {
	data := EventPlayerGiveDamageActor{
		Player:   playerFromPointer(*args.list.player),
		Actor:    *args.list.actor,
		Amount:   float32(*args.list.amount),
		Weapon:   int(*args.list.weapon),
		BodyPart: int(*args.list.part),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerShotMissed
func onPlayerShotMissed(args *C.struct_EventArgs_onPlayerShotMissed) C.bool {
	data := EventPlayerShotMissed{
		Player: playerFromPointer(*args.list.player),
		Weapon: int(*args.list.weapon),
		X:      float32(*args.list.x),
		Y:      float32(*args.list.y),
		Z:      float32(*args.list.z),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerShotPlayer
func onPlayerShotPlayer(args *C.struct_EventArgs_onPlayerShotPlayer) C.bool {
	data := EventPlayerShotPlayer{
		Player: playerFromPointer(*args.list.player),
		Target: playerFromPointer(*args.list.target),
		Weapon: int(*args.list.weapon),
		X:      float32(*args.list.x),
		Y:      float32(*args.list.y),
		Z:      float32(*args.list.z),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerShotVehicle
func onPlayerShotVehicle(args *C.struct_EventArgs_onPlayerShotVehicle) C.bool {
	data := EventPlayerShotVehicle{
		Player: playerFromPointer(*args.list.player),
		Target: &Vehicle{ptr: *args.list.target},
		Weapon: int(*args.list.weapon),
		X:      float32(*args.list.x),
		Y:      float32(*args.list.y),
		Z:      float32(*args.list.z),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerShotObject
func onPlayerShotObject(args *C.struct_EventArgs_onPlayerShotObject) C.bool {
	data := EventPlayerShotObject{
		Player: playerFromPointer(*args.list.player),
		Target: *args.list.target,
		Weapon: int(*args.list.weapon),
		X:      float32(*args.list.x),
		Y:      float32(*args.list.y),
		Z:      float32(*args.list.z),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerShotPlayerObject
func onPlayerShotPlayerObject(args *C.struct_EventArgs_onPlayerShotPlayerObject) C.bool {
	data := EventPlayerShotPlayerObject{
		Player: playerFromPointer(*args.list.player),
		Target: *args.list.target,
		Weapon: int(*args.list.weapon),
		X:      float32(*args.list.x),
		Y:      float32(*args.list.y),
		Z:      float32(*args.list.z),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerPickUpPickup
func onPlayerPickUpPickup(args *C.struct_EventArgs_onPlayerPickUpPickup) C.bool {
	data := EventPlayerPickUpPickup{
		Player: playerFromPointer(*args.list.player),
		Pickup: *args.list.pickup,
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerObjectMove
func onPlayerObjectMove(args *C.struct_EventArgs_onPlayerObjectMove) C.bool {
	data := EventPlayerObjectMove{
		Player: playerFromPointer(*args.list.player),
		Object: *args.list.object,
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerEditObject
func onPlayerEditObject(args *C.struct_EventArgs_onPlayerEditObject) C.bool {
	data := EventPlayerEditObject{
		Player:    playerFromPointer(*args.list.player),
		Object:    *args.list.object,
		Response:  int(*args.list.response),
		OffsetX:   float32(*args.list.offsetX),
		OffsetY:   float32(*args.list.offsetY),
		OffsetZ:   float32(*args.list.offsetZ),
		RotationX: float32(*args.list.rotationX),
		RotationY: float32(*args.list.rotationY),
		RotationZ: float32(*args.list.rotationZ),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerEditAttachedObject
func onPlayerEditAttachedObject(args *C.struct_EventArgs_onPlayerEditAttachedObject) C.bool {
	data := EventPlayerEditAttachedObject{
		Player:    playerFromPointer(*args.list.player),
		Saved:     bool(*args.list.saved),
		Index:     int(*args.list.index),
		Model:     int(*args.list.model),
		Bone:      int(*args.list.bone),
		OffsetX:   float32(*args.list.offsetX),
		OffsetY:   float32(*args.list.offsetY),
		OffsetZ:   float32(*args.list.offsetZ),
		RotationX: float32(*args.list.rotationX),
		RotationY: float32(*args.list.rotationY),
		RotationZ: float32(*args.list.rotationZ),
		ScaleX:    float32(*args.list.scaleX),
		ScaleY:    float32(*args.list.scaleY),
		ScaleZ:    float32(*args.list.scaleZ),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerSelectObject
func onPlayerSelectObject(args *C.struct_EventArgs_onPlayerSelectObject) C.bool {
	data := EventPlayerSelectObject{
		Player: playerFromPointer(*args.list.player),
		Object: *args.list.object,
		Model:  int(*args.list.model),
		X:      float32(*args.list.x),
		Y:      float32(*args.list.y),
		Z:      float32(*args.list.z),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerClickMap
func onPlayerClickMap(args *C.struct_EventArgs_onPlayerClickMap) C.bool {
	data := EventPlayerClickMap{
		Player: playerFromPointer(*args.list.player),
		X:      float32(*args.list.x),
		Y:      float32(*args.list.y),
		Z:      float32(*args.list.z),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerClickTextDraw
func onPlayerClickTextDraw(args *C.struct_EventArgs_onPlayerClickTextDraw) C.bool {
	data := EventPlayerClickTextDraw{
		Player:   playerFromPointer(*args.list.player),
		TextDraw: *args.list.textdraw,
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerClickPlayerTextDraw
func onPlayerClickPlayerTextDraw(args *C.struct_EventArgs_onPlayerClickPlayerTextDraw) C.bool {
	data := EventPlayerClickPlayerTextDraw{
		Player:         playerFromPointer(*args.list.player),
		PlayerTextDraw: *args.list.textdraw,
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(args *C.struct_EventArgs_onPlayerClickPlayer) C.bool {
	data := EventPlayerClickPlayer{
		Player:  playerFromPointer(*args.list.player),
		Clicked: playerFromPointer(*args.list.clicked),
		Source:  int(*args.list.source),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerStreamIn
func onPlayerStreamIn(args *C.struct_EventArgs_onPlayerStreamIn) C.bool {
	data := EventPlayerStreamIn{
		Player:    playerFromPointer(*args.list.player),
		ForPlayer: playerFromPointer(*args.list.forPlayer),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerStreamOut
func onPlayerStreamOut(args *C.struct_EventArgs_onPlayerStreamOut) C.bool {
	data := EventPlayerStreamOut{
		Player:    playerFromPointer(*args.list.player),
		ForPlayer: playerFromPointer(*args.list.forPlayer),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerExitedMenu
func onPlayerExitedMenu(args *C.struct_EventArgs_onPlayerExitedMenu) C.bool {
	data := EventPlayerExitedMenu{
		Player: playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerSelectedMenuRow
func onPlayerSelectedMenuRow(args *C.struct_EventArgs_onPlayerSelectedMenuRow) C.bool {
	data := EventPlayerSelectedMenuRow{
		Player: playerFromPointer(*args.list.player),
		Row:    int(*args.list.row),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerRequestDownload
func onPlayerRequestDownload(args *C.struct_EventArgs_onPlayerRequestDownload) C.bool {
	data := EventPlayerRequestDownload{
		Player:   playerFromPointer(*args.list.player),
		Type:     int(*args.list._type),
		Checksum: int(*args.list.checksum),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onTick
func onTick() {
	data := EventTick{}

	emmitEvent(data)
}

//export onIncomingConnection
func onIncomingConnection(args *C.struct_EventArgs_onIncomingConnection) C.bool {
	data := EventIncomingConnection{
		Player:    playerFromPointer(*args.list.player),
		IPAddress: C.GoStringN(args.list.ipAddress.data, C.int(args.list.ipAddress.len)),
		Port:      int(*args.list.port),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onRconLoginAttempt
func onRconLoginAttempt(args *C.struct_EventArgs_onRconLoginAttempt) C.bool {
	data := EventRconLoginAttempt{
		Address:  C.GoStringN(args.list.address.data, C.int(args.list.address.len)),
		Password: C.GoStringN(args.list.password.data, C.int(args.list.password.len)),
		Success:  bool(*args.list.success),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onConsoleText
func onConsoleText(args *C.struct_EventArgs_onConsoleText) C.bool {
	data := EventConsoleText{
		Command:    C.GoStringN(args.list.command.data, C.int(args.list.command.len)),
		Parameters: C.GoStringN(args.list.parameters.data, C.int(args.list.parameters.len)),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onDialogResponse
func onDialogResponse(args *C.struct_EventArgs_onDialogResponse) C.bool {
	data := EventDialogResponse{
		Player:    playerFromPointer(*args.list.player),
		DialogID:  int(*args.list.dialogId),
		Response:  int(*args.list.response),
		ListItem:  int(*args.list.listItem),
		InputText: C.GoStringN(args.list.inputText.data, C.int(args.list.inputText.len)),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onVehicleSpawn
func onVehicleSpawn(args *C.struct_EventArgs_onVehicleSpawn) C.bool {
	data := EventVehicleSpawn{
		Vehicle: &Vehicle{ptr: *args.list.vehicle},
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onVehicleDeath
func onVehicleDeath(args *C.struct_EventArgs_onVehicleDeath) C.bool {
	data := EventVehicleDeath{
		Vehicle: &Vehicle{ptr: *args.list.vehicle},
		Player:  playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onVehicleMod
func onVehicleMod(args *C.struct_EventArgs_onVehicleMod) C.bool {
	data := EventVehicleMod{
		Player:    playerFromPointer(*args.list.player),
		Vehicle:   &Vehicle{ptr: *args.list.vehicle},
		Component: int(*args.list.component),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onVehiclePaintJob
func onVehiclePaintJob(args *C.struct_EventArgs_onVehiclePaintJob) C.bool {
	data := EventVehiclePaintJob{
		Player:   playerFromPointer(*args.list.player),
		Vehicle:  &Vehicle{ptr: *args.list.vehicle},
		PaintJob: int(*args.list.paintJob),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onVehicleRespray
func onVehicleRespray(args *C.struct_EventArgs_onVehicleRespray) C.bool {
	data := EventVehicleRespray{
		Player:  playerFromPointer(*args.list.player),
		Vehicle: &Vehicle{ptr: *args.list.vehicle},
		Color1:  int(*args.list.color1),
		Color2:  int(*args.list.color2),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(args *C.struct_EventArgs_onVehicleDamageStatusUpdate) C.bool {
	data := EventVehicleDamageStatusUpdate{
		Vehicle: &Vehicle{ptr: *args.list.vehicle},
		Player:  playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(args *C.struct_EventArgs_onVehicleSirenStateChange) C.bool {
	data := EventVehicleSirenStateChange{
		Player:     playerFromPointer(*args.list.player),
		Vehicle:    &Vehicle{ptr: *args.list.vehicle},
		SirenState: int(*args.list.sirenState),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onVehicleStreamIn
func onVehicleStreamIn(args *C.struct_EventArgs_onVehicleStreamIn) C.bool {
	data := EventVehicleStreamIn{
		Vehicle: &Vehicle{ptr: *args.list.vehicle},
		Player:  playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onVehicleStreamOut
func onVehicleStreamOut(args *C.struct_EventArgs_onVehicleStreamOut) C.bool {
	data := EventVehicleStreamOut{
		Vehicle: &Vehicle{ptr: *args.list.vehicle},
		Player:  playerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(args *C.struct_EventArgs_onUnoccupiedVehicleUpdate) C.bool {
	data := EventUnoccupiedVehicleUpdate{
		Vehicle:   &Vehicle{ptr: *args.list.vehicle},
		Player:    playerFromPointer(*args.list.player),
		Seat:      int(*args.list.seat),
		PosX:      float32(*args.list.posX),
		PosY:      float32(*args.list.posY),
		PosZ:      float32(*args.list.posZ),
		VelocityX: float32(*args.list.velocityX),
		VelocityY: float32(*args.list.velocityY),
		VelocityZ: float32(*args.list.velocityZ),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onTrailerUpdate
func onTrailerUpdate(args *C.struct_EventArgs_onTrailerUpdate) C.bool {
	data := EventTrailerUpdate{
		Player:  playerFromPointer(*args.list.player),
		Trailer: &Vehicle{ptr: *args.list.trailer},
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onActorStreamIn
func onActorStreamIn(args *C.struct_EventArgs_onActorStreamIn) C.bool {
	data := EventActorStreamIn{
		Actor:     *args.list.actor,
		ForPlayer: playerFromPointer(*args.list.forPlayer),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onActorStreamOut
func onActorStreamOut(args *C.struct_EventArgs_onActorStreamOut) C.bool {
	data := EventActorStreamOut{
		Actor:     *args.list.actor,
		ForPlayer: playerFromPointer(*args.list.forPlayer),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onObjectMove
func onObjectMove(args *C.struct_EventArgs_onObjectMove) C.bool {
	data := EventObjectMove{
		Object: *args.list.object,
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onEnterExitModShop
func onEnterExitModShop(args *C.struct_EventArgs_onEnterExitModShop) C.bool {
	data := EventEnterExitModShop{
		Player:     playerFromPointer(*args.list.player),
		EnterExit:  int(*args.list.enterexit),
		InteriorID: int(*args.list.interiorId),
	}

	emmitEvent(data)

	return C.bool(true)
}
