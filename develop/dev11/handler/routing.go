package handler

import (
	"net/http"
	"regexp"
)

type MyContext struct {
	slug string
	//slice of structs
	datamap map[int][]Calendar
}

func NewMyContext() *MyContext {
	calendar := make(map[int][]Calendar)
	ctx := MyContext{
		slug:    "",
		datamap: calendar,
	}
	return &ctx
}

func (s MyStore) Serve(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	switch {

	case match(p, "/create_event/+"):
		*s.Handl = s.logRequest(post(s.create_event))
	case match(p, "/update_event/+"):
		*s.Handl = s.logRequest(post(s.update_event))
	case match(p, "/delete_event"):
		*s.Handl = s.logRequest(post(s.delete_event))
	case match(p, "/events_for_day"):
		*s.Handl = s.logRequest(get(s.events_for_day))
	case match(p, "/events_for_week"):
		*s.Handl = s.logRequest(get(s.events_for_week))
	case match(p, "/events_for_month"):
		*s.Handl = s.logRequest(get(s.events_for_month))

	default:
		http.NotFound(w, r)
		return
	}
	x := *s.Handl
	x.ServeHTTP(w, r)
}

func match(path, pattern string, vars ...interface{}) bool {
	//регулярное выражение
	regex := regexp.MustCompile(pattern)
	//поиск списка подстрок
	matches := regex.FindStringSubmatch(path)
	if len(matches) <= 0 {
		return false
	}
	return true
}

func allowMethod(h http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			w.Header().Set("Allow", method)
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}

func get(h http.HandlerFunc) http.HandlerFunc {
	return allowMethod(h, "GET")
}

func post(h http.HandlerFunc) http.HandlerFunc {
	return allowMethod(h, "POST")
}
