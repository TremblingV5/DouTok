package cache

import (
	"errors"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type MapCache[V any] struct {
	m *cmap.ConcurrentMap[string, V]
}

var KeyNotFound = errors.New("key not found")

func NewMapCache[V any]() *MapCache[V] {
	m := cmap.New[V]()
	return &MapCache[V]{
		m: &m,
	}
}

func (c *MapCache[V]) Set(key string, value V) {
	c.m.Set(key, value)
}

func (c *MapCache[V]) Get(key string) (V, bool) {
	v, ok := c.m.Get(key)
	if ok {
		return v, ok
	} else {
		var zero V
		return zero, ok
	}
}

func (c *MapCache[V]) Update(
	key string, value V,
	f func(exist bool, valueInMap V, newValue V) V,
) V {
	return c.m.Upsert(key, value, f)
}

func (c *MapCache[V]) Remove(key string) {
	c.m.Remove(key)
}

func (c *MapCache[V]) Iter(f func(key string, v V)) {
	c.m.IterCb(f)
}

func (c *MapCache[V]) Count() int {
	return c.m.Count()
}

func (c *MapCache[V]) Clear() {
	keys := make([]string, 0, c.m.Count())
	c.Iter(func(key string, v V) {
		keys = append(keys, key)
	})
	for _, key := range keys {
		c.m.Remove(key)
	}
}
