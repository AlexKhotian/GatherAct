package serverhandling

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPHandler handles http requests
type HTTPHandler struct {
}

// InitHandler sets up routes
func (handler *HTTPHandler) InitHandler() {
	r := mux.NewRouter()
	r.HandleFunc("/changeactivity/{activityid}/{changetype}/{activitydiff}", handler.activityChange).Methods("PUT")
	r.HandleFunc("/getallactivities", handler.getAllActivities).Methods("GET")
	r.HandleFunc("/getactivitybudget/{userid}", handler.getActivityBudget).Methods("GET")
	// TODO: Change to TLS
	http.ListenAndServe(":80", r)
}

func (handler *HTTPHandler) activityChange(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["activityid"], vars["changetype"], vars["activitydiff"])
}

func (handler *HTTPHandler) getAllActivities(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to get all activities")
}

func (handler *HTTPHandler) getActivityBudget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["userid"])
}

