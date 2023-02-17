package safeMap

import cmap "github.com/orcaman/concurrent-map"

func New() *SafeMap {
	return &SafeMap{
		Map: cmap.New(),
	}
}
