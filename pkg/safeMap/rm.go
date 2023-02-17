package safeMap

func (m *SafeMap) Remove(key string) {
	m.Map.Remove(key)
}
