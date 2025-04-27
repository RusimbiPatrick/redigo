package main
import (
	"log"
	"net"
	"github.com/RusimbiPatrick/redigo/internal/server"
)

func main() {
	//Todo get dynamic port
	listen, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	srv := server.NewServer()
	log.Println("server listening on :6379")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go srv.HandleConn(conn)
	}
}