package colmap

/*
#include "lib/bullet_capi.h"
*/
import "C"
import "math"

type RayCastResult struct {
	X, Y, Z float32
	Model   int
}

type RayCastResultEx struct {
	X, Y, Z        float32
	QX, QY, QZ, QW float32
	PX, PY, PZ     float32
	Model          int
}

type RayCastAngleResult struct {
	X, Y, Z    float32
	RX, RY, RZ float32
	Model      int
}

const radToDeg = float32(57.29577951)

func RayCastLine(sx, sy, sz, ex, ey, ez float32) (RayCastResult, bool) {
	var rx, ry, rz, nx, ny, nz C.float
	var modelId, bodyTag C.int
	worldMu.Lock()
	hit := C.cm_raytest(world, C.float(sx), C.float(sy), C.float(sz),
		C.float(ex), C.float(ey), C.float(ez),
		&rx, &ry, &rz, &nx, &ny, &nz, &modelId, &bodyTag)
	worldMu.Unlock()
	if hit == 0 {
		return RayCastResult{}, false
	}
	return RayCastResult{float32(rx), float32(ry), float32(rz), int(modelId)}, true
}

func RayCastLineID(sx, sy, sz, ex, ey, ez float32) (RayCastResult, int, bool) {
	var rx, ry, rz, nx, ny, nz C.float
	var modelId, bodyTag C.int
	worldMu.Lock()
	hit := C.cm_raytest(world, C.float(sx), C.float(sy), C.float(sz),
		C.float(ex), C.float(ey), C.float(ez),
		&rx, &ry, &rz, &nx, &ny, &nz, &modelId, &bodyTag)
	worldMu.Unlock()
	if hit == 0 {
		return RayCastResult{}, -1, false
	}
	return RayCastResult{float32(rx), float32(ry), float32(rz), int(modelId)}, int(bodyTag), true
}

func RayCastLineExtraID(extraType int, sx, sy, sz, ex, ey, ez float32) (RayCastResult, int32, bool) {
	var rx, ry, rz, nx, ny, nz C.float
	var modelId, bodyTag C.int
	worldMu.Lock()
	hit := C.cm_raytest(world, C.float(sx), C.float(sy), C.float(sz),
		C.float(ex), C.float(ey), C.float(ez),
		&rx, &ry, &rz, &nx, &ny, &nz, &modelId, &bodyTag)
	worldMu.Unlock()
	if hit == 0 {
		return RayCastResult{}, -1, false
	}
	data, _ := objManager.getExtraID(int(bodyTag), extraType)
	return RayCastResult{float32(rx), float32(ry), float32(rz), int(modelId)}, data, true
}

func RayCastLineEx(sx, sy, sz, ex, ey, ez float32) (RayCastResultEx, bool) {
	var rx, ry, rz, qx, qy, qz, qw, px, py, pz C.float
	var modelId C.int
	worldMu.Lock()
	hit := C.cm_raytest_ex(world, C.float(sx), C.float(sy), C.float(sz),
		C.float(ex), C.float(ey), C.float(ez),
		&rx, &ry, &rz, &qx, &qy, &qz, &qw, &px, &py, &pz, &modelId)
	worldMu.Unlock()
	if hit == 0 {
		return RayCastResultEx{}, false
	}
	return RayCastResultEx{
		float32(rx), float32(ry), float32(rz),
		float32(qx), float32(qy), float32(qz), float32(qw),
		float32(px), float32(py), float32(pz),
		int(modelId),
	}, true
}

func RayCastLineAngle(sx, sy, sz, ex, ey, ez float32) (RayCastAngleResult, bool) {
	var rx, ry, rz, nx, ny, nz C.float
	var modelId, bodyTag C.int
	worldMu.Lock()
	hit := C.cm_raytest(world, C.float(sx), C.float(sy), C.float(sz),
		C.float(ex), C.float(ey), C.float(ez),
		&rx, &ry, &rz, &nx, &ny, &nz, &modelId, &bodyTag)
	worldMu.Unlock()
	if hit == 0 {
		return RayCastAngleResult{}, false
	}
	rotX := -float32(math.Asin(float64(ny))) * radToDeg
	rotY := float32(math.Asin(float64(nx))) * radToDeg
	return RayCastAngleResult{
		float32(rx), float32(ry), float32(rz),
		rotX, rotY, 0,
		int(modelId),
	}, true
}

