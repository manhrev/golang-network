package main

import (
	"encoding/json"
	"golang-network/state"
	"net"
)

type Client struct {
	conn *net.UDPConn
	name string
	addr *net.UDPAddr
	ch   chan *state.World
}

func NewClient(name string, conn *net.UDPConn, addr *net.UDPAddr, ch chan *state.World) *Client {
	return &Client{
		name: name,
		conn: conn,
		addr: addr,
		ch:   ch,
	}
}

func (c Client) HandleRequest() {
	for {
		state := <-c.ch
		println("sending to: ", c.name)
		stateSerialized, err := json.Marshal(state)
		if err != nil {
			panic(err)
		}

		_, err = c.conn.WriteToUDP(stateSerialized, c.addr)
		if err != nil {
			panic(err)
		}
	}
}
