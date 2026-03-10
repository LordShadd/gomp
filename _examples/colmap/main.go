package main

import (
	"fmt"

	"github.com/LordShadd/gomp"
	"github.com/LordShadd/gomp/colmap"
)

func init() {
	gomp.SetComponentName("COLMAP EXAMPLE")
	gomp.SetComponentVersion(0, 0, 1, 0)

	gomp.OnReady(func() {
		if !colmap.Init("scriptfiles/ColAndreas.cadb") {
			fmt.Println("colmap: failed to load ColAndreas.cadb")
		}
	})

	gomp.OnEvent(func(event gomp.EventPlayerRequestClass) {
		event.Player.Spawn()
	})

	gomp.OnEvent(func(event gomp.EventPlayerClickMap) {
		result, ok := colmap.RayCastLine(
			event.X, event.Y, 1000.0,
			event.X, event.Y, -1000.0,
		)

		if !ok {
			event.Player.SendClientMessage(0xFF4444FF, "No collision found at that position.")
			return
		}

		event.Player.SetPos(result.X, result.Y, result.Z+0.5)
		event.Player.SendClientMessage(0xFFFFFFFF, fmt.Sprintf(
			"Teleported to X: %.2f Y: %.2f Z: %.2f (model: %d)",
			result.X, result.Y, result.Z, result.Model,
		))
	})

	gomp.OnFree(func() {
		colmap.Destroy()
	})
}

func main() {}
