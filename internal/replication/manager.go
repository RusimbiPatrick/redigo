package replication

import (
	"net"
	"sync"
	"time"
)

type Replica struct {
	conn net.Conn
	offset int64
	lastAck time.Time
}

type ReplicationManager struct {
	mu sync.RWMutex
	replicas []*Replica
	master bool
}