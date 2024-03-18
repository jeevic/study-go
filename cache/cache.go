package main

import (
	"fmt"
	"time"

	"github.com/muesli/cache2go"
)

type myStruct struct {
	text     string
	moreData []byte
}

func main() {
	cache := cache2go.Cache("myCache")

	val := myStruct{"This is a test!", []byte{}}
	cache.Add("someKey", 5*time.Second, &val)
	res, err := cache.Value("someKey")
	time.Sleep(6 * time.Second)
	if err == nil {
		fmt.Println("Found value in cache:", res.Data().(*myStruct).text)
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	// Wait for the item to expire in cache.
	time.Sleep(6 * time.Second)
	res, err = cache.Value("someKey")
	if err != nil {
		fmt.Println("Item is not cached (anymore).")
	}

	// Add another item that never expires.
	cache.Add("someKey", 0, &val)

	// cache2go supports a few handy callbacks and loading mechanisms.
	cache.SetAboutToDeleteItemCallback(func(e *cache2go.CacheItem) {
		fmt.Println("Deleting:", e.Key(), e.Data().(*myStruct).text, e.CreatedOn())
	})

	// Remove the item from the cache.
	cache.Delete("someKey")

	// And wipe the entire cache table.
	cache.Flush()

}
