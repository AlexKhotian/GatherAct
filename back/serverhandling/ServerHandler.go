package serverhandling

import "github.com/AlexKhotian/GatherAct/functions"
import "log"

// RunServer runs a main server routine
func RunServer() {
	log.Println("Starting server")

	activityHandler := functions.NewActivityHandler("root:admin@tcp(database:3306)/db")

	httpHandler := new(HTTPHandler)
	httpHandler.InitHandler(activityHandler)
}

