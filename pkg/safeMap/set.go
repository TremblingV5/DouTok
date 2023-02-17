package safeMap

func (m *SafeMap) Set(key string, value any) {
	m.Map.Set(key, value)
}
