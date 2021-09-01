//паттерн фасад
//Простой (но урезанный) интерфейс к сложной системе объектов
//Кроме того, что Фасад позволяет снизить общую сложность программы,
//он также помогает вынести код, зависимый от внешней системы в единственное место.
package main

import (
	"fmt"
	"log"
)

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
}

func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	return nil
}

func main() {
	fmt.Println()
	//создаем новый интерфейс реализующий фасад
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()

}

/*Starting create account
Account created

Starting add money to wallet
Account Verified
SecurityCode Verified
Wallet balance added successfully*/
