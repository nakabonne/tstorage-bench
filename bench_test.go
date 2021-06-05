package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/nakabonne/tstorage"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"
)

func BenchmarkStorage_Insert(b *testing.B) {
	tmpDir, err := ioutil.TempDir("", "tstorage-bench")
	require.NoError(b, err)
	defer os.RemoveAll(tmpDir)

	storage, err := tstorage.NewStorage(
		tstorage.WithDataPath(tmpDir),
	)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		storage.InsertRows([]tstorage.Row{
			{Metric: "metric1", DataPoint: tstorage.DataPoint{Timestamp: int64(i), Value: 0.1}},
		})
	}
}

func BenchmarkLevelDB_Insert(b *testing.B) {
	tmpDir, err := ioutil.TempDir("", "tstorage-bench")
	require.NoError(b, err)
	defer os.RemoveAll(tmpDir)

	db, err := leveldb.OpenFile(tmpDir, nil)
	require.NoError(b, err)
	defer db.Close()

	b.ResetTimer()
	for i := 1; i < b.N; i++ {
	}
}

func BenchmarkTstorage_Select(b *testing.B) {
	tmpDir, err := ioutil.TempDir("", "tstorage-bench")
	require.NoError(b, err)
	defer os.RemoveAll(tmpDir)

	storage, err := tstorage.NewStorage(
		tstorage.WithDataPath(tmpDir),
	)
	require.NoError(b, err)
	for i := 1; i < 1000000; i++ {
		storage.InsertRows([]tstorage.Row{
			{Metric: "metric1", DataPoint: tstorage.DataPoint{Timestamp: int64(i), Value: 0.1}},
		})
	}
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		_, _ = storage.Select("metric1", nil, 10, 100)
	}
}
