package main

import (
	"golang-network/state"
	"net"
	"strings"
	"time"
)

type Server struct {
	addr    string
	clients map[string]*Client
	state   *state.World
	conn    *net.UDPConn
}

func NewServer(addr string) *Server {
	state := initState()
	return &Server{
		addr:    addr,
		state:   state,
		clients: make(map[string]*Client),
	}
}

func (s *Server) AddClientIfNotExist(c *Client) {
	_, ok := s.clients[c.name]
	if ok {
		return
	}

	s.clients[c.name] = c
}

func (s *Server) StreamState() {
	for {
		time.Sleep(2 * time.Second)
		for _, c := range s.clients {
			c.ch <- s.state
		}
	}
}

func (s *Server) Serve() error {
	udpAddr, err := net.ResolveUDPAddr("udp", s.addr)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	s.conn = conn

	go s.StreamState()

	for {
		var buf [512]byte
		_, caddr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			// fmt.Println(err)
			return err
		}

		println(caddr.String())

		connectString := string(buf[:])
		ss := strings.Split(connectString, " ")
		if len(ss) != 2 {
			println("Not a connect request 1")
			continue
		}

		if ss[0] != "connect" {
			println("Not a connect request 2")
			continue
		}

		name := ss[1]
		ch := make(chan *state.World)
		client := NewClient(name, conn, caddr, ch)
		s.AddClientIfNotExist(client)

		go client.HandleRequest()
	}
}
