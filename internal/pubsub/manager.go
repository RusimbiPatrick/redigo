package pubsub

import (
	"net"
	"sync"
)

type PubSub struct {
	mu sync.RWMutex
	channel map[string][]net.Conn
}

func(p *PubSub) Subscribe(conn net.Conn, channels ...string) {
	//Add connection to channels lists
}

func(p *PubSub) Publish(channel string, message string) int {
	//Broadcast message to all subscribers
}