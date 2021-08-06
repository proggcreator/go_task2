package defhttp

import (
	"fmt"
	"net/http"
)

//строка параметров
type apiWidget struct {
	slug string
}

func (h apiWidget) create_event(w http.ResponseWriter, r *http.Request) {
	//параметры передаются фиксированно
	regex := mustCompileCached("user_id=(.*)&date=(.*)&name=(.*)")
	//формирование списка подстрок
	matches := regex.FindStringSubmatch(h.slug)
	user_id := matches[1]
	date := matches[2]
	name := matches[3]
	fmt.Fprintln(w, user_id)
	fmt.Fprintln(w, date)
	fmt.Fprintln(w, name)

}

func (h apiWidget) update_event(w http.ResponseWriter, r *http.Request) {
	regex := mustCompileCached("user_id=(.*)&date=(.*)&name=(.*)")
	//формирование списка подстрок
	matches := regex.FindStringSubmatch(h.slug)
	user_id := matches[1]
	date := matches[2]
	name := matches[3]
	fmt.Fprintln(w, user_id)
	fmt.Fprintln(w, date)
	fmt.Fprintln(w, name)

}
func delete_event(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete_evedelete_eventnt\n")
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
