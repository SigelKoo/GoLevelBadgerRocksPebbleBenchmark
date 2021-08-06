ulimit -a
cd /home/gopath/GoLevelBadgerRocksPebbleBenchmark/BadgerDB
CGO_CFLAGS="-I/usr/local/rocksdb/include" CGO_LDFLAGS="-L/usr/local/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" go test -bench=. -run=none -benchtime=100000x -timeout 24h
cd /tmp/badgertest-0/dbbench
du --total
cd /home/gopath/GoLevelBadgerRocksPebbleBenchmark
./clean.sh
cd /home/gopath/GoLevelBadgerRocksPebbleBenchmark/PebbleDB
CGO_CFLAGS="-I/usr/local/rocksdb/include" CGO_LDFLAGS="-L/usr/local/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" go test -bench=. -run=none -benchtime=100000x -timeout 24h
cd /tmp/pebbledbtest-0/dbbench
du --total
cd /home/gopath/GoLevelBadgerRocksPebbleBenchmark
./clean.sh
cd /home/gopath/GoLevelBadgerRocksPebbleBenchmark/GoLevelDB
CGO_CFLAGS="-I/usr/local/rocksdb/include" CGO_LDFLAGS="-L/usr/local/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" go test -bench=. -run=none -benchtime=100000x -timeout 24h
cd /tmp/leveldbtest-1/dbbench
du --total
cd /home/gopath/GoLevelBadgerRocksPebbleBenchmark
./clean.sh
cd /home/gopath/GoLevelBadgerRocksPebbleBenchmark/RocksDB
CGO_CFLAGS="-I/usr/local/rocksdb/include" CGO_LDFLAGS="-L/usr/local/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" go test -bench=. -run=none -benchtime=100000x -timeout 24h
cd /tmp/rocksdbtest-0/dbbench
du --total
cd /home/gopath/GoLevelBadgerRocksPebbleBenchmark
./clean.sh