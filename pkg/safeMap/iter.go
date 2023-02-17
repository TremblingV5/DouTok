package safeMap

import cmap "github.com/orcaman/concurrent-map"

func (m *SafeMap) Iter(f cmap.IterCb) {
	m.Mu.RLock()
	defer m.Mu.Unlock()

	m.Map.IterCb(f)
}
