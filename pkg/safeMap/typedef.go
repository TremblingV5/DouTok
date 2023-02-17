package safeMap

import (
	cmap "github.com/orcaman/concurrent-map"
)

type SafeMap struct {
	Map cmap.ConcurrentMap
}
