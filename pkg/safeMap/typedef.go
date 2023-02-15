package safeMap

import (
	"sync"

	cmap "github.com/orcaman/concurrent-map"
)

type SafeMap struct {
	Map cmap.ConcurrentMap
	Mu  sync.RWMutex
}
