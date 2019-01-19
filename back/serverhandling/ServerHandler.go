package serverhandling

func RunServer() {
	httpHandler := new(HTTPHandler)
	httpHandler.InitHandler()
}

