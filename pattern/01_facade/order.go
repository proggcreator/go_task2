package main

import "fmt"

type order struct {
}

func newOrder() *order {
	return &order{}
}

func (w *order) checkItems() error {
	fmt.Println("Проверка наличия товаров")
	return nil
}
