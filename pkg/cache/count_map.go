package cache

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
)

type CountMapCache struct {
	countMap cmap.ConcurrentMap
}

func NewCountMapCache() *CountMapCache {
	return &CountMapCache{
		countMap: cmap.New(),
	}
}

func (c *CountMapCache) Get(id int64) (int64, bool) {
	value, ok := c.countMap.Get(fmt.Sprint(id))
	if !ok {
		return 0, false
	}

	commentCount, ok := value.(int64)
	if !ok {
		return 0, false
	}

	return commentCount, true
}

func (c *CountMapCache) Set(id, value int64) {
	c.countMap.Set(fmt.Sprint(id), value)
}

func (c *CountMapCache) Add(id, increment int64) {
	value, ok := c.Get(id)
	if ok {
		increment += value

	}

	c.countMap.Set(fmt.Sprint(id), increment)
}

func (c *CountMapCache) Iter(f cmap.IterCb) {
	c.countMap.IterCb(f)
}
