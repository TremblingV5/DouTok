package safeMap

func (m *SafeMap) Get(key string) (any, bool) {
	return m.Map.Get(key)
}
