package commentTotalCountCache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func assertEqual(t *testing.T, expect int, actual int64) {
	require.Equal(t, int64(expect), actual)
}

func TestCache(t *testing.T) {
	cache := NewCache()

	cache.Set(1, 1)
	v1, ok := cache.Get(1)
	assertEqual(t, 1, v1)
	require.Equal(t, true, ok)

	cache.Set(1, 5)
	v1, ok = cache.Get(1)
	assertEqual(t, 5, v1)
	require.Equal(t, true, ok)

	batchData := make(map[int64]int64)
	batchData[1] = 1
	batchData[2] = 2
	cache.SetBatch(batchData)
	v1, ok = cache.Get(1)
	assertEqual(t, 1, v1)
	require.Equal(t, true, ok)
	v2, ok := cache.Get(2)
	assertEqual(t, 2, v2)
	require.Equal(t, true, ok)

	cache.Clear()
	_, ok = cache.Get(1)
	require.Equal(t, false, ok)
}
