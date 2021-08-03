package defhttp

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	//var slug string
	//var id int

	p := r.URL.Path
	switch {
	case match(p, "/"):
		h = get(home)
	case match(p, "/create_event"):
		h = post(create_event)
	case match(p, "/update_event"):
		h = post(update_event)
	case match(p, "/delete_event"):
		h = post(delete_event)
	case match(p, "/events_for_day"):
		h = get(events_for_day)
	case match(p, "/events_for_week"):
		h = get(events_for_week)
	case match(p, "/events_for_month"):
		h = get(events_for_month)

	default:
		http.NotFound(w, r)
		return
	}
	h.ServeHTTP(w, r)
}

// match reports whether path matches ^regex$, and if it matches,
// assigns any capture groups to the *string or *int vars.
func match(path, pattern string, vars ...interface{}) bool {
	regex := mustCompileCached(pattern)
	matches := regex.FindStringSubmatch(path)
	if len(matches) <= 0 {
		return false
	}
	for i, match := range matches[1:] {
		switch p := vars[i].(type) {
		case *string:
			*p = match
		case *int:
			n, err := strconv.Atoi(match)
			if err != nil {
				return false
			}
			*p = n
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
