package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func Mycurrtime() {

	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Println(err) // вывод ошибки
		os.Exit(-1)      //код выхода

	}
	fmt.Printf("Точное время: %s / Текущее время: %s", t, time.Now())

}
