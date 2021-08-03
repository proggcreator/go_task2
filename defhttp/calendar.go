package defhttp

import (
	"fmt"
	"net/http"
)

func create_event(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create_event\n")
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
