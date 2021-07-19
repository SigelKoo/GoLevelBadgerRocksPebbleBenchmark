package BadgerDB

import (
	config "GoLevelBadgerRocksPebbleBenchmark/config"
	"log"

	"github.com/dgraph-io/badger/v3"
)

type Badgerdb struct {
	db *badger.DB
}

func (bdb Badgerdb) WriteBatchData(num int) (error, bool) {
	wb := bdb.db.NewWriteBatch()

	for i := 0; i < num; i++ {
		wb.Set([]byte(config.RandStr(64)), []byte(config.RandStr(1024)))
	}
	err := wb.Flush()
	if err != nil {
		log.Fatal(err)
	}
	return nil, true
}

