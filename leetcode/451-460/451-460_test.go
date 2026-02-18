package _451_460

import "testing"

func TestLFUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	if cache.Get(1) != 1 {
		t.Error("get 1 should be 1")
	}
	cache.Put(2, 2)
	if cache.Get(2) != 2 {
		t.Error("get 2 should be 2")
	}
	cache.Put(3, 3)
	if cache.Get(2) != 2 {
		t.Error("get 2 should be 2")
	}
	if cache.Get(1) != -1 {
		t.Error("get 1 should be -1")
	}
}
