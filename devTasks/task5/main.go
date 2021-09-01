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

func Find_String_After(strlist []string, pos int, N int) ([]string, error) {
	changedList := []string{}
	//проверка на количество строк для флага a
	if (N + pos) <= len(strlist) {
		for i := pos; i <= pos+N; i++ {
			changedList = append(changedList, strlist[i])
		}
		return changedList, nil
	}
	err := errors.New("Error out of len")
	return changedList, err

}
func Find_String_Before(strlist []string, pos int, N int) ([]string, error) {
	changedList := []string{}
	//проверка на количество строк для флага b
	if (pos - N) >= 0 {
		for i := pos - N; i <= pos; i++ {
			changedList = append(changedList, strlist[i])
		}
		return changedList, nil
	}
	err := errors.New("Error out of len")
	return changedList, err

}
func Find_String_Around(strlist []string, pos int, N int) ([]string, error) {
	changedList := []string{}
	if (pos-N >= 0) && (pos+N <= len(strlist)) {
		for i := pos - N; i <= pos+N; i++ {
			changedList = append(changedList, strlist[i])
		}
		return changedList, nil
	}
	err := errors.New("Out range")
	return changedList, err
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
func Find_Count(strlist []string, lookFor string) int {
	count := 0
	for _, curstr := range strlist {
		if strings.Contains(curstr, lookFor) == true {
			//+1
			count++
		}
	}
	return count
}
func print_res(strlist []string) {
	for x := range strlist {
		fmt.Println(x)
	}
}

func Remove_Match(strlist []string, lookFor string) []string {
	//соединяем в текст строки
	newtext := strings.Join(strlist, " ")
	//разбиваем по словам
	strlist = strings.Split(newtext, " ")
	newlist := []string{}
	for _, x := range strlist {
		if x != lookFor {
			newlist = append(newlist, x)
		}
	}
	return newlist
}
func main() {

	flaga := flag.Int("A", 0, "after number of string")
	flagb := flag.Int("B", 0, "before number of string")
	flagC := flag.Int("C", 0, "arount number of string")
	flagv := flag.Bool("v", false, "print strings without match ")
	flagc := flag.Bool("c", false, "count string match ")
	flagn := flag.Bool("n", false, "index of string")
	flagf := flag.Bool("F", false, "index of full string")
	flagi := flag.Bool("i", false, "ignore case")

	flag.Parse()
	//что искать
	lookFor := flag.Arg(1)
	var strlist []string
	var err error
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
	//реализация флага c печать количества совпадений
	if *flagc {
		count := Find_Count(strlist, lookFor)
		fmt.Println("Count match: ", count)
		return

	}
	//реализация флага v удаление совппадений
	if *flagv {
		strlist := Remove_Match(strlist, lookFor)
		print_res(strlist)
		return
	}

	//реализация флага a
	if *flaga != 0 {
		strlist, err = Find_String_After(strlist, pos_str, *flaga)
		if err != nil {
			fmt.Println("error Print_String_After", err)
			return
		}
	}

	//реализация флага b
	if *flagb != 0 {
		strlist, err = Find_String_Before(strlist, pos_str, *flaga)
		if err != nil {
			fmt.Println("error Print_String_Before", err)
			return
		}
	}
	//реализация флага C
	if *flagC != 0 {
		strlist, err = Find_String_Around(strlist, pos_str, *flaga)
		if err != nil {
			fmt.Println("error Print_String_Before", err)
			return
		}
	}
	//печать результата
	print_res(strlist)

}
