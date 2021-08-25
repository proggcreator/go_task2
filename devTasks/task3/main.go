//считается что все колонки заполеннны //go run main.go -k=0 text.txt
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

//создание матрицы
func crmatrix(lst []string) [][]string {
	var matrix [][]string
	for _, str := range lst {
		matrix = append(matrix, strings.Fields(str))
	}
	return matrix
}

//из строки в int
func StrToInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}
func mySort(strlist []string, flagk int, flagn bool, flagr bool, flagu bool) [][]string {
	//реализация флага r
	matrix := crmatrix(strlist)
	if flagr == true {
		sort.Slice(matrix, func(i, j int) bool { return matrix[i][flagk] > matrix[j][flagk] })
	}
	//реализация флага n
	if flagn == true {
		sort.Slice(matrix, func(i, j int) bool { return StrToInt(matrix[i][flagk]) < StrToInt(matrix[j][flagk]) })
	}
	//реализация флага k или по умолчанию
	sort.Slice(matrix, func(i, j int) bool {
		return matrix[i][flagk] < matrix[j][flagk]
	})

	//реализация флага u
	if flagu == true {
		for i, str := range matrix {
			if i != len(matrix)-1 { //если не последняя строка
				if reflect.DeepEqual(matrix[i], matrix[i+1]) == false { //рекурсивное сравнение
					fmt.Println(str)
				}
			}
		}
	} else {
		for _, str := range matrix {
			fmt.Println(str)
		}

	}
	return matrix
}
func main() {
	//флаг колонка для сортировки
	flagk := flag.Int("k", 0, "number of column")
	//сортировать по числовому значению
	flagn := flag.Bool("n", false, "int sort")
	//сортировать в обратном порядке
	flagr := flag.Bool("r", false, "reverse sort")
	//не выводить повторяющиеся строки
	flagu := flag.Bool("u", false, "a bool")
	flag.Parse()

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

	mySort(strlist, *flagk, *flagn, *flagr, *flagu)

}
