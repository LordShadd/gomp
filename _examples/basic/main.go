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

	gomp.OnEvent(func(data gomp.EventPlayerRequestClass) {
		data.Player.Spawn()
	})

	gomp.OnEvent(func(data gomp.EventPlayerCommandText) {
		if strings.HasPrefix(data.Command, "/skin") {
			var skinID int
			var err error

			split := strings.Split(data.Command, " ")

			if len(split) > 1 {
				skinID, err = strconv.Atoi(split[1])

				if err != nil {
					skinID = 23
				}
			}

			data.Player.SetSkin(skinID)

			return
		}

		if strings.HasPrefix(data.Command, "/spawn") {
			data.Player.Spawn()

			return
		}

		if strings.HasPrefix(data.Command, "/veh") {
			var modelID int = 415
			var err error

			split := strings.Split(data.Command, " ")

			if len(split) > 1 {
				modelID, err = strconv.Atoi(split[1])

				if err != nil {
					modelID = 415
				}
			}

			x, y, z, _ := data.Player.GetPos()
			angle := data.Player.GetFacingAngle()

			oldVeh := data.Player.GetVehicle()
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

			data.Player.PutInVehicle(vehicle, 0)
			vehicle.SetVelocity(vx, vy, vz)
		}
	})

	gomp.OnEvent(func(data gomp.EventPlayerConnect) {
		name, _ := data.Player.GetName()

		fmt.Printf("Player %s joined.\n", name)
	})
}

func main() {}
