package safeMap

func (m *SafeMap) Clean() {
	keys := make([]string, 0, m.Map.Count())
	m.Iter(func(key string, v interface{}) {
		keys = append(keys, key)
	})
	for _, key := range keys {
		m.Set(key, int64(0))
	}
}
