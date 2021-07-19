package config

import (
	"log"
	"math/rand"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/dgraph-io/badger/v3"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/tecbot/gorocksdb"
)

func GoLevelDbConn() *leveldb.DB {
	opts := &opt.Options{NoSync: true}
	db, err := leveldb.OpenFile("/tmp/leveldbtest-1/dbbench", opts)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func BadgerConn() *badger.DB {
	opts := badger.DefaultOptions("/tmp/badgertest-0/dbbench")
	opts.SyncWrites = false
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func PebbleConn() *pebble.DB {
	db, err := pebble.Open("/tmp/pebbledbtest-0/dbbench", &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func RocksDbConn() *gorocksdb.DB {
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "/tmp/rocksdbtest-0/dbbench")
	if err != nil {
		log.Fatal("err")
	}
	return db
}

func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

