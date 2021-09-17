package main

import (
	"fmt"
)

type newMethodstruct struct {
}

func (a *newMethodstruct) visitForJpeg(s *pictureJpeg) {
	fmt.Println("Новый метод для формата JPEG")
}

func (a *newMethodstruct) visitForPng(s *pictureJpeg) {
	fmt.Println("Новый метод для формата PNG")
}
