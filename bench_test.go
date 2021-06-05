package main

import (
	"fmt"
	"testing"

	"github.com/nakabonne/tstorage"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"
)

func BenchmarkStorage_Insert(b *testing.B) {
	tmpDir := b.TempDir()
	storage, err := tstorage.NewStorage(
		tstorage.WithDataPath(tmpDir),
	)
	require.NoError(b, err)
	b.Cleanup(func() {
		storage.Close()
	})

	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		err := storage.InsertRows([]tstorage.Row{
			{Metric: "metric1", DataPoint: tstorage.DataPoint{Timestamp: int64(i), Value: 0.1}},
		})
		require.NoError(b, err)
	}
}

func BenchmarkLevelDB_Insert(b *testing.B) {
	tmpDir := b.TempDir()
	db, err := leveldb.OpenFile(tmpDir, nil)
	require.NoError(b, err)
	b.Cleanup(func() {
		db.Close()
	})

	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		err := db.Put(
			[]byte(fmt.Sprintf("%d-metric1", i)),
			[]byte("0.1"),
			nil,
		)
		require.NoError(b, err)
	}
}

func BenchmarkStorage_InsertParallel(b *testing.B) {
	tmpDir := b.TempDir()
	storage, err := tstorage.NewStorage(
		tstorage.WithDataPath(tmpDir),
	)
	require.NoError(b, err)
	b.Cleanup(func() {
		storage.Close()
	})

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			err := storage.InsertRows([]tstorage.Row{
				{Metric: "metric1", DataPoint: tstorage.DataPoint{Timestamp: int64(i), Value: 0.1}},
			})
			require.NoError(b, err)
			i++
		}
	})
}

func BenchmarkLevelDB_InsertParallel(b *testing.B) {
	tmpDir := b.TempDir()
	db, err := leveldb.OpenFile(tmpDir, nil)
	require.NoError(b, err)
	b.Cleanup(func() {
		db.Close()
	})

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			err := db.Put(
				[]byte(fmt.Sprintf("%d-metric1", i)),
				[]byte("0.1"),
				nil,
			)
			require.NoError(b, err)
			i++
		}
	})
}

/*
func BenchmarkTstorage_Select(b *testing.B) {
	tmpDir := b.TempDir()
	storage, err := tstorage.NewStorage(
		tstorage.WithDataPath(tmpDir),
	)
	require.NoError(b, err)
	for i := 1; i < 1000000; i++ {
		err := storage.InsertRows([]tstorage.Row{
			{Metric: "metric1", DataPoint: tstorage.DataPoint{Timestamp: int64(i), Value: 0.1}},
		})
		require.NoError(b, err)
	}

	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		_, _ = storage.Select("metric1", nil, 10, 100)
	}
}

func BenchmarkLevelDB_Select(b *testing.B) {
	tmpDir := b.TempDir()
	db, err := leveldb.OpenFile(tmpDir, nil)
	require.NoError(b, err)
	defer db.Close()

	require.NoError(b, err)
	for i := 1; i < 1000000; i++ {
		err := db.Put(
			[]byte(fmt.Sprintf("%d-metric1", i)),
			[]byte("0.1"),
			nil,
		)
		require.NoError(b, err)
	}
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		_, _ = db.NewIterator(util.BytesPrefix())
	}
}
*/
