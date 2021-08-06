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
		batch.Set(config.RandStr(i), []byte(config.Value1024), &pebble.WriteOptions{Sync: true})
	}

	batch.Commit(&pebble.WriteOptions{Sync: true})
	return true, nil
}
