package main

import (
	"fmt"

	"github.com/interactics/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)

	fmt.Println(account)
}

// 펑션을 만들어서 생성자를 만들자
// d에러를 위한 딕셔너러니 생성 가능
// Go 가 자동으로 호출하는 매서드 string
