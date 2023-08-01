package eviction

import "github.com/mr-emerald-wolf/go-cache-api/structs"

// LRUEvictionPolicy implements the Least Recently Used eviction policy.
type LRUEvictionPolicy struct{}

func (LRUEvictionPolicy) Evict(c *structs.Cache) {
	// Evict the least recently used item in the queue (LRU)
	oldestElement := c.Queue.Front()
	if oldestElement != nil {
		oldestKey := oldestElement.Value.(string)
		delete(c.Items, oldestKey)
		c.Queue.Remove(oldestElement)
	}
}

// updateCacheAccess updates the cache access order based on the LRU (Least Recently Used) eviction policy.
func (LRUEvictionPolicy) UpdateCacheAccess(c *structs.Cache, key string) {
	// Move the accessed item to the back of the queue to mark it as the most recently used.
	// This ensures that the least recently used item remains at the front of the queue.
	for e := c.Queue.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			c.Queue.MoveToBack(e)
			break
		}
	}
}
