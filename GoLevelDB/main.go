package GoLevelDB

import (
	config "GoLevelBadgerRocksPebbleBenchmark/config"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type GoLeveldb struct {
	db *leveldb.DB
}

func (ldb GoLeveldb) WriteBatchData(num int) (bool, error) {
	batch := new(leveldb.Batch)

	for i := 0; i < num; i++ {
		batch.Put(config.RandStr(i), []byte(config.Value1024))
	}

	err := ldb.db.Write(batch, &opt.WriteOptions{Sync: true})
	if err != nil {
		log.Fatal(err)
	}
	return true, nil
}
