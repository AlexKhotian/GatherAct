package serverhandling

import "github.com/AlexKhotian/GatherAct/functions"
import "log"

// RunServer runs a main server routine
func RunServer() {
	log.Println("Starting server")

	activityHandler := functions.NewActivityHandler("root:kukush@tcp(db:3306)/test_db")

	httpHandler := new(HTTPHandler)
	httpHandler.InitHandler(activityHandler)
}

