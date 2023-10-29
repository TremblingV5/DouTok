package cache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMapCache(t *testing.T) {
	stringCache := NewMapCache[string]()

	stringCache.Set("1", "1")
	stringCache.Set("2", "2")

	v1, ok := stringCache.Get("1")
	require.Equal(t, "1", v1)
	require.Equal(t, true, ok)

	_, ok = stringCache.Get("999")
	require.Equal(t, false, ok)

	v2, ok := stringCache.Get("2")
	require.Equal(t, "2", v2)
	require.Equal(t, true, ok)
	stringCache.Remove("2")

	_, ok = stringCache.Get("2")
	require.Equal(t, false, ok)

	cnt := stringCache.Count()
	require.Equal(t, 1, cnt)
	stringCache.Clear()
	cnt = stringCache.Count()
	require.Equal(t, 0, cnt)

	integerCache := NewMapCache[int64]()
	integerCache.Set("1", 1)
	v3, ok := integerCache.Get("1")
	require.Equal(t, int64(1), v3)
	require.Equal(t, true, ok)
}
