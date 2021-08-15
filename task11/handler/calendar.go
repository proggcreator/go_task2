package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
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

func fromString(id string, name string, date string) (int, Calendar, error) {
	//id to int
	newid, err := strconv.Atoi(id)
	if err != nil {
		logrus.Fatalf("Error 400")
		return 0, Calendar{}, err
	}
	//datetime
	/*
		t, err := time.Parse(time.RFC3339, date)
		if err != nil {
			logrus.Fatalf("Error 400")
			return 0, Calendar{}, err
		}
	*/
	return newid, Calendar{
		data: date,
		name: name,
	}, nil
}

func toJson(id string, name string, date string) (error, []byte) {
	jData, err := json.Marshal("Result: " + id + " " + name + " " + date)
	if err != nil {
		return err, nil
	}
	return nil, jData
}
func (h *MyStore) create_event(w http.ResponseWriter, r *http.Request) {
	//parse querry raw
	idstr := r.URL.Query().Get("user_id")
	name := r.URL.Query().Get("name")
	date := r.URL.Query().Get("date")

	//convert to json
	err, myJson := toJson(idstr, name, date)
	if err != nil {
		//error 503
		logrus.Fatalf("Error 503")
		w.WriteHeader(http.StatusServiceUnavailable)
		return

	}

	//conver from string type
	myid, mycalend, err := fromString(idstr, name, date)
	if err != nil {
		logrus.Fatalf("Error 400")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//save in memory
	h.contx.datamap[myid] = mycalend

	//responce
	//return JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(myJson)

}

func (h *MyStore) update_event(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Мапа")
	if val, ok := h.contx.datamap[5]; ok {
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
