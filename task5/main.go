package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Print_String_After(strlist []string, pos int, N int) {

	for i := pos; i <= pos+N; i++ {
		fmt.Println(strlist[i])

	}
}
func FindStr(strlist []string, lookFor string) int {
	for i, curstr := range strlist {
		if strings.Contains(curstr, lookFor) == true {
			return i
		}
	}
	return 0
}
func main() {

	flaga := flag.Int("a", 0, "number of string")
	//flagn := flag.Bool("n", false, "int sort")
	flag.Parse()
	lookFor := flag.Arg(1)
	fmt.Println(lookFor)
	var strlist []string
	var in io.Reader

	//чтение из файла или stdin
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()
		in = f
	} else {
		in = os.Stdin
	}
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		//считывам построчно, добавляем в массив
		strlist = append(strlist, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//проверка на количество строк
	if *flaga > len(strlist) {
		fmt.Println("Error too many N")
		return
	}
	pos_str := FindStr(strlist, lookFor) //поиск строки с совпадением
	if *flaga != 0 {

		Print_String_After(strlist, pos_str, *flaga)
	} else {
		fmt.Println(strlist[pos_str])
	}

}
