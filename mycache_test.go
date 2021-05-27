package lrucache

import (
	"fmt"
	"testing"
	"time"
)

func TestLruCache(t *testing.T) {
	cache, err := New(1, 1)
	if err != nil {
		t.Fatal(err)
	}

	cache.Set("foo", "bar")
	if data, err := cache.Get("foo"); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(data)
	}

	time.Sleep(2 * time.Second)
	if _, err := cache.Get("foo"); err != nil {
		fmt.Printf("expired: %v, non-existed: %v\n", IsExpired(err), IsNonExisted(err))
	} else {
		t.Fatal()
	}

	cache.Set("fff", "bbb")
	if _, err := cache.Get("foo"); err != nil {
		fmt.Printf("expired: %v, non-existed: %v\n", IsExpired(err), IsNonExisted(err))
	} else {
		t.Fatal()
	}
}
