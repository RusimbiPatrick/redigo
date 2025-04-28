package persistence

import (
	"time"

	"github.com/RusimbiPatrick/redigo/internal/storage"
)

type RDB struct {
	saveInterval time.Duration
	lastSave time.Time
}

func (r *RDB) Save(engine *storage.Engine) error {
	//Serialize in-memory data to disk
}
