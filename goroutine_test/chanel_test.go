package goroutine_test

import (
	"testing"
	"fmt"
	"time"
)

func TestCreateChanel(t *testing.T) {
	chanel := make(chan string)
	// name := "zul"
	// chanel <- name // kirim data ke chanel
	// name2 := <- chanel // mengirim langsung ke varibel

	// fmt.Println(name2)
	 defer close(chanel)
	go func () {
		time.Sleep(2 * time.Second) 
		chanel <- "zul"
		fmt.Println("selesai mengirim data ke chanel")
	}()

	data := <- chanel

	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestChanelAsParameter(t *testing.T) {
	chanel := make(chan string)

	go GiveMessageResponse(chanel)

	data := <- chanel

	fmt.Println(data)

	close(chanel)
}

func GiveMessageResponse(chanel chan string) {
	time.Sleep(2 * time.Second)

	chanel <- "Zulkifli Mahardhika"
}


func OnlyIn(chanel chan <- string) { // setup chanel hanya bisa utk menerima
	time.Sleep(2 * time.Second)

	chanel <- "naruto uzumaki"

	//data := <- chanel
}

func OnlyOut(chanel <- chan string) { // setup chanel hanya bisa utk mengirim
	data := <- chanel

	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestInOutChanel(t *testing.T) {
	chanel := make(chan string)
	go OnlyIn(chanel)

	OnlyOut(chanel)

	close(chanel)
}

func TestSelectMutltipleChanel(t *testing.T) {
	chanel1 := make(chan string)
	chanel2 := make(chan string)
	couter := 0

	defer close(chanel1)
	defer close(chanel2)

	go GiveMessageResponse(chanel1)

	go GiveMessageResponse(chanel2)

	for {
		select {
		case data := <- chanel1:
			fmt.Println("data dari chanel 1", data)
			couter++
		case data := <- chanel2:
			fmt.Println("data dari chanel 2", data)
			couter++
		}

		if(couter == 2) {
			break
		}
	}
}

func TestSelectMutltipleChanelDefault(t *testing.T) {
	chanel1 := make(chan string)
	chanel2 := make(chan string)
	couter := 0

	defer close(chanel1)
	defer close(chanel2)

	go GiveMessageResponse(chanel1)

	go GiveMessageResponse(chanel2)

	for {
		select {
		case data := <- chanel1:
			fmt.Println("data dari chanel 1", data)
			couter++
		case data := <- chanel2:
			fmt.Println("data dari chanel 2", data)
			couter++
		default :
			fmt.Println("menunggu data")
		}
		if(couter == 2) {
			break
		}
	}
}