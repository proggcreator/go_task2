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
func (c *MyContext) create_event(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	name := r.URL.Query().Get("name")
	date := r.URL.Query().Get("date")
	c.datamap[id] = Calendar{
		data: date,
		name: name,
	}

	c.datamap["7"] = Calendar{
		data: date,
		name: name,
	}
	c.id = c.id + 2
	fmt.Fprintln(w, c)

}

func (c *MyContext) update_event(w http.ResponseWriter, r *http.Request) {
	c.id = c.id + 1
	fmt.Fprintln(w, "Мапа")
	if val, ok := c.datamap["3"]; ok {
		fmt.Fprintln(w, val)
	} else {
		fmt.Fprintln(w, "Ошибка")
	}
	fmt.Fprintln(w, c)

	//fmt.Fprintln(w, h.datamap["3"])
	//fmt.Fprintln(w, h.mapdata[1])

}
func (h *MyContext) delete_event(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete_evedelete_eventnt\n")
}

func (h *MyContext) events_for_day(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_day\n")
}
func (h *MyContext) events_for_week(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_week\n")
}
func (h *MyContext) events_for_month(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "events_for_month\n")
}
