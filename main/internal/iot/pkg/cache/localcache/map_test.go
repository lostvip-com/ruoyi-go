package localcache

import (
	"fmt"
	"testing"

	"github.com/patrickmn/go-cache"
)

func TestNewRamCacheClient(t *testing.T) {
	c := NewRamCacheClient()
	var (
		key   = "hello"
		value = "world"
	)
	c.Set(key, value, cache.DefaultExpiration)
	fmt.Println(c.Get(key))
	c.Del(key)
	fmt.Println(c.Get(key))

	fmt.Println("=====================")

	var (
		hashKey = "hash_test"
		field   = "boo"
		values  = []interface{}{"boo", "far", "boo1", "far1"}
	)
	fmt.Println("hash get " + hashKey)
	fmt.Println(c.HGet(hashKey, field))
	fmt.Println("=====================")
	fmt.Println("hash set " + hashKey)
	fmt.Println(c.HSet(hashKey, values...))
	fmt.Println("=====================")
	fmt.Println("hash get " + hashKey)
	fmt.Println(c.HGet(hashKey, field))
	fmt.Println("=====================")
	fmt.Println("hash get all")
	fmt.Println(c.HGetAll(hashKey))
	fmt.Println("=====================")
	fmt.Println("hash delete " + field)
	fmt.Println(c.HDel(hashKey, field))
	fmt.Println("=====================")
	fmt.Println("hash get all")
	fmt.Println(c.HGetAll(hashKey))
}
