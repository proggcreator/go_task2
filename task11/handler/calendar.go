package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Calendar struct {
	Date string
	Name string
}

//from string to int, fill the structure
func fromString(id string, name string, date string) (int, Calendar, error) {
	//id to int
	newid, err := strconv.Atoi(id)
	if err != nil {
		logrus.Fatalf("Error 400")
		return 0, Calendar{}, err
	}
	return newid, Calendar{
		Date: date,
		Name: name,
	}, nil
}

//struct to JSON
func toJson(id int, name string, date string) (error, []byte) {

	jData, err := json.Marshal(Result{
		Name: name,
		Date: date,
		Id:   id,
	})
	if err != nil {
		logrus.Error("Marshal error: code 503")
		return err, nil
	}
	return nil, jData
}

//parse return param (querry string)
func parseQuerryRaw(r *http.Request) (string, string, string) {
	idstr := r.URL.Query().Get("user_id")
	name := r.URL.Query().Get("name")
	date := r.URL.Query().Get("date")
	return idstr, name, date

}
func (h *MyStore) create_event(w http.ResponseWriter, r *http.Request) {
	//parse
	idstr, name, date := parseQuerryRaw(r)

	//conver from string type
	myid, mycalend, err := fromString(idstr, name, date)
	if err != nil {
		h.Logger.Fatalf("Error 400")
		http.Error(w, "StatusBadRequest", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//save in memory
	h.Contx.datamap[myid] = append(h.Contx.datamap[myid], mycalend)

	//convert to json
	err, myJson := toJson(myid, name, date)
	if err != nil {
		resp := Response{http.StatusServiceUnavailable, "error convert to JSON"}
		respErr(w, resp)
		return

	}

	// response return JSON
	responceRes(w, myJson)

}

func (h *MyStore) update_event(w http.ResponseWriter, r *http.Request) {
	idstr, name, date := parseQuerryRaw(r)

	myid, mycalend, err := fromString(idstr, name, date)
	if err != nil {
		h.Logger.Fatalf("Error 400")
		http.Error(w, "StatusBadRequest", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//update date of event in memory
	flag := false
	for i, cur := range h.Contx.datamap[myid] {
		if cur.Name == mycalend.Name {
			h.Contx.datamap[myid][i].Date = mycalend.Date
			flag = true
		}
	}
	//if this date not exist for user_id
	if !flag {
		resp := Response{http.StatusServiceUnavailable, "this date is not exist"}
		respErr(w, resp)
		return
	}
	//convert to json
	err, myJson := toJson(myid, name, date)
	if err != nil {
		resp := Response{http.StatusServiceUnavailable, "error convert to JSON"}
		respErr(w, resp)
		return
	}

	// response return JSON
	responceRes(w, myJson)

}
func (h *MyStore) delete_event(w http.ResponseWriter, r *http.Request) {
	idstr, name, date := parseQuerryRaw(r)

	myid, mycalend, err := fromString(idstr, name, date)
	if err != nil {
		h.Logger.Fatalf("Error 400")
		http.Error(w, "StatusBadRequest", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mylist := h.Contx.datamap[myid]
	for i, cur := range h.Contx.datamap[myid] {
		if (cur.Date == mycalend.Date) && (cur.Name == mycalend.Name) {
			//delete from slice, order not important
			mylist[i] = mylist[len(mylist)-1]
			h.Contx.datamap[myid] = mylist[:len(mylist)-1]
		}
	}

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
