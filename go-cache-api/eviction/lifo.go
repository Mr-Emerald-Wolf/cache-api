package eviction

import "github.com/mr-emerald-wolf/go-cache-api/structs"

// LIFOEvictionPolicy implements the Last-In-First-Out eviction policy.
type LIFOEvictionPolicy struct{}

func (LIFOEvictionPolicy) Evict(c *structs.Cache) {
	// Evict the newest item in the queue (LIFO)
	newestElement := c.Queue.Back()
	if newestElement != nil {
		newestKey := newestElement.Value.(string)
		delete(c.Items, newestKey)
		c.Queue.Remove(newestElement)
	}
}

func (LIFOEvictionPolicy) UpdateCacheAccess(c *structs.Cache, key string) {
	
}
