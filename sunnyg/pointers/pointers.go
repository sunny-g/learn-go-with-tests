package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringifiable interface {
	Stringify() string
}

var ErrInsufficientFunds = errors.New("insufficient funds")

func (b Bitcoin) Stringify() (string) {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Balance() (Bitcoin) {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) (*Wallet, error) {
	w.balance += amount
	return w, nil
}

func (w *Wallet) Withdraw(amount Bitcoin) (*Wallet, error) {
	if amount > w.balance {
		return w, ErrInsufficientFunds
	}

	w.balance -= amount
	return w, nil
}
