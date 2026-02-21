package main

import (
	"fmt"
	"github.com/LordShadd/gomp"
	"strconv"
	"strings"
)

func init() {
	gomp.SetComponentName("BASIC GAMEMODE")
	gomp.SetComponentVersion(0, 0, 1, 0)

	gomp.OnReady(func() {
		gomp.EnableStuntBonus(false)
		gomp.UsePedAnims()
	})

	gomp.OnEvent(func(event gomp.EventPlayerRequestClass) {
		event.Player.Spawn()
	})

	gomp.OnEvent(func(event gomp.EventPlayerCommandText) {
		if strings.HasPrefix(event.Command, "/skin") {
			var skinID int
			var err error

			split := strings.Split(event.Command, " ")

			if len(split) > 1 {
				skinID, err = strconv.Atoi(split[1])

				if err != nil {
					skinID = 23
				}
			}

			event.Player.SetSkin(skinID)

			return
		}

		if strings.HasPrefix(event.Command, "/spawn") {
			event.Player.Spawn()

			return
		}

		if strings.HasPrefix(event.Command, "/veh") {
			var modelID int = 415
			var err error

			split := strings.Split(event.Command, " ")

			if len(split) > 1 {
				modelID, err = strconv.Atoi(split[1])

				if err != nil {
					modelID = 415
				}
			}

			x, y, z, _ := event.Player.GetPos()
			angle := event.Player.GetFacingAngle()

			oldVeh := event.Player.GetVehicle()
			oldVehAngle := oldVeh.GetZAngle()

			vx, vy, vz, _ := oldVeh.GetVelocity()

			if oldVeh != nil {
				angle = oldVehAngle
				oldVeh.Destroy()
			}

			vehicle := gomp.VehicleCreate(modelID, x, y, z, angle, 0, 0, -1, false)

			if vehicle == nil {
				return
			}

			event.Player.PutInVehicle(vehicle, 0)
			vehicle.SetVelocity(vx, vy, vz)
		}
	})

	gomp.OnEvent(func(event gomp.EventPlayerConnect) {
		name, _ := event.Player.GetName()

		fmt.Printf("Player %s joined.\n", name)
	})
}

func main() {}
