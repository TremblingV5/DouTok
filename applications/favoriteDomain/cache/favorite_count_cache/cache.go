package favorite_count_cache

import (
	"fmt"
	"github.com/TremblingV5/DouTok/pkg/cache"
)

type Cache struct {
	cache *cache.MapCache[int64]
}

func New() *Cache {
	return &Cache{
		cache: cache.NewMapCache[int64](),
	}
}

func (c *Cache) Get(videoId int64) (int64, bool) {
	return c.cache.Get(fmt.Sprint(videoId))
}

func (c *Cache) Set(videoId, count int64) {
	c.cache.Set(fmt.Sprint(videoId), count)
}

func (c *Cache) SetBatch(bacth map[int64]int64) {
	for k, v := range bacth {
		c.cache.Set(fmt.Sprint(k), v)
	}
}

func (c *Cache) Clear() {
	c.cache.Clear()
}
