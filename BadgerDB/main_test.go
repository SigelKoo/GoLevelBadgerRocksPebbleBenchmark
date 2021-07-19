package BadgerDB

import (
	config "GoLevelBadgerRocksPebbleBenchmark/config"
	"testing"
)

func BenchmarkWithoutMap(b *testing.B) {
	db := config.BadgerConn()
	bdbModel := Badgerdb{
		db: db,
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = bdbModel.WriteBatchData(1000)
	}
	b.StopTimer()
	db.Close()
}
