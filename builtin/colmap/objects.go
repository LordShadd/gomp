package colmap

/*
#include "lib/bullet_capi.h"
*/
import "C"
import "sync"

const maxMapObjects = 50000

type objectEntry struct {
	used       bool
	bodyHandle int
	extraData  [10]int32
}

type objectManagerT struct {
	mu      sync.RWMutex
	objects [maxMapObjects]objectEntry
}

var objManager objectManagerT

func (m *objectManagerT) reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.objects = [maxMapObjects]objectEntry{}
}

func (m *objectManagerT) add(bodyHandle int) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := 0; i < maxMapObjects; i++ {
		if !m.objects[i].used {
			m.objects[i] = objectEntry{used: true, bodyHandle: bodyHandle}
			return i
		}
	}
	return -1
}

func (m *objectManagerT) remove(idx int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if idx < 0 || idx >= maxMapObjects || !m.objects[idx].used {
		return false
	}
	m.objects[idx] = objectEntry{}
	return true
}

func (m *objectManagerT) valid(idx int) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return idx >= 0 && idx < maxMapObjects && m.objects[idx].used
}

func (m *objectManagerT) bodyHandle(idx int) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if idx < 0 || idx >= maxMapObjects || !m.objects[idx].used {
		return 0, false
	}
	return m.objects[idx].bodyHandle, true
}

func (m *objectManagerT) setExtraID(idx, typ int, data int32) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if idx < 0 || idx >= maxMapObjects || !m.objects[idx].used || typ < 0 || typ >= 10 {
		return false
	}
	m.objects[idx].extraData[typ] = data
	return true
}

func (m *objectManagerT) getExtraID(idx, typ int) (int32, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if idx < 0 || idx >= maxMapObjects || !m.objects[idx].used || typ < 0 || typ >= 10 {
		return -1, false
	}
	return m.objects[idx].extraData[typ], true
}

func CreateObject(modelID int, x, y, z, rx, ry, rz float32) (int, bool) {
	qx, qy, qz, qw := EulerToQuat(rx, ry, rz)
	shapeMu.RLock()
	s, ok := modelShapes[int32(modelID)]
	shapeMu.RUnlock()
	if !ok {
		return -1, false
	}
	worldMu.Lock()
	handle := int(C.cm_world_add_body(world, s,
		C.float(x), C.float(y), C.float(z),
		C.float(qx), C.float(qy), C.float(qz), C.float(qw),
		C.int(modelID), C.int(0)))
	worldMu.Unlock()
	if handle < 0 {
		return -1, false
	}
	idx := objManager.add(handle)
	if idx < 0 {
		worldMu.Lock()
		C.cm_world_remove_body(world, C.int(handle))
		worldMu.Unlock()
		return -1, false
	}
	return idx, true
}

func DestroyObject(idx int) bool {
	handle, ok := objManager.bodyHandle(idx)
	if !ok {
		return false
	}
	worldMu.Lock()
	C.cm_world_remove_body(world, C.int(handle))
	worldMu.Unlock()
	return objManager.remove(idx)
}

func IsValidObject(idx int) bool {
	return objManager.valid(idx)
}

func SetObjectPos(idx int, x, y, z float32) bool {
	handle, ok := objManager.bodyHandle(idx)
	if !ok {
		return false
	}
	worldMu.Lock()
	result := int(C.cm_world_set_body_pos(world, C.int(handle), C.float(x), C.float(y), C.float(z)))
	worldMu.Unlock()
	return result == 1
}

func SetObjectRot(idx int, rx, ry, rz float32) bool {
	handle, ok := objManager.bodyHandle(idx)
	if !ok {
		return false
	}
	qx, qy, qz, qw := EulerToQuat(rx, ry, rz)
	worldMu.Lock()
	result := int(C.cm_world_set_body_rot(world, C.int(handle), C.float(qx), C.float(qy), C.float(qz), C.float(qw)))
	worldMu.Unlock()
	return result == 1
}

func SetObjectExtraID(idx, typ int, data int32) bool {
	return objManager.setExtraID(idx, typ, data)
}

func GetObjectExtraID(idx, typ int) (int32, bool) {
	return objManager.getExtraID(idx, typ)
}

func ContactTest(modelID int, x, y, z, rx, ry, rz float32) bool {
	shapeMu.RLock()
	s, ok := convexShapes[int32(modelID)]
	shapeMu.RUnlock()
	if !ok {
		return false
	}
	qx, qy, qz, qw := EulerToQuat(rx, ry, rz)
	worldMu.Lock()
	result := int(C.cm_contact_test(world, s,
		C.float(x), C.float(y), C.float(z),
		C.float(qx), C.float(qy), C.float(qz), C.float(qw)))
	worldMu.Unlock()
	return result == 1
}
