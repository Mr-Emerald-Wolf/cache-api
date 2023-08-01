package initializers

import (
	"github.com/mr-emerald-wolf/go-cache-api/eviction"
	"github.com/mr-emerald-wolf/go-cache-api/structs"
)

type CacheMemory struct {
	LRU  *structs.Cache
	FIFO *structs.Cache
	LIFO *structs.Cache
}

func InitCache() *CacheMemory {
	
	lru := structs.NewCache(3, &eviction.LRUEvictionPolicy{})
	fifo := structs.NewCache(3, &eviction.FIFOEvictionPolicy{})
	lifo := structs.NewCache(3, &eviction.LIFOEvictionPolicy{})

	return &CacheMemory{
		LRU:  lru,
		FIFO: fifo,
		LIFO: lifo,
	}
}
