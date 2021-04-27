package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

var errNoMoney = errors.New("can't withdraw")

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
} // 고에서 receiver라고 한다.
// 리시버의 규칙은 구조체의 첫글자를 따서 해야한다.
// 메서드이다.
// 포인터 리시버.

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
} //에러 핸들링 전용 함수가 존재하지 않음.
//

// 이거 근데 함수 내부 변수 소멸 안되나..?
// 메서드랑 펑션이랑 다른가?

// ChangeOwner of account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account\nHas: ", a.Balance())
} //Sprint는 스트링 형태를 의미한다.
