package structs

import (
	"container/list"
	"sync"
)

type Cache struct {
	Capacity       int
	Items          map[string]interface{}
	Queue          *list.List
	Mutex          sync.Mutex
	EvictionPolicy EvictionPolicy // Custom eviction policy

}

func NewCache(Capacity int, policy EvictionPolicy) *Cache {
	return &Cache{
		Capacity:       Capacity,
		Items:          make(map[string]interface{}),
		Queue:          list.New(),
		EvictionPolicy: policy,
	}
}

func (c *Cache) removeFromQueue(key string) {
	for e := c.Queue.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			c.Queue.Remove(e)
			break
		}
	}
}

func (c *Cache) Clear() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Items = make(map[string]interface{})
	c.Queue = list.New()
}

// Delete removes the given key from the cache.
func (c *Cache) Delete(key string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if _, found := c.Items[key]; found {
		delete(c.Items, key)
		c.removeFromQueue(key)
	}
}

func (c *Cache) Length() int {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	return len(c.Items)
}

// Put adds an item to the cache using the selected eviction policy.
func (c *Cache) Put(key string, value interface{}) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if _, ok := c.Items[key]; !ok {
		if len(c.Items) >= c.Capacity {
			c.EvictionPolicy.Evict(c) // Evict based on the custom eviction policy
		}
	}

	c.Items[key] = value
	c.Queue.PushBack(key)

	// Call updateCacheAccess to reflect the newly added item in the cache access order (if eviction policy is used).
	c.EvictionPolicy.UpdateCacheAccess(c, key)
	return nil 
}

func (c *Cache) Get(key string) (value interface{}, found bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	value, found = c.Items[key]
	if found {
		// No need to call the eviction policy when the key is found (cache hit).
		// We only need to update the cache access order based on the eviction policy.
		c.EvictionPolicy.UpdateCacheAccess(c, key)
	}

	return value, found
}
