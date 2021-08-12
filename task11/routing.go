package defhttp

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

type MyContext struct {
	slug    string
	id      int
	datamap map[string]Calendar
}
type serverHandler struct {
	srv *Server
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

func (sh serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := sh.srv.httpServer.Handler
	mycache := NewMyContext()

	//rep := MyRep{}
	p := r.URL.Path
	switch {
	case match(p, "/create_event/+"):
		h = post(mycache.create_event)
		fmt.Fprintln(w, mycache)
	case match(p, "/update_event/+"):
		h = post(mycache.update_event)
		fmt.Fprintln(w, mycache)
	case match(p, "/delete_event"):
		h = post(mycache.delete_event)
	case match(p, "/events_for_day"):
		h = get(mycache.events_for_day)
	case match(p, "/events_for_week"):
		h = get(mycache.events_for_week)
	case match(p, "/events_for_month"):
		h = get(mycache.events_for_month)

	default:
		http.NotFound(w, r)
		return
	}

	h.ServeHTTP(w, r)
}

// match reports whether path matches ^regex$, and if it matches,
// assigns any capture groups to the *string or *int vars.
func match(path, pattern string, vars ...interface{}) bool {
	//регулярное выражение
	regex := mustCompileCached(pattern)
	//поиск списка подстрок
	matches := regex.FindStringSubmatch(path)
	if len(matches) <= 0 {
		return false
	}
	for i, match := range matches[1:] {
		switch p := vars[i].(type) {
		case *string:
			*p = match
		default:
			panic("vars must be *string or *int")
		}
	}
	return true
}

var (
	regexen = make(map[string]*regexp.Regexp)
	relock  sync.Mutex
)

func mustCompileCached(pattern string) *regexp.Regexp {
	relock.Lock()
	defer relock.Unlock()

	regex := regexen[pattern]
	if regex == nil {
		regex = regexp.MustCompile("^" + pattern + "$")
		regexen[pattern] = regex
	}
	return regex
}

// allowMethod takes a HandlerFunc and wraps it in a handler that only
// responds if the request method is the given method, otherwise it
// responds with HTTP 405 Method Not Allowed.
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

// get takes a HandlerFunc and wraps it to only allow the GET method
func get(h http.HandlerFunc) http.HandlerFunc {
	return allowMethod(h, "GET")
}

// post takes a HandlerFunc and wraps it to only allow the POST method
func post(h http.HandlerFunc) http.HandlerFunc {
	return allowMethod(h, "POST")
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home\n")
}
