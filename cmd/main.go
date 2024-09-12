package main

import (
	"fmt"
	"time"

	"go-snippets/internal/utils"
)

func main() {
	cache := utils.NewCache()
	cache.Set("hello", "hello world", 3*time.Second)
	fmt.Println(cache.Get("hello"))
	time.Sleep(1 * time.Second)
	fmt.Println(cache.Get("hello"))
	time.Sleep(2 * time.Second)
	fmt.Println(cache.Get("hello"))
}
