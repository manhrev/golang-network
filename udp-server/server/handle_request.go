package server

import (
	"strconv"
	"strings"
)

type ReqType string

const (
	reqType_NONE       = "none"
	reqType_CONNECT    = "connect"
	reqType_RECONNECT  = "reconnect"
	reqType_DISCONNECT = "disconnect"
)

func (s *Server) HandleRequest() error {
	for {
		var buf [512]byte
		l, caddr, err := s.conn.ReadFromUDP(buf[0:])
		if err != nil {
			return err
		}

		connectString := string(buf[:l])
		parts := strings.Split(connectString[:len(connectString)-1], " ")
		if len(parts) < 1 {
			continue
		}

		switch parts[0] {
		case reqType_CONNECT:
			s.clientManager.AddNewClient(caddr)
		case reqType_DISCONNECT:
			if len(parts) != 2 {
				continue
			}

			id, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}

			s.clientManager.RemoveClient(int64(id))
		case reqType_RECONNECT:
			if len(parts) != 2 {
				continue
			}

			id, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}

			s.clientManager.ReconnectClient(int64(id), caddr)
		default:
			continue
		}
	}
}
