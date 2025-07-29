package mutex

import (
	"fmt"
	"math/rand"
	"sync"
)

type Account struct {
	id      int
	balance int
	mu      sync.Mutex
}

func NewAccount(id int, balance int) *Account {
	return &Account{
		id:      id,
		balance: balance,
	}
}

func (a *Account) Deposit(amount int) bool {
	// Sleep(2)
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
	return true
}

func (a *Account) GetBalance() int {
	// Sleep(2)
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func (a *Account) Withdraw(amount int) bool {
	// Sleep(2)
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < amount {
		return false
	}
	a.balance -= amount
	return true
}

func Transfer(from *Account, to *Account, amount int) bool {
	if from.Withdraw(amount) {
		return to.Deposit(amount)
	}
	fmt.Println("Insufficient balance")
	return false
}

func SimulateMutex() {
	accounts := make(map[int]*Account)
	rand := rand.New(rand.NewSource(4))
	totAccounts := 10
	totTransfers := 10
	for i := range totAccounts {
		accounts[i] = NewAccount(i, rand.Intn(1000))
	}
	var wg sync.WaitGroup
	for range totTransfers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a := rand.Intn(totAccounts)
			b := rand.Intn(totAccounts)
			if a != b {
				Transfer(accounts[a], accounts[b], rand.Intn(100))
			}
		}()
	}
	wg.Wait()
	xor := 0
	for _, account := range accounts {
		balance := account.GetBalance()
		fmt.Println(balance)
		xor ^= balance
	}
	fmt.Println("hash =", xor)
}

// Need to work on atomicity
