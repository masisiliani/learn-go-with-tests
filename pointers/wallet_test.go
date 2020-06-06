package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.WithDraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}
		err := wallet.WithDraw(Bitcoin(100))

		assertBalance(t, wallet, startBalance)
		assertError(t, err, ErrInsufficientFunds.Error())

	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, gotErr error, wantError string) {
	t.Helper()
	if gotErr == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if gotErr.Error() != wantError {
		t.Errorf("got %q, want %q", gotErr.Error(), wantError)
	}
}

func assertNoError(t *testing.T, gotErr error) {
	t.Helper()
	if gotErr != nil {
		t.Fatal("got an error but didn't want one")
	}
}