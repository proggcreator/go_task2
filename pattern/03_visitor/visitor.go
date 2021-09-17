package main

type visitor interface {
	visitForJpeg(*pictureJpeg)
	visitForPng(*picturePng)
}
