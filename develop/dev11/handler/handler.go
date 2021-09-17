package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type MyStore struct {
	Handl  *http.Handler
	Contx  *MyContext
	Logger *logrus.Logger
}

func NewStore() *MyStore {
	return &MyStore{
		Handl:  new(http.Handler),
		Contx:  NewMyContext(),
		Logger: logrus.New(),
	}
}
