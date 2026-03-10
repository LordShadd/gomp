package colmap

/*
#cgo CXXFLAGS: -std=c++17 -I${SRCDIR}/bullet32/include/bullet -I${SRCDIR}/lib
#cgo LDFLAGS: -L${SRCDIR}/bullet32/lib -lBulletDynamics -lBulletCollision -lLinearMath -lstdc++
#include <stdlib.h>
#include "lib/bullet_capi.h"
*/
import "C"
import "sync"

var (
	world   C.CMWorld
	worldMu sync.Mutex
)
