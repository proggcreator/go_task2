package main

// go run main.go -n  text.txt the
//реализована работа с первым совпадением
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
func Print_String_Before(strlist []string, pos int, N int) {
	for i := pos - N; i <= pos; i++ {
		fmt.Println(strlist[i])

	}
}
func FindStr(strlist []string, lookFor string) int {
	for i, curstr := range strlist {
		if strings.Contains(curstr, lookFor) == true {
			//вернуть индекс строки
			return i
		}
	}
	log.Fatalf("No match found ")
	return 0
}
func FindFullStr(strlist []string, lookFor string) int {
	for i, curstr := range strlist {
		if curstr == lookFor {
			//вернуть индекс строки
			return i
		}
	}
	log.Fatalf("No match found ")
	return 0
}
func main() {

	flaga := flag.Int("A", 0, "after number of string")
	flagb := flag.Int("B", 0, "before number of string")
	flagn := flag.Bool("n", false, "index of string")
	flagf := flag.Bool("F", false, "index of full string")
	flagi := flag.Bool("i", false, "ignore case")

	flag.Parse()
	lookFor := flag.Arg(1)
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
	//считывам построчно, добавляем в массив
	for scanner.Scan() {
		//если флаг приводим  все к нижнему регистру
		if *flagi {
			strlist = append(strlist, strings.ToLower(scanner.Text()))
		} else {
			strlist = append(strlist, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	pos_str := 0
	//реализация флага f
	if *flagf {
		pos_str = FindFullStr(strlist, lookFor)

	} else {
		pos_str = FindStr(strlist, lookFor) //поиск строки с совпадением

	}

	if *flagn {
		fmt.Println("Номер строки: ", pos_str)

	}
	//проверка на количество строк
	if (*flaga + pos_str) > len(strlist) {
		fmt.Println("Error out of len  ")
		return
	}
	if *flaga != 0 {
		Print_String_After(strlist, pos_str, *flaga)
		return
	}

	//проверка на количество строк
	if (pos_str - *flagb) >= 0 {
		Print_String_Before(strlist, pos_str, *flaga)
		return
	}

}
