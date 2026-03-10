package colmap

/*
#include "lib/bullet_capi.h"
*/
import "C"
import "math"

const degToRad = float32(0.0174532925)

func EulerToQuat(rx, ry, rz float32) (qx, qy, qz, qw float32) {
	rx *= degToRad
	ry *= degToRad
	rz *= degToRad

	c1 := float32(math.Cos(float64(ry / 2)))
	s1 := float32(math.Sin(float64(ry / 2)))
	c2 := float32(math.Cos(float64(rx / 2)))
	s2 := float32(math.Sin(float64(rx / 2)))
	c3 := float32(math.Cos(float64(rz / 2)))
	s3 := float32(math.Sin(float64(rz / 2)))

	c1c2 := c1 * c2
	s1s2 := s1 * s2

	qw = c1c2*c3 - s1s2*s3
	qz = c1c2*s3 + s1s2*c3
	qy = s1*c2*c3 + c1*s2*s3
	qx = c1*s2*c3 - s1*c2*s3
	return
}

func QuatToEuler(qx, qy, qz, qw float32) (rx, ry, rz float32) {
	ry = float32(-math.Asin(float64(2*(qx*qz+qw*qy)))) * radToDeg
	rx = float32(math.Atan2(float64(2*(qy*qz+qw*qx)),
		float64(qw*qw-qx*qx-qy*qy+qz*qz))) * radToDeg
	rz = float32(-math.Atan2(float64(2*(qx*qy+qw*qz)),
		float64(qw*qw+qx*qx-qy*qy-qz*qz))) * radToDeg
	return
}

func GetModelBoundingSphere(modelID int) (cx, cy, cz, radius float32, ok bool) {
	shapeMu.RLock()
	s, exists := modelShapes[int32(modelID)]
	shapeMu.RUnlock()
	if !exists {
		return
	}
	var ccx, ccy, ccz, r C.float
	C.cm_shape_bounding_sphere(s, &ccx, &ccy, &ccz, &r)
	cx, cy, cz, radius = float32(ccx), float32(ccy), float32(ccz), float32(r)
	ok = true
	return
}

func GetModelBoundingBox(modelID int) (min, max [3]float32, ok bool) {
	shapeMu.RLock()
	s, exists := modelShapes[int32(modelID)]
	shapeMu.RUnlock()
	if !exists {
		return
	}
	var minx, miny, minz, maxx, maxy, maxz C.float
	C.cm_shape_bounding_box(s, &minx, &miny, &minz, &maxx, &maxy, &maxz)
	min = [3]float32{float32(minx), float32(miny), float32(minz)}
	max = [3]float32{float32(maxx), float32(maxy), float32(maxz)}
	ok = true
	return
}
