package server

import (
	"net"
	"github.com/RusimbiPatrick/redigo/internal/storage"
	"github.com/RusimbiPatrick/redigo/internal/protocol"
)

type Server struct {
	store *storage.Engine
}

func NewServer() *Server {
	return &Server{
		store: storage.NewEngine(),
	}
}

func (s *Server) HandleConn(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 4096)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            return
        }
        cmd, args, err := protocol.Parse(buf[:n])
        if err != nil {
            conn.Write([]byte("Err" + err.Error() + "\r\n"))
            continue
        }
        response := s.handleCommand(cmd, args)
        conn.Write(response)
    }
}

func(s *Server) handleCommand(cmd string, args []string) []byte {
	switch cmd {
	case "PING":
		return []byte("+PONG\r\n")
	case "SET":
		return s.store.Set(args[0], args[1])
	case "GET":
		return s.store.Get(args[0])
	default:
		return []byte("-ERR unknown command\r\n")
	}
}