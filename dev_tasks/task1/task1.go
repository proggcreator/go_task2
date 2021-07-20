package task1

import (
	"fmt"

	"github.com/beevik/ntp"
)

func mytime() {

	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	fmt.Println(time)
	fmt.Println(err)
}
