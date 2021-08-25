package main

// go run main.go -n  text.txt the
//реализована работа с первым совпадением
import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Print_String_After(strlist []string, pos int, N int) error {
	//проверка на количество строк для флага a
	if (N + pos) <= len(strlist) {
		for i := pos; i <= pos+N; i++ {
			fmt.Println(strlist[i])
		}
		return nil
	}
	err := errors.New("Error out of len")
	return err

}
func Print_String_Before(strlist []string, pos int, N int) error {
	//проверка на количество строк для флага b
	if (pos - N) >= 0 {
		for i := pos - N; i <= pos; i++ {
			fmt.Println(strlist[i])
		}
		return nil
	}
	err := errors.New("Error out of len")
	return err

}
func Print_String_Around(strlist []string, pos int, N int) error {
	if (pos-N >= 0) && (pos+N <= len(strlist)) {
		for i := pos - N; i <= pos+N; i++ {
			fmt.Println(strlist[i])
		}
		return nil
	}
	err := errors.New("Out range")
	return err
}
func FindStr(strlist []string, lookFor string) (int, error) {
	for i, curstr := range strlist {
		if strings.Contains(curstr, lookFor) == true {
			//вернуть индекс строки
			return i, nil
		}
	}
	err := errors.New("No match found")
	return 0, err
}
func FindFullStr(strlist []string, lookFor string) (int, error) {
	for i, curstr := range strlist {
		if curstr == lookFor {
			//вернуть индекс строки
			return i, nil
		}
	}
	err := errors.New("No match found ")
	return 0, err
}
func main() {

	flaga := flag.Int("A", 0, "after number of string")
	flagb := flag.Int("B", 0, "before number of string")
	flagc := flag.Int("c", 0, "arount number of string")
	flagn := flag.Bool("n", false, "index of string")
	flagf := flag.Bool("F", false, "index of full string")
	flagi := flag.Bool("i", false, "ignore case")

	flag.Parse()
	//что искать
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
		//если флаг i приводим  все к нижнему регистру
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
	//если флаг i приводим все к нижнему регистру
	if *flagi {
		lookFor = strings.ToLower(lookFor)
	}
	//реализация флага f точное совпадение со строкой
	if *flagf {
		pos_str, _ = FindFullStr(strlist, lookFor)

	} else {
		pos_str, _ = FindStr(strlist, lookFor) //сначала поиск строки с совпадением

	}
	//реализация флага n печать номера строки
	if *flagn {
		fmt.Println("Number of string: ", pos_str)

	}

	//реализация флага a
	if *flaga != 0 {
		Print_String_After(strlist, pos_str, *flaga)
		return
	}

	//реализация флага b
	if *flagb != 0 {
		Print_String_Before(strlist, pos_str, *flaga)
		return
	}
	//реализация флага C
	if *flagc != 0 {
		Print_String_Around(strlist, pos_str, *flaga)
		return
	}

}
