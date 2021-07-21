package task1

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func mycurrtime() {

	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {

		fmt.Println(err) // вывод ошибки
		os.Exit(-1)      //код выхода

	}
	fmt.Println(time) ///вывод точного времени

}
