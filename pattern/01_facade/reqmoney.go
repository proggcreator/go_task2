package main

import "fmt"

type reqMoney struct {
}

func newreqMoney() *reqMoney {
	return &reqMoney{}
}

func (s *reqMoney) checMooney() error {
	fmt.Println("Проверка сретств")
	return nil
}
