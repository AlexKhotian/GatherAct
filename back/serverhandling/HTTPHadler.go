package serverhandling

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
}

func (handler *HTTPHandler) InitHandler() {
	r := mux.NewRouter()
	r.HandleFunc("/changeactivity/{activityid}/{changetype}/{activitydiff}", handler.HandleActivityChange).Methods("PUT")
	// TODO: Change to TLS
	http.ListenAndServe(":80", r)
}

func (handler *HTTPHandler) HandleActivityChange(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["activityid"], vars["changetype"], vars["activitydiff"])
}

