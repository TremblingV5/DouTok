package safeMap

func (m *SafeMap) Set(key string, value any) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	m.Map.Set(key, value)
}
