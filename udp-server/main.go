package main

import (
	"golang-network/state"
)

func main() {
	server := NewServer("localhost:1200")
	server.Serve()

}

func initState() *state.World {
	world := state.NewWorld("manh's world")
	world.SetObject(state.NewObject("Manh", 0, 0))
	world.SetObject(state.NewObject("Van", 0, 1))

	return world
}
