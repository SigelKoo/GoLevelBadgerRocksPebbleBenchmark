package RocksDB

import (
	config "GoLevelBadgerRocksPebbleBenchmark/config"
	"log"

	"github.com/tecbot/gorocksdb"
)

type Rocksdb struct {
	db *gorocksdb.DB
}

func (rdb Rocksdb) WriteBatchData(num int) (error, bool) {
	wo := gorocksdb.NewDefaultWriteOptions()
	wo.SetSync(false)
	wb := gorocksdb.NewWriteBatch()
	for i := 0; i < num; i++ {
		wb.Put(config.RandStr(64), config.RandStr(1024))
	}
	err := rdb.db.Write(wo, wb)
	if err != nil {
		log.Fatal(err)
	}
	return nil, true
}
