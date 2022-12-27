package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"rest-api-tutorial/internal/user"
	"rest-api-tutorial/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create router")
	router := httprouter.New()

	logger.Info("register user new handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening port 1234")
	logger.Fatal(server.Serve(listener))
}
