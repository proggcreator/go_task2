package main

//библиотека, разные структуры (форматы файлов) реализуют один интерфейс
//библиотека обработки изображений, новый метод обработки
import "fmt"

type picture interface {
	getType() string
	accept(visitor)
}

func main() {
	pictureJpeg := &pictureJpeg{}
	picturePng := &picturePng{}

	newMethodstruct := &newMethodstruct{}

	pictureJpeg.accept(newMethodstruct)
	picturePng.accept(newMethodstruct)
	fmt.Println()
}
