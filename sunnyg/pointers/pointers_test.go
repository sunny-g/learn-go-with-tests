package pointers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBitcoin(t *testing.T) {

}

func TestWallet(t *testing.T) {

	t.Run("should deposit", func(t *testing.T) {
		t.Helper()
		wallet := Wallet{balance: Bitcoin(0)}
		expected := Bitcoin(10)

		_, err := wallet.Deposit(Bitcoin(10))
		actual := wallet.Balance()
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should withdraw", func(t *testing.T) {
		t.Helper()
		wallet := Wallet{balance: Bitcoin(20)}
		expected := Bitcoin(10)

		_, err := wallet.Withdraw(Bitcoin(10))
		actual := wallet.Balance()
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should not withdraw insufficient funds", func(t *testing.T) {
		t.Helper()
		wallet := Wallet{balance: Bitcoin(0)}
		expected := Bitcoin(0)

		_, err := wallet.Withdraw(Bitcoin(10))
		actual := wallet.Balance()
		assert.EqualError(t, err, ErrInsufficientFunds.Error())
		assert.Equal(t, expected, actual)
	})

}
