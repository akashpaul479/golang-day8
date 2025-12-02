package day8

import "testing"

func TestDeposit(t *testing.T) {
	acc := &account{Balance: 100}
	acc.Deposit(50)

	if acc.Balance != 150 {
		t.Errorf("Expected 150, got %d", acc.Balance)
	}

}
func TestWithdrawl(t *testing.T) {
	acc := &account{Balance: 100}
	if !acc.Withdrawl(50) {
		t.Errorf("Withdraw should suceed")
	}
	if acc.Balance != 50 {
		t.Errorf("Expected balance 50 , got %d", acc.Balance)
	}

}
