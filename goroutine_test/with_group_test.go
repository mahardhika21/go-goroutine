package goroutine_test

import (
	"testing"
	"fmt"
	"time"
	"sync"
)

func RunAsyncronus(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")

	time.Sleep(1 * time.Second)
}

func TestWaitGroupData(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronus(group)
	}

	group.Wait()
	fmt.Println("selesai")
}