package clientmanager

import (
	"golang-network/state"
	"net"
)

type Service struct {
	maxID   int64
	clients map[int64]*Client
	conn    *net.UDPConn
	//mutexlock
}

func NewService(conn *net.UDPConn) *Service {
	return &Service{
		maxID:   1,
		clients: make(map[int64]*Client),
		conn:    conn,
	}
}

func (s *Service) StreamState(w *state.World) {
	for _, c := range s.clients {
		c.ch <- w
	}
}

func (s *Service) AddNewClient(addr *net.UDPAddr) {
	client := NewClient(s.maxID, s.conn, addr, make(chan *state.World))

	s.clients[s.maxID] = client
	s.maxID++

	go client.ListenWrite()
}

func (s *Service) RemoveClient(id int64) {
	client, ok := s.clients[id]
	if ok {
		client.Stop()
		delete(s.clients, id)
	}
}

func (s *Service) ReconnectClient(id int64, addr *net.UDPAddr) {
	client := NewClient(id, s.conn, addr, make(chan *state.World))
	s.clients[id] = client

	go client.ListenWrite()
}
