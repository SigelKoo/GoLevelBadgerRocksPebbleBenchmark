package PebbleDB

import (
	config "GoLevelBadgerRocksPebbleBenchmark/config"

	"github.com/cockroachdb/pebble"
)

type Pebbledb struct {
	db *pebble.DB
}

func (pdb Pebbledb) WriteBatchData(num int) (bool, error) {
	batch := pdb.db.NewBatch()
	defer batch.Close()

	for i := 0; i < num; i++ {
		batch.Set([]byte(config.RandStr(64)), []byte(config.RandStr(1024)), &pebble.WriteOptions{Sync: false})
	}

	batch.Commit(&pebble.WriteOptions{Sync: false})
	return true, nil
}

