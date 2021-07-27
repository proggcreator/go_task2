//считается что все колонки заполеннны
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

func crmatrix(lst []string) [][]string {
	var matrix [][]string
	for _, str := range lst {
		matrix = append(matrix, strings.Fields(str))
	}
	return matrix
}

func StrToInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func main() {

	flagk := flag.Int("k", 0, "number of column")
	flagn := flag.Bool("n", false, "int sort")
	flagr := flag.Bool("r", false, "reverse sort")
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
	matrix := crmatrix(strlist)

	if *flagn == true {
		//реализация флага r
		fmt.Println("d")
	} else if *flagr == true {
		sort.Slice(matrix, func(i, j int) bool { return matrix[i][*flagk] > matrix[j][*flagk] })
		//реализация флага n
	} else if *flagn == true {
		sort.Slice(matrix, func(i, j int) bool { return StrToInt(matrix[i][*flagk]) < StrToInt(matrix[j][*flagk]) })
	} else {
		//реализация флага k или по умолчанию
		sort.Slice(matrix, func(i, j int) bool {

			return StrToInt(matrix[i][*flagk]) < StrToInt(matrix[j][*flagk])
		})
	}
	//реализация флага u
	if *flagu == true {
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

}
