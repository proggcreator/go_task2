package main

import "fmt"

type app interface {
	send()
	get()
}

type writeApp struct {
}

func (w *writeApp) send() {
	fmt.Println("Отправить сообщение")
}

func (w *writeApp) get() {
	fmt.Println("Принять сообщение")
}
