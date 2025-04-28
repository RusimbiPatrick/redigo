package persistence

import (
	"os"
	"sync"
	"time"

	"github.com/containerd/containerd/cmd/containerd/command"
)


type AOF struct {
	file *os.File
	mu  sync.Mutex
	interval time.Duration
}


func (a *AOF) Append(command string, args ...string)  {
	a.mu.Lock()
	defer a.mu.Unlock()
	//write Redis protocol formated command to file
}