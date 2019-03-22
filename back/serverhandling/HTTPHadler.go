package serverhandling

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/AlexKhotian/GatherAct/functions"
)

// HTTPHandler handles http requests
type HTTPHandler struct {
	activityHandler *functions.ActivityHandler
}

// InitHandler sets up routes
func (handler *HTTPHandler) InitHandler(activityHandler *functions.ActivityHandler) {
	handler.activityHandler = activityHandler
	r := mux.NewRouter()
	r.HandleFunc("/changeactivity/{activityid}/{changetype}/{activitydiff}", handler.activityChange).Methods("PUT")
	r.HandleFunc("/getallactivities", handler.getAllActivities).Methods("GET")
	r.HandleFunc("/getactivitybudget/{userid}", handler.getActivityBudget).Methods("GET")
	// TODO: Change to TLS
	http.ListenAndServe(":80", r)
}

func (handler *HTTPHandler) activityChange(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	handler.activityHandler.ChangeActivity(1, strconv.Atoi(vars["activitydiff"]), strconv.Atoi(vars["activityid"]))
}

func (handler *HTTPHandler) getAllActivities(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to get all activities")
}

func (handler *HTTPHandler) getActivityBudget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["userid"])
}

