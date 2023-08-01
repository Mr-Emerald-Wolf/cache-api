# Cache API with Go Fiber

A powerful and flexible in-memory caching API for general use in Go (Golang) with Go Fiber. This API supports multiple standard eviction policies (FIFO, LRU, LIFO) and allows the addition of custom eviction policies. The API ensures thread safety using mutexes to handle concurrent access.

## Postman Documentation

For detailed API documentation, including request examples, response format, and usage instructions, please refer to the Postman documentation:

[![Postman Documentation](https://img.shields.io/badge/Postman-Documentation-orange)](https://documenter.getpostman.com/view/21877920/2s9XxvTFE7)

## Features

- General-purpose in-memory caching API with Go Fiber
- RESTful endpoints for cache management
- Support for three standard eviction policies: FIFO, LRU, and LIFO
- Option to add custom eviction policies
- Thread-safe caching with the use of mutexes
- Simple and intuitive API design

## Usage

To use this caching API, follow these steps:

```bash
go run main.go
```


Your caching API with Go Fiber is now up and running on `localhost:8080`. You can use the specified endpoints to manage the cache efficiently.

## Endpoints

### LIFO Cache Endpoints

- **`GET /lifo/get`**: Retrieve the value for the specified key from the LIFO cache.
- **`POST /lifo/put`**: Add or update a value in the LIFO cache for the specified key.
- **`GET /lifo/size`**: Get the current size of the LIFO cache.
- **`DELETE /lifo/clear`**: Clear all key-value pairs from the LIFO cache.
- **`GET /lifo/cache`**: Retrieve all key-value pairs from the LIFO cache.

### FIFO Cache Endpoints

- **`GET /fifo/get`**: Retrieve the value for the specified key from the FIFO cache.
- **`POST /fifo/put`**: Add or update a value in the FIFO cache for the specified key.
- **`GET /fifo/size`**: Get the current size of the FIFO cache.
- **`DELETE /fifo/clear`**: Clear all key-value pairs from the FIFO cache.
- **`GET /fifo/cache`**: Retrieve all key-value pairs from the FIFO cache.

### LRU Cache Endpoints

- **`GET /lru/get`**: Retrieve the value for the specified key from the LRU cache.
- **`POST /lru/put`**: Add or update a value in the LRU cache for the specified key.
- **`GET /lru/size`**: Get the current size of the LRU cache.
- **`DELETE /lru/clear`**: Clear all key-value pairs from the LRU cache.
- **`GET /lru/cache`**: Retrieve all key-value pairs from the LRU cache.

## Thread Safety

The caching API ensures thread safety using mutexes. This means that the API can handle concurrent access from multiple goroutines without any data corruption.

## Custom Eviction Policies

To add a custom eviction policy, you need to implement the `structs.EvictionPolicy` interface. By creating a new interface under the eviction folder you can implement a custom eviction policy. Here's an example of how you can create a custom eviction policy: 

```go
type CustomEvictionPolicy struct {
	// Add any custom fields or parameters here
}

func (c *CustomEvictionPolicy) Evict(cache *cache.Cache) {
	// Add your custom eviction logic here
}
func (c *CustomEvictionPolicy) UpdateCacheAccess(cache *cache.Cache) {
	// updateCacheAccess updates the cache access order based on the custom eviction policy 
}
```

Once you have implemented the `Evict` method with your custom eviction logic, you can use it when initializing a new custom cache:

```go
var customCache = structs.NewCache(100, &eviction.CustomEvictionPolicy{})
```

Now, this new cache will use your custom eviction policy.
