package goroutine_test

import (
	"testing"
	"fmt"
	_"time"
	"sync"
)

var counter = 0;

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}
	//var mutex sync.Mutex
	for i :=0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			// mutex.Lock()
			// OnlyOnce()
			// mutex.Unlock()
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("counter ",counter)
}