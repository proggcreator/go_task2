package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type MyStore struct {
	handl  *http.Handler
	contx  *MyContext
	logger *logrus.Logger
}

func NewStore() *MyStore {
	return &MyStore{
		handl:  new(http.Handler),
		contx:  NewMyContext(),
		logger: logrus.New(),
	}
}
