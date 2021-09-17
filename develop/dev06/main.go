package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//создаем разбиваем по строкам и столбцам
func CrmMatrix(lst []string, separator string) [][]string {
	var matrix [][]string
	for _, str := range lst {
		matrix = append(matrix, strings.Split(str, separator))
	}
	return matrix
}

//вывести только строки с разделителем
func FindOnlySep(matr [][]string) [][]string {
	newmatr := [][]string{}
	for _, str := range matr {
		//если в срезе несколько элементов - есть разделитель
		if len(str) > 1 {
			newmatr = append(newmatr, str)
		}
	}
	return newmatr
}

//вывод колонок
func StrFflagToIntList(str string) (lst []int) {
	//разбиваем строку по "," - номера нужных столбцов
	lststring := strings.Split(str, ",")
	for _, s := range lststring {
		//преобразовать в int нужные колонки
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		//список int колонок
		lst = append(lst, i)
	}
	return lst
}

//вывод нужных колонок
func PrintFlagf(matr [][]string, lst []int) {
	for _, str := range matr {
		for i := 0; i < len(lst); i++ {
			fmt.Print(str[lst[i]] + "   ")
		}
		fmt.Println()
	}
}

func main() {
	flags := flag.Bool("s", false, "print only separator")
	flagd := flag.String("d", "   ", "init separator")
	flagf := flag.String("f", "", "chose columns")
	flag.Parse()
	//lookFor := flag.Arg(1)
	var strlist []string
	var in io.Reader

	//чтение из файла или stdin
	separator := *flagd //флаг d
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
		//считывам построчно, добавляем все в массив
		strlist = append(strlist, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//поддерржка флага d кастомный разделитель
	matrix := CrmMatrix(strlist, separator)

	//поддержка флага f выбор колонок (-f=0,1 text.txt)
	if len(*flagf) > 0 {
		lst := StrFflagToIntList(*flagf)
		PrintFlagf(matrix, lst)
	}
	//поддрержка флага s вывод строки у которых есть разделитель
	if *flags == true {
		//новая матрица
		matrix = FindOnlySep(matrix)
		for x := range matrix {
			fmt.Print(x)
		}
	}

}
