package main

import (
	"fmt"
	"gomp"
)

func init() {
	gomp.SetComponentName("BASIC GAMEMODE")
	gomp.SetComponentVersion(0, 0, 1, 0)

	gomp.OnEvent(func(data gomp.EventPlayerConnect) {
		playerName, ok := data.Player.GetName()

		if !ok {
			fmt.Println("It was not possible to get player name")
			return
		}

		fmt.Printf("Player %s joined.\n", playerName)
	})
}

func main() {}
