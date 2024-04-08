package commentCntCache

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

func assertEqual(t *testing.T, expect int, actual int64) {
	require.Equal(t, int64(expect), actual)
}

func TestCache(t *testing.T) {
	cache := NewCache()

	cache.Add(1, 1)
	v1, ok := cache.Get(1)
	assertEqual(t, 1, v1)
	require.Equal(t, true, ok)
	_, ok = cache.Get(2)
	require.Equal(t, false, ok)

	cache.Add(1, 5)
	v1, ok = cache.Get(1)
	assertEqual(t, 6, v1)
	require.Equal(t, true, ok)

	cache.Add(1, -13)
	v1, ok = cache.Get(1)
	assertEqual(t, -7, v1)
	require.Equal(t, true, ok)

	cache.Clear()
	_, ok = cache.Get(1)
	require.Equal(t, false, ok)
}

func TestCacheMultiThread(t *testing.T) {
	cache := NewCache()
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			cache.Add(1, 1)

			cache.Get(1)
		}()
	}

	wg.Wait()

	value, ok := cache.Get(1)
	require.Equal(t, int64(100), value)
	require.Equal(t, true, ok)
}
