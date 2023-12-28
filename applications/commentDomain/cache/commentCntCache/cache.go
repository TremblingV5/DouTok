package commentCntCache

import (
	"fmt"
	"github.com/TremblingV5/DouTok/pkg/cache"
	"strconv"
)

type Cache struct {
	cache *cache.MapCache[int64]
}

func NewCache() *Cache {
	mapCache := cache.NewMapCache[int64]()
	return &Cache{
		cache: mapCache,
	}
}

func (c *Cache) Add(videoId int64, modification int64) {
	c.cache.Update(fmt.Sprint(videoId), modification, func(exist bool, valueInMap int64, newValue int64) int64 {
		if exist {
			return valueInMap + modification
		} else {
			return modification
		}
	})
}

func (c *Cache) Get(videoId int64) (int64, bool) {
	return c.cache.Get(fmt.Sprint(videoId))
}

func (c *Cache) GetAll() map[int64]int64 {
	result := make(map[int64]int64)
	c.cache.Iter(func(key string, v int64) {
		i64, _ := strconv.ParseInt(key, 10, 64)
		result[i64] = v
	})
	return result
}

func (c *Cache) Clear() {
	c.cache.Clear()
}
