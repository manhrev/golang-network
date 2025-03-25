package server

import (
	"golang-network/state"
	"golang-network/udp-server/service/clientmanager"
	"net"
	"time"
)

type Server struct {
	addr          string
	clientManager *clientmanager.Service
	state         *state.World
	conn          *net.UDPConn
}

func NewServer(addr string, state *state.World) *Server {
	return &Server{
		addr:  addr,
		state: state,
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
	s.clientManager = clientmanager.NewService(conn)

	go func() {
		for range time.Tick(3 * time.Second) {
			s.clientManager.StreamState(s.state)
		}
	}()

	return s.HandleRequest()
}
