package handler

import "net/http"

type Result struct {
	Id   int
	Date string
	Name string
}

func responceRes(w http.ResponseWriter, myJson []byte) {
	// response return JSON
	w.Header().Set("Content-Type", "calendar/json")
	w.Write(myJson)

}
