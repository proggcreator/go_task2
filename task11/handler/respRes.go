package handler

import (
	"net/http"
	"time"
)

type Result struct {
	Id   int
	Date time.Time
	Name string
}

func responceRes(w http.ResponseWriter, myJson []byte) {
	// response return JSON
	w.Header().Set("Content-Type", "calendar/json")
	w.Write(myJson)

}
