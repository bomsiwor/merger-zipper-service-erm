package main

import (
	"fmt"
	"mymodule/handler"
	"mymodule/utils"
	"os"

	// "mymodule/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Start new server
	r := mux.NewRouter()

	// Use global middleware
	r.Use(utils.CheckToken)

	// Registering router
	handler.NewRouter(r)

	// Create server config
	// Use port from env
	port := os.Getenv("SERVICE_PORT")
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:" + port,
	}

	// Start server
	fmt.Println("Starting server on port " + port)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
