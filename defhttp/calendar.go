package defhttp

import (
	"fmt"
	"net/http"
)

type apiWidget struct {
	slug string
	id   int
}

func (h apiWidget) create_event(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, h.slug)
	regex := mustCompileCached("user_id=(.*)&date=(.*)")
	//поиск списка подстрок
	matches := regex.FindStringSubmatch(h.slug)
	user_id := matches[1]
	date := matches[2]
	fmt.Fprintln(w, user_id)
	fmt.Fprintln(w, date)

}

func update_event(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "update_event\n")
}
func delete_event(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete_event\n")
}

func events_for_day(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_day\n")
}
func events_for_week(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_week\n")
}
func events_for_month(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_month\n")
}
