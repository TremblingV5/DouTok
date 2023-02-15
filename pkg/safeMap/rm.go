package safeMap

func (m *SafeMap) Remove(key string) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	m.Map.Remove(key)
}
