package account

import "sync"

type Account struct {
	mu     sync.Mutex
	amount int64
	closed bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{amount: amount}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.amount, !a.closed
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	newAmount := a.amount + amount
	if newAmount < 0 {
		return a.amount, false
	}

	a.amount = newAmount
	return a.amount, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	prevAmount := a.amount
	a.amount = 0

	success := true
	if a.closed {
		success = false
	}
	a.closed = true

	return prevAmount, success
}
