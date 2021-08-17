package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Response struct {
	StatusCode int
	Msg        string
}

func respErr(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "calendar/json")
	date, err := json.Marshal(r)

	http.Error(w, "StatusServiceUnavailable", 503)
	w.WriteHeader(http.StatusServiceUnavailable)

	if err != nil {
		logrus.Error("Marshal error: code 503")
		return
	}
	logrus.Error("StatusServiceUnavailable: code 503")
	w.Write(date)
}
