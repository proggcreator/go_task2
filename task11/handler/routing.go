package handler

import (
	"fmt"
	"net/http"
	"regexp"
)

type MyContext struct {
	slug    string
	id      int
	datamap map[string]Calendar
}

func NewMyContext() *MyContext {
	calendar := make(map[string]Calendar)
	ctx := MyContext{
		id:      0,
		slug:    "",
		datamap: calendar,
	}
	return &ctx
}

func (h Handler) Serve(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	switch {
	case match(p, "/create_event/+"):
		*h.handl = post(h.create_event)
		fmt.Fprintln(w, "Сощдатьь")
	case match(p, "/update_event/+"):
		*h.handl = post(h.update_event)
		fmt.Fprintln(w, "Обновить")
	case match(p, "/delete_event"):
		*h.handl = post(h.delete_event)
	case match(p, "/events_for_day"):
		*h.handl = get(h.events_for_day)
	case match(p, "/events_for_week"):
		*h.handl = get(h.events_for_week)
	case match(p, "/events_for_month"):
		*h.handl = get(h.events_for_month)

	default:
		http.NotFound(w, r)
		return
	}
	x := *h.handl
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
