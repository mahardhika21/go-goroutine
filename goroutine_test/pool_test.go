package goroutine_test

import (
	"testing"
	"fmt"
	"time"
	"sync"
)

func TestPool(t *testing.T) {
	pool := sync.Pool {
		New: func() interface {} {
			return "New"
		},
	}
	pool.Put("Naruto")
	pool.Put("Sasuke")
	pool.Put("Sakura")

	for i:=0; i<10; i++ {
		go func () {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("selesai")
}