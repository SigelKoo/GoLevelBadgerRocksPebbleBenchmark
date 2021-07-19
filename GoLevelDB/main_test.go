package GoLevelDB

import (
	config "GoLevelBadgerRocksPebbleBenchmark/config"
	"testing"
)

func BenchmarkWithoutMap(b *testing.B) {
	db := config.GoLevelDbConn()
	ldbModel := GoLeveldb{
		db: db,
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = ldbModel.WriteBatchData(1000)
	}
	b.StopTimer()
	db.Close()
}

