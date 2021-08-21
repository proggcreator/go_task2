package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Calendar struct {
	Date time.Time
	Name string
}

//from string to int, date, fill the structure
func fromString(id string, name string, date string) (int, Calendar, error) {
	//id to int
	newid, err := strconv.Atoi(id)
	if err != nil {
		logrus.Fatalf("Error 400")
		return 0, Calendar{}, err
	}
	//date string to date

	myDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		logrus.Fatalf("Error 400")
		return 0, Calendar{}, err
	}
	return newid, Calendar{
		Date: myDate,
		Name: name,
	}, nil
}

//struct to JSON
func toJson(id int, calend Calendar) (error, []byte) {

	jData, err := json.Marshal(Result{
		Name: calend.Name,
		Date: calend.Date,
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
	err, myJson := toJson(myid, mycalend)
	if err != nil {
		resp := Response{http.StatusServiceUnavailable, "error convert to JSON"}
		respErr(w, resp)
		return

	}
	fmt.Fprint(w, h.Contx.datamap)
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
	err, myJson := toJson(myid, mycalend)
	if err != nil {
		resp := Response{http.StatusServiceUnavailable, "error convert to JSON"}
		respErr(w, resp)
		return
	}
	fmt.Fprint(w, h.Contx.datamap)
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
	//marshal JSON
	jData, err := json.Marshal(h.Contx.datamap[myid])
	if err != nil {
		resp := Response{http.StatusServiceUnavailable, "marshall error"}
		respErr(w, resp)
		return
	}
	fmt.Fprint(w, h.Contx.datamap)
	responceRes(w, jData)

}

func (h *MyStore) events_for_day(w http.ResponseWriter, r *http.Request) {

	idstr, name, date := parseQuerryRaw(r)
	myid, mycalend, err := fromString(idstr, name, date)
	if err != nil {
		h.Logger.Fatalf("Error 400")
		http.Error(w, "StatusBadRequest", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//list for events for day
	listDayEvents := []string{}

	for _, cur := range h.Contx.datamap[myid] {
		day := cur.Date.Day()
		if day == mycalend.Date.Day() {
			listDayEvents = append(listDayEvents, cur.Name)
		}
	}
	//marshal JSON
	jData, err := json.Marshal(listDayEvents)
	if err != nil {
		resp := Response{http.StatusServiceUnavailable, "marshall error"}
		respErr(w, resp)
		return
	}
	responceRes(w, jData)
}
func (h *MyStore) events_for_week(w http.ResponseWriter, r *http.Request) {
	idstr, name, date := parseQuerryRaw(r)
	myid, mycalend, err := fromString(idstr, name, date)
	if err != nil {
		h.Logger.Fatalf("Error 400")
		http.Error(w, "StatusBadRequest", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//list for events for week
	listWeekEvents := []string{}
	//necessary week number
	_, myweek := mycalend.Date.ISOWeek()
	for _, cur := range h.Contx.datamap[myid] {
		//current week number
		_, curweek := cur.Date.ISOWeek()
		if myweek == curweek {
			listWeekEvents = append(listWeekEvents, cur.Name)
		}
	}
	//marshal JSON
	jData, err := json.Marshal(listWeekEvents)
	if err != nil {
		resp := Response{http.StatusServiceUnavailable, "marshall error"}
		respErr(w, resp)
		return
	}
	//return list events for week
	responceRes(w, jData)
}

func (h *MyStore) events_for_month(w http.ResponseWriter, r *http.Request) {

	idstr, name, date := parseQuerryRaw(r)
	myid, mycalend, err := fromString(idstr, name, date)
	if err != nil {
		h.Logger.Fatalf("Error 400")
		http.Error(w, "StatusBadRequest", 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//list for events for month
	listMonthEvents := []string{}
	//necessary month number
	mymonth := mycalend.Date.Month()
	for _, cur := range h.Contx.datamap[myid] {
		//current month number
		curmonth := cur.Date.Month()
		if mymonth == curmonth {
			listMonthEvents = append(listMonthEvents, cur.Name)
		}
	}
	//marshal JSON
	jData, err := json.Marshal(listMonthEvents)
	if err != nil {
		resp := Response{http.StatusServiceUnavailable, "marshall error"}
		respErr(w, resp)
		return
	}
	//return list events for month
	responceRes(w, jData)
}