func RayCastLineAngleEx(sx, sy, sz, ex, ey, ez float32) (result RayCastAngleResult, objPos [3]float32, objEuler [3]float32, ok bool) {
	res, ok := RayCastLineEx(sx, sy, sz, ex, ey, ez)
	if !ok {
		return
	}
	nx := -float32(math.Asin(float64(res.QY))) * radToDeg
	ny := float32(math.Asin(float64(res.QX))) * radToDeg
	result = RayCastAngleResult{res.X, res.Y, res.Z, nx, ny, 0, res.Model}
	objPos = [3]float32{res.PX, res.PY, res.PZ}
	ex1, ey1, ez1 := QuatToEuler(res.QX, res.QY, res.QZ, res.QW)
	objEuler = [3]float32{ex1, ey1, ez1}
	return
}

func RayCastMultiLine(sx, sy, sz, ex, ey, ez float32, maxSize int) ([]RayCastResult, bool) {
	if maxSize <= 0 || maxSize > 99 {
		return nil, false
	}
	rxArr := make([]C.float, maxSize)
	ryArr := make([]C.float, maxSize)
	rzArr := make([]C.float, maxSize)
	distArr := make([]C.float, maxSize)
	modelArr := make([]C.int, maxSize)
	bodyTagArr := make([]C.int, maxSize)
	worldMu.Lock()
	count := C.cm_raytest_all(world,
		C.float(sx), C.float(sy), C.float(sz),
		C.float(ex), C.float(ey), C.float(ez),
		&rxArr[0], &ryArr[0], &rzArr[0], &distArr[0], &modelArr[0], &bodyTagArr[0],
		C.int(maxSize))
	worldMu.Unlock()
	if count <= 0 {
		return nil, false
	}
	out := make([]RayCastResult, int(count))
	for i := range out {
		out[i] = RayCastResult{float32(rxArr[i]), float32(ryArr[i]), float32(rzArr[i]), int(modelArr[i])}
	}
	return out, true
}

func RayCastLineNormal(sx, sy, sz, ex, ey, ez float32) (hit [3]float32, normal [3]float32, model int, ok bool) {
	var rx, ry, rz, nx, ny, nz C.float
	var modelId, bodyTag C.int
	worldMu.Lock()
	h := C.cm_raytest(world, C.float(sx), C.float(sy), C.float(sz),
		C.float(ex), C.float(ey), C.float(ez),
		&rx, &ry, &rz, &nx, &ny, &nz, &modelId, &bodyTag)
	worldMu.Unlock()
	if h == 0 {
		return
	}
	hit = [3]float32{float32(rx), float32(ry), float32(rz)}
	normal = [3]float32{float32(nx), float32(ny), float32(nz)}
	model = int(modelId)
	ok = true
	return
}

func RayCastReflectionVector(sx, sy, sz, ex, ey, ez float32) (hitPos [3]float32, reflection [3]float32, model int, ok bool) {
	var rx, ry, rz, nx, ny, nz C.float
	var modelId, bodyTag C.int
	worldMu.Lock()
	h := C.cm_raytest(world, C.float(sx), C.float(sy), C.float(sz),
		C.float(ex), C.float(ey), C.float(ez),
		&rx, &ry, &rz, &nx, &ny, &nz, &modelId, &bodyTag)
	worldMu.Unlock()
	if h == 0 {
		return
	}
	px, py, pz := float32(rx), float32(ry), float32(rz)
	nnx, nny, nnz := float32(nx), float32(ny), float32(nz)
	mag := dist3D(sx, sy, sz, px, py, pz)
	if mag == 0 {
		return
	}
	ux := (px - sx) / mag
	uy := (py - sy) / mag
	uz := (pz - sz) / mag
	dot := ux*nnx + uy*nny + uz*nnz
	hitPos = [3]float32{px, py, pz}
	reflection = [3]float32{ux - 2*dot*nnx, uy - 2*dot*nny, uz - 2*dot*nnz}
	model = int(modelId)
	ok = true
	return
}

func dist3D(ax, ay, az, bx, by, bz float32) float32 {
	dx := bx - ax
	dy := by - ay
	dz := bz - az
	return float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}
