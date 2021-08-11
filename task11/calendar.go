package defhttp

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
func (h *Myapi) create_event(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	name := r.URL.Query().Get("name")
	date := r.URL.Query().Get("date")
	h.datamap[id] = Calendar{
		data: date,
		name: name,
	}

	h.datamap["4"] = Calendar{
		data: date,
		name: name,
	}
	fmt.Fprintln(w, h.datamap)
	fmt.Fprintln(w, id)
	fmt.Fprintln(w, name)
	fmt.Fprintln(w, date)

}

func (h *Myapi) update_event(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Мапа")
	if val, ok := h.datamap["3"]; ok {
		fmt.Fprintln(w, val)
	} else {
		fmt.Fprintln(w, "Ошибка")
	}
	fmt.Fprintln(w, h.datamap)

	//fmt.Fprintln(w, h.datamap["3"])
	//fmt.Fprintln(w, h.mapdata[1])

}
func (h *Myapi) delete_event(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete_evedelete_eventnt\n")
}

func (h *Myapi) events_for_day(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_day\n")
}
func (h *Myapi) events_for_week(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_week\n")
}
func (h *Myapi) events_for_month(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_month\n")
}