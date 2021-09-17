//фасад
package main

import (
	"fmt"
	"log"
)

type orderFacade struct {
	order    *order
	reqMoney *reqMoney
	confirm  *confirm
}

func newOrderFacade() *orderFacade {
	walletFacacde := &orderFacade{
		order:    newOrder(),
		reqMoney: newreqMoney(),
		confirm:  newConfirm(),
	}
	return walletFacacde
}

func (w *orderFacade) addOrder() error {
	fmt.Println("Формирование нового заказа")
	err := w.order.checkItems()
	if err != nil {
		return err
	}
	err = w.reqMoney.checMooney()
	if err != nil {
		return err
	}
	w.confirm.checkConfirm()
	return nil
}

func main() {
	//создаем новый интерфейс реализующий фасад
	walletFacade := newOrderFacade()
	fmt.Println()
	err := walletFacade.addOrder()
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}

/*Формирование нового заказа
Проверка наличия товаров
Проверка сретств
Заказ успешно обработан*/
