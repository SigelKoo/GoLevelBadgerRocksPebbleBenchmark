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

/*func (ldb GoLeveldb) WriteBatchData(data map[string]string) (bool, error) {
	batch := new(leveldb.Batch)

	for k, v := range data {
		batch.Put([]byte(k), []byte(v))
	}

	err := ldb.db.Write(batch, nil)
	if err != nil {
		log.Fatal(err)·
	}
	return true, nil
}*/

func (ldb GoLeveldb) WriteBatchData(num int) (bool, error) {
	batch := new(leveldb.Batch)

	for i := 0; i < num; i++ {
		batch.Put([]byte(config.RandStr(64)), []byte(config.RandStr(1024)))
	}

	err := ldb.db.Write(batch, &opt.WriteOptions{Sync: false})
	if err != nil {
		log.Fatal(err)
	}
	return true, nil
}

