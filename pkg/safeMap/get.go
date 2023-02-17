package safeMap

func (m *SafeMap) Get(key string) (any, bool) {
	m.Mu.RLock()
	defer m.Mu.RUnlock()

	return m.Map.Get(key)
}
