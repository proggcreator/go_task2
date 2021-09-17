package main

import (
	"fmt"
	"sort"
)

func myanagrams(strs []string) map[string]string {

	primaryDict := make(map[string][]string)
	m := make(map[string]string)

	for _, word := range strs {
		byteWords := []byte(word)
		//сортировка символов в слове
		sort.SliceStable(byteWords, func(i, j int) bool {
			return byteWords[i] < byteWords[j]
		})
		//в строку
		stringyWord := string(byteWords)
		//словарь с ключами - отсортированным набором букв
		primaryDict[stringyWord] = append(primaryDict[stringyWord], word)
	}
	//карта с множествами где больше 1 анаграммы
	for _, v := range primaryDict {
		if len(v) > 1 {
			for i := 1; i < len(v); i++ {
				m[v[0]] = v[i]
			}
		}

	}
	return m

}

func main() {
	masstr := []string{"пятак", "тапок", "тяпка", "дом", "улица"}
	fmt.Println(myanagrams(masstr))

}
