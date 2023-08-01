package eviction

import "github.com/mr-emerald-wolf/go-cache-api/structs"

// FIFOEvictionPolicy implements the First-In-First-Out eviction policy.
type FIFOEvictionPolicy struct{}

func (FIFOEvictionPolicy) Evict(c *structs.Cache) {
	// Evict the oldest item in the queue (FIFO)
	oldestElement := c.Queue.Front()
	if oldestElement != nil {
		oldestKey := oldestElement.Value.(string)
		delete(c.Items, oldestKey)
		c.Queue.Remove(oldestElement)
	}
}

func (FIFOEvictionPolicy) UpdateCacheAccess(c *structs.Cache, key string) {

}
