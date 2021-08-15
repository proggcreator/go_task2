package handler

import (
	"fmt"
	"net/http"
)

type Calendar struct {
	data string
	name string
}

/*func parse_string(str string) []string {
	regex := mustCompileCached("user_id=(.*)&date=(.*)&name=(.*)")
	//формирование списка подстрок
	matches := regex.FindStringSubmatch(str)
	return matches
}*/
func (h *MyStore) create_event(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Этоо")
	id := r.URL.Query().Get("user_id")
	name := r.URL.Query().Get("name")
	date := r.URL.Query().Get("date")
	h.contx.datamap[id] = Calendar{
		data: date,
		name: name,
	}

	h.contx.datamap["70"] = Calendar{
		data: date,
		name: name,
	}
	h.contx.id = h.contx.id + 2
	fmt.Fprintln(w, h.contx)
}

func (h *MyStore) update_event(w http.ResponseWriter, r *http.Request) {
	h.contx.id = h.contx.id + 1
	fmt.Fprintln(w, "Мапа")
	if val, ok := h.contx.datamap["3"]; ok {
		fmt.Fprintln(w, val)
	} else {
		fmt.Fprintln(w, "Ошибка")
	}
	fmt.Fprintln(w, h.contx)

	//fmt.Fprintln(w, h.datamap["3"])
	//fmt.Fprintln(w, h.mapdata[1])

}
func (h *MyStore) delete_event(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete_evedelete_eventnt\n")
}

func (h *MyStore) events_for_day(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_day\n")
}
func (h *MyStore) events_for_week(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_week\n")
}
func (h *MyStore) events_for_month(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_month\n")
}
