package channels

import (
	"fmt"
	"math/rand"
	"sync"
)

type deopsitReq struct {
	amount int
	resp   chan bool
}

type withdrawReq struct {
	amount int
	resp   chan bool
}

type balanceReq struct {
	resp chan int
}

type Account struct {
	id              int
	balance         int
	depositReqChnl  chan deopsitReq
	withdrawReqChnl chan withdrawReq
	balanceReqChnl  chan balanceReq
}

func NewAccount(id int, balance int) *Account {
	account := &Account{
		id:              id,
		balance:         balance,
		depositReqChnl:  make(chan deopsitReq),
		withdrawReqChnl: make(chan withdrawReq),
		balanceReqChnl:  make(chan balanceReq),
	}
	go account.Serve()
	return account
}

func (a *Account) deposit(amount int) bool {
	a.balance += amount
	return true
}

func (a *Account) Deposit(amount int) bool {
	req := deopsitReq{
		amount: amount,
		resp:   make(chan bool),
	}
	a.depositReqChnl <- req
	return <-req.resp
}

func (a *Account) withdraw(amount int) bool {
	if a.balance < amount {
		return false
	}
	a.balance -= amount
	return true
}

func (a *Account) Withdraw(amount int) bool {
	req := withdrawReq{
		amount: amount,
		resp:   make(chan bool),
	}
	a.withdrawReqChnl <- req
	return <-req.resp
}

func (a *Account) getBalance() int {
	return a.balance
}

func (a *Account) GetBalance() int {
	req := balanceReq{
		resp: make(chan int),
	}
	a.balanceReqChnl <- req
	return <-req.resp
}

func (a *Account) Serve() {
	for {
		select {
		case req := <-a.depositReqChnl:
			req.resp <- a.deposit(req.amount)
		case req := <-a.withdrawReqChnl:
			req.resp <- a.withdraw(req.amount)
		case req := <-a.balanceReqChnl:
			req.resp <- a.getBalance()
		}
	}
}

func Transfer(from *Account, to *Account, amount int) bool {
	if from.Withdraw(amount) {
		return to.Deposit(amount)
	}
	fmt.Println("Insufficient balance")
	return false
}

func SimulateChannels() {
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
