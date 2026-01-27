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
}

//export onPlayerConnect
func onPlayerConnect(args *C.struct_EventArgs_onPlayerConnect) C.bool {
	data := EventPlayerConnect{
		Player: PlayerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}

//export onPlayerDisconnect
func onPlayerDisconnect(args *C.struct_EventArgs_onPlayerDisconnect) C.bool {
	data := EventPlayerDisconnect{
		Player: PlayerFromPointer(*args.list.player),
	}

	emmitEvent(data)

	return C.bool(true)
}
