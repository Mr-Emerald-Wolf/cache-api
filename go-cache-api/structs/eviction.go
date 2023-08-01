package structs

type EvictionPolicy interface {
	Evict(c *Cache)
	UpdateCacheAccess(c *Cache, key string)
}
