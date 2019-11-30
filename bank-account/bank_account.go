package account
import (
	"sync"
)

type Account struct {
	money int64
	available bool
	mu sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{money: initialDeposit, available: true}
}

func (a *Account) Close()(payout int64, ok bool) {
	a.mu.Lock()
	if a.available {
		payout = a.money
		a.available = false
		ok = true
	} else {
		payout = 0
		ok = false
	}
	a.mu.Unlock()
	return
}

func (a *Account) Balance()(balance int64, ok bool) {
	a.mu.Lock()
	if a.available {
		balance = a.money 
	} else {
		balance = 0
	}
	ok = a.available
	a.mu.Unlock()
	return 
}

func (a *Account) Deposit(amount int64)(newBalance int64, ok bool){
	a.mu.Lock()
	ok = a.available
	if a.available {
		a.money += amount
		if a.money < 0 {
			ok = false
			a.money -= amount
			newBalance = 0
		} else {
			newBalance = a.money
		}
	} else {
		newBalance = 0
	}
	a.mu.Unlock()
	return 
}