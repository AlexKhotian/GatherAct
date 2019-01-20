package serverhandling

// RunServer runs a main server routine
func RunServer() {
	httpHandler := new(HTTPHandler)
	httpHandler.InitHandler()
}

