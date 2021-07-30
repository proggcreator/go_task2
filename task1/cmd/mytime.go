package cmd

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func Mycurrtime() {

	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {

		fmt.Println(err) // вывод ошибки
		os.Exit(-1)      //код выхода

	}
	fmt.Println("Точное время:")
	fmt.Println(time) ///вывод точного времени

}
