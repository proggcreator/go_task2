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

func CrmMatrix(lst []string, separator string) [][]string {
	var matrix [][]string
	for _, str := range lst {
		matrix = append(matrix, strings.Split(str, separator))
	}
	return matrix
}

func PrintOnlySep(matr [][]string) {
	for _, str := range matr {
		if len(str) > 1 { //если в срезе несколько элементов, есть разделитель
			fmt.Println(str)
		}
	}
}

func StrFflagToIntList(str string) (lst []int) { //список int для flagf
	lststring := strings.Split(str, ",")
	for _, s := range lststring {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		lst = append(lst, i)
	}
	return lst
}
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
	flagf := flag.String("f", "", "print only separator")
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
		//считывам построчно, добавляем в массив
		strlist = append(strlist, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	matrix := CrmMatrix(strlist, separator) //поддерржка флага d

	if len(*flagf) > 0 { //поддержка флага f (-f=0,1 text.txt)
		lst := StrFflagToIntList(*flagf)
		PrintFlagf(matrix, lst)
	}

	if *flags == true {
		PrintOnlySep(matrix) //поддрержка флага s
	}

}
