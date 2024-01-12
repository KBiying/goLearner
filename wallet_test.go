package main

import (
	"errors"
	"fmt"
	"testing"
)

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct{
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin){
	fmt.Println("address of balance in test is", &w.balance)
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin{
	return w.balance
}

func (w *Wallet)Withdraw(amount Bitcoin) error {
	if amount > w.balance{
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

func(b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
} 

func TestWallet(t *testing.T){

	
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		fmt.Println("address of balance in test is", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(10))
		
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance:Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance:startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t,err,InsufficientFundsError)
	})

}

func assertBalance (t *testing.T, wallet Wallet, want Bitcoin){
	got := wallet.Balance()
	if got!=want{
		t.Errorf("got %s want %s ", got, want)
	}		
}
// func assertError (t *testing.T, got error, want error){
// 	if got == nil {
// 		t.Fatal("didn't get an error but wanted one")
// 	}
// 	if got.Error() != want.Error(){
// 		t.Errorf("got %s want %s ", got, want)
// 	}
// }