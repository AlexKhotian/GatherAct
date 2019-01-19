package main

import "github.com/AlexKhotian/GatherAct/back/serverhandling"

func main() {
	server := new(serverhandling.ServerHandler)
	server.RunServer()
}

