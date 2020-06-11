package lru

import (
	"fmt"
	"testing"
)

func TestCache_Get(t *testing.T) {
	cache := New(4)
	cache.Set("1", "1")
	cache.Set("2", "2")
	cache.Set("3", "3")
	cache.Set("4", "4")
	cache.Set("5", "5")

	all := cache.GetAll()
	fmt.Println(all)

	cache.Rem("3")

	all2 := cache.GetAll()
	fmt.Println(all2)
}