package RocksDB

import (
	config "GoLevelBadgerRocksPebbleBenchmark/config"
	"testing"
)

func BenchmarkWithoutMap(b *testing.B) {
	db := config.RocksDbConn()
	rdbModel := Rocksdb{
		db: db,
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = rdbModel.WriteBatchData(1000)
	}
	b.StopTimer()
	db.Close()
}