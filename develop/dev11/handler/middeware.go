package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s MyStore) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logg := s.Logger.WithFields(logrus.Fields{
			"remode_addr": r.RemoteAddr,
		})
		logg.Infof("Started %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})

}
