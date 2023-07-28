package structs

import (
	"container/list"
	"sync"
)

type Cache struct {
	Capacity int
	Items    map[string]interface{}
	Queue    *list.List
	Mutex    sync.Mutex
}

func NewCache(Capacity int) *Cache {
	return &Cache{
		Capacity: Capacity,
		Items:    make(map[string]interface{}),
		Queue:    list.New(),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if value, found := c.Items[key]; found {
		c.Queue.MoveToFront(value.(*list.Element))
		return value, true
	}

	return nil, false
}

func (c *Cache) Set(key string, value interface{}) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if existing, found := c.Items[key]; found {
		c.Queue.MoveToFront(existing.(*list.Element))
		c.Items[key] = c.Queue.Front()
		return
	}

	if c.Queue.Len() >= c.Capacity {
		back := c.Queue.Back()
		delete(c.Items, back.Value.(string))
		c.Queue.Remove(back)
	}

	element := c.Queue.PushFront(key)
	c.Items[key] = element
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

