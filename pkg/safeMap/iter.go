package safeMap

import cmap "github.com/orcaman/concurrent-map"

func (m *SafeMap) Iter(f cmap.IterCb) {
	m.Map.IterCb(f)
}
