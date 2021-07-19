package PebbleDB

import (
	config "GoLevelBadgerRocksPebbleBenchmark/config"
	"testing"
)

func BenchmarkWithoutMap(b *testing.B) {
	db := config.PebbleConn()
	pdbModel := Pebbledb{
		db: db,
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = pdbModel.WriteBatchData(1000)
	}
	b.StopTimer()
	db.Close()
}

