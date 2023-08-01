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
		Capacity: Capacity,
		Items:    make(map[string]interface{}),
		Queue:    list.New(),
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


