package goroutine_test

import (
	"testing"
	"fmt"
	_"time"
	"sync"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaiCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	
	fmt.Prinln("Done ", value)

	cond.L.Unlock()
}

func TesCond(t *testing.T) {
	for i:=0; i<10; i++ {
		go WaiCondition(i)
	}

	group.Wait()
}