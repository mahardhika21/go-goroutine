package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld(name string){
	fmt.Println("hello ", name);
}

func TestCreateGorotine(t *testing.T) {
	go RunHelloWorld("zul") // ada kemunkinan proses di kil jika funsi terkir selesai di eksekusi dan ini belum di eksekusi, solusi menggunkann time utk menunggu
	// RunHelloWorld("zul")
	fmt.Println("ups");

	time.Sleep(1 * time.Second)
}


func DisplayNumber(number int) {
	fmt.Println("display number goruotine", number)
}

func TestManyGoroutine(t *testing.T) {
	for i :=0; i<100000; i++ {
		go DisplayNumber(i)
	}

	//time.Sleep(20 * time.Second)
}