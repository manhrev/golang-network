package clientmanager

import (
	"encoding/json"
	"fmt"
	"golang-network/state"
	"net"
)

type Client struct {
	id     int64
	conn   *net.UDPConn
	addr   *net.UDPAddr
	ch     chan *state.World
	stopch chan bool
}

func NewClient(id int64, conn *net.UDPConn, addr *net.UDPAddr, ch chan *state.World) *Client {
	return &Client{
		id:     id,
		conn:   conn,
		addr:   addr,
		ch:     ch,
		stopch: make(chan bool),
	}
}

func (c Client) ListenWrite() {
	_, err := c.conn.WriteToUDP([]byte(fmt.Sprintf("your id: %d", c.id)), c.addr)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case state := <-c.ch:
			println("sending to: ", c.id)
			stateSerialized, err := json.Marshal(state)
			if err != nil {
				panic(err)
			}

			_, err = c.conn.WriteToUDP(stateSerialized, c.addr)
			if err != nil {
				panic(err)
			}
		case <-c.stopch:
			return
		}
	}
}

func (c Client) Stop() {
	c.stopch <- true
}
