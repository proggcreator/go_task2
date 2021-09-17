package main

import "fmt"

type confirm struct {
}

func newConfirm() *confirm {
	return &confirm{}
}

func (a *confirm) checkConfirm() {
	fmt.Println("Заказ успешно обработан")

}
