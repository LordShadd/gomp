package colmap

/*
#include "lib/bullet_capi.h"
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sync"
	"unsafe"
)

var (
	shapeMu      sync.RWMutex
	modelShapes  = make(map[int32]C.CMShape)
	convexShapes = make(map[int32]C.CMShape)

	modelPlacements  []itemPlacement
	removedBuildings []removedBuilding
)

type vec3 struct{ X, Y, Z float32 }
type quat struct{ X, Y, Z, W float32 }

type sphereData struct {
	Offset vec3
	Radius float32
}

type boxData struct {
	Center vec3
	Size   vec3
}

type faceData struct {
	A, B, C vec3
}

type collisionModel struct {
	ModelID uint16
	Spheres []sphereData
	Boxes   []boxData
	Faces   []faceData
}

type itemPlacement struct {
	ModelID  uint16
	Position vec3
	Rotation quat
}

type removedBuilding struct {
	Model           int16
	X, Y, Z, Radius float32
}

func readF32(r io.Reader) float32 {
	var v [4]byte
	r.Read(v[:])
	bits := binary.LittleEndian.Uint32(v[:])
	return *(*float32)(unsafe.Pointer(&bits))
}

func readU16(r io.Reader) uint16 {
	var v [2]byte
	r.Read(v[:])
	return binary.LittleEndian.Uint16(v[:])
}

func readU32(r io.Reader) uint32 {
	var v [4]byte
	r.Read(v[:])
	return binary.LittleEndian.Uint32(v[:])
}

func readVec3(r io.Reader) vec3 {
	return vec3{readF32(r), readF32(r), readF32(r)}
}

func readQuat(r io.Reader) quat {
	return quat{readF32(r), readF32(r), readF32(r), readF32(r)}
}

func loadCADB(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "colmap: cannot open %s: %v\n", path, err)
		return false
	}
	defer f.Close()

	sig := make([]byte, 4)
	f.Read(sig)
	if string(sig) != "cadf" {
		fmt.Fprintln(os.Stderr, "colmap: invalid .cadb signature")
		return false
	}

	version := readU16(f)
	if version != 2 {
		fmt.Fprintf(os.Stderr, "colmap: unsupported .cadb version %d\n", version)
		return false
	}

	modelCount := uint32(readU16(f))
	iplCount := readU32(f)

	models := make([]collisionModel, modelCount)
	for i := uint32(0); i < modelCount; i++ {
		m := &models[i]
		m.ModelID = readU16(f)
		sc := readU16(f)
		bc := readU16(f)
		fc := readU16(f)
		m.Spheres = make([]sphereData, sc)
		for j := uint16(0); j < sc; j++ {
			m.Spheres[j] = sphereData{readVec3(f), readF32(f)}
		}
		m.Boxes = make([]boxData, bc)
		for j := uint16(0); j < bc; j++ {
			m.Boxes[j] = boxData{readVec3(f), readVec3(f)}
		}
		m.Faces = make([]faceData, fc)
		for j := uint16(0); j < fc; j++ {
			m.Faces[j] = faceData{readVec3(f), readVec3(f), readVec3(f)}
		}
	}

	modelPlacements = make([]itemPlacement, iplCount)
	for i := uint32(0); i < iplCount; i++ {
		modelPlacements[i] = itemPlacement{
			ModelID:  readU16(f),
			Position: readVec3(f),
			Rotation: readQuat(f),
		}
	}

	buildShapes(models)
	return true
}

func buildShapes(models []collisionModel) {
	shapeMu.Lock()
	defer shapeMu.Unlock()
	for i := range models {
		m := &models[i]
		s := C.cm_shape_create()
		cs := C.cm_shape_create()

		for _, sp := range m.Spheres {
			C.cm_shape_add_sphere(s, C.float(sp.Offset.X), C.float(sp.Offset.Y), C.float(sp.Offset.Z), C.float(sp.Radius))
			C.cm_shape_add_sphere(cs, C.float(sp.Offset.X), C.float(sp.Offset.Y), C.float(sp.Offset.Z), C.float(sp.Radius))
		}
		for _, b := range m.Boxes {
			C.cm_shape_add_box(s, C.float(b.Center.X), C.float(b.Center.Y), C.float(b.Center.Z), C.float(b.Size.X), C.float(b.Size.Y), C.float(b.Size.Z))
			C.cm_shape_add_box(cs, C.float(b.Center.X), C.float(b.Center.Y), C.float(b.Center.Z), C.float(b.Size.X), C.float(b.Size.Y), C.float(b.Size.Z))
		}
		if len(m.Faces) > 0 {
			verts := make([]C.float, len(m.Faces)*9)
			indices := make([]C.int, len(m.Faces)*3)
			for j, f := range m.Faces {
				base := j * 9
				verts[base] = C.float(f.A.X)
				verts[base+1] = C.float(f.A.Y)
				verts[base+2] = C.float(f.A.Z)
				verts[base+3] = C.float(f.B.X)
				verts[base+4] = C.float(f.B.Y)
				verts[base+5] = C.float(f.B.Z)
				verts[base+6] = C.float(f.C.X)
				verts[base+7] = C.float(f.C.Y)
				verts[base+8] = C.float(f.C.Z)
				indices[j*3] = C.int(j * 3)
				indices[j*3+1] = C.int(j*3 + 1)
				indices[j*3+2] = C.int(j*3 + 2)
			}
			nv := C.int(len(m.Faces) * 3)
			nt := C.int(len(m.Faces))
			C.cm_shape_add_trimesh(s, &verts[0], nv, &indices[0], nt, 0)
			C.cm_shape_add_trimesh(cs, &verts[0], nv, &indices[0], nt, 1)
		}
		modelShapes[int32(m.ModelID)] = s
		convexShapes[int32(m.ModelID)] = cs
	}
}

func initMap() {
	objManager.reset()
	for i := range modelPlacements {
		p := &modelPlacements[i]
		if isRemoved(int16(p.ModelID), p.Position) {
			continue
		}
		shapeMu.RLock()
		s, ok := modelShapes[int32(p.ModelID)]
		shapeMu.RUnlock()
		if !ok {
			continue
		}
		C.cm_world_add_body(world, s,
			C.float(p.Position.X), C.float(p.Position.Y), C.float(p.Position.Z),
			C.float(p.Rotation.X), C.float(p.Rotation.Y), C.float(p.Rotation.Z), C.float(p.Rotation.W),
			C.int(p.ModelID), C.int(-1))
	}
	initWaterMesh()
}

func isRemoved(model int16, pos vec3) bool {
	for _, r := range removedBuildings {
		if r.Model != model && r.Model != -1 {
			continue
		}
		dx := pos.X - r.X
		dy := pos.Y - r.Y
		dz := pos.Z - r.Z
		if dx*dx+dy*dy+dz*dz <= r.Radius*r.Radius {
			return true
		}
	}
	return false
}

func ModelHasCollision(modelID int) bool {
	shapeMu.RLock()
	defer shapeMu.RUnlock()
	_, ok := modelShapes[int32(modelID)]
	return ok
}

func RemoveBuilding(model int, x, y, z, radius float32) {
	removedBuildings = append(removedBuildings, removedBuilding{
		Model: int16(model), X: x, Y: y, Z: z, Radius: radius,
	})
}

func RestoreBuilding(model int, x, y, z, radius float32) {
	worldMu.Lock()
	defer worldMu.Unlock()
	for i := range modelPlacements {
		p := &modelPlacements[i]
		if int16(p.ModelID) != int16(model) && model != -1 {
			continue
		}
		dx := p.Position.X - x
		dy := p.Position.Y - y
		dz := p.Position.Z - z
		if dx*dx+dy*dy+dz*dz > radius*radius {
			continue
		}
		shapeMu.RLock()
		s, ok := modelShapes[int32(p.ModelID)]
		shapeMu.RUnlock()
		if !ok {
			continue
		}
		C.cm_world_add_body(world, s,
			C.float(p.Position.X), C.float(p.Position.Y), C.float(p.Position.Z),
			C.float(p.Rotation.X), C.float(p.Rotation.Y), C.float(p.Rotation.Z), C.float(p.Rotation.W),
			C.int(p.ModelID), C.int(-1))
	}
}
