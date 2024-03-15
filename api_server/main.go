package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"compumed/auth"
	"compumed/device"
	"compumed/logging"
	"compumed/mqtt"
)

func heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	logging.Log("Starting API server")
	router := mux.NewRouter()
	router.HandleFunc("/heartbeat", heartbeat)
	authRouter := router.PathPrefix("/auth").Subrouter()
	auth.CreateAuthRoutes(authRouter)
	deviceRouter := router.PathPrefix("/device").Subrouter()
	device.CreateRoutes(deviceRouter)

	logging.Log("STARTING MQTT")
	mqtt.Run()

	http.Handle("/", router)
	// Start and listen to requests
	http.ListenAndServe(":8080", router)
}
