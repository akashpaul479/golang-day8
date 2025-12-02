package day8

type account struct {
	Balance int
}

func (a *account) Deposit(amount int) {
	a.Balance += amount
}
func (a *account) Withdrawl(amount int) bool {
	if a.Balance < amount {
		return false
	}
	a.Balance -= amount
	return true
}
