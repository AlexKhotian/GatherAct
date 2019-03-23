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
	r.HandleFunc("/change_activity/{team_id}/{activity_id}/{activity_diff}", handler.activityChange).Methods("PUT")
	r.HandleFunc("/add_activity/{team_id}/{activity_name}", handler.addActivity).Methods("PUT")
	r.HandleFunc("/get_allactivities_for_team/{team_id}", handler.getAllActivitiesForTeam).Methods("GET")
	r.HandleFunc("/getactivitybudget/{userid}", handler.getActivityBudget).Methods("GET")
	// TODO: Change to TLS
	http.ListenAndServe(":80", r)
}

func (handler *HTTPHandler) activityChange(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	activityDiff, err := strconv.Atoi(vars["activity_diff"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	activityId, err := strconv.Atoi(vars["activity_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	teamId, err := strconv.Atoi(vars["team_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("activityChange: success")
	handler.activityHandler.ChangeActivity(uint32(teamId), int32(activityDiff), uint32(activityId))
	w.WriteHeader(http.StatusOK)
}

func (handler *HTTPHandler) addActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamId, err := strconv.Atoi(vars["team_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("addActivity: success")
	handler.activityHandler.AddActivity(uint32(teamId), vars["activity_name"])
	w.WriteHeader(http.StatusCreated)
}

func (handler *HTTPHandler) getAllActivitiesForTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId, err := strconv.Atoi(vars["team_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := handler.activityHandler.GetActivitiesForTeam(uint32(teamId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("getAllActivitiesForTeam: success")
	for _, activity := range result {
		log.Println(activity)
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *HTTPHandler) getActivityBudget(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["userid"])
}

