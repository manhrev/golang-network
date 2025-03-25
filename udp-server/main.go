package main

import (
	"golang-network/state"
	"golang-network/udp-server/server"
)

func main() {
	server := server.NewServer("localhost:1200", initState())
	server.Serve()

}

func initState() *state.World {
	world := state.NewWorld("manh's world")
	world.SetObject(state.NewObject("Manh", 0, 0))
	world.SetObject(state.NewObject("Van", 0, 1))

	return world
}
