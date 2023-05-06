package goroutine_test

import (
	"testing"
	"fmt"
	"time"
	"sync"
)

func TestRaceCOndition(t *testing.T) {
	x := 0;

	for i := 0; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1;
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("counter :", x)
}

func TestRaceCOnditionMutex(t *testing.T) { // locking
	x := 0;
	var mutex sync.Mutex

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1;
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("counter :", x)
}

func TestRwMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i<100; i ++ {
		go func () {
			for j := 1; j < 100; j++ {
				account.AddBalance(j)
				fmt.Println(account.GetBalance())
			}
		} ()
	}

	time.Sleep(5 * time.Second);

	fmt.Println("total balance", account.GetBalance())
}

type BankAccount struct {
	RwMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance (ammount int) {
	account.RwMutex.Lock()
	account.Balance = account.Balance + ammount
	account.RwMutex.Unlock()
}

func (account *BankAccount) GetBalance () int {
	account.RwMutex.RLock()
	balance := account.Balance
	account.RwMutex.RUnlock()

	return balance
}

// simulasi deadlock

type UserBalance struct {
	Mutex sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change (amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock();


}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance {
		Name : "Naruto",
		Balance: 1000000,
	}

	user2 := UserBalance {
		Name: "Sasuke",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)

	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second) // alternatif withGroup()
	fmt.Println("User1 ", user1.Name," Balance ", user1.Balance);

	fmt.Println("User2 ", user2.Name," Balance ", user2.Balance);
}

