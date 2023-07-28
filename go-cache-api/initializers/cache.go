package initializers

import "github.com/mr-emerald-wolf/go-cache-api/structs"


type CacheMemory struct {
	LRU  *structs.Cache
	FIFO *structs.Cache
	LIFO *structs.Cache
}

func InitCache() *CacheMemory {
	lru := structs.NewCache(3)
	fifo := structs.NewCache(3)
	lifo := structs.NewCache(3)

	return &CacheMemory{
		LRU:  lru,
		FIFO: fifo,
		LIFO: lifo,
	}
}
