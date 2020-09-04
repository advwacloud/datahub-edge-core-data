package main

import (
	"github.com/gorilla/mux"
)

func loadRestRoutes() *mux.Router {
	r := mux.NewRouter()
	b := r.PathPrefix("/api/v1").Subrouter()

	// EVENTS
	// /api/v1/event
	b.HandleFunc("/event", eventHandler).Methods("GET", "PUT", "POST")
	//e := b.PathPrefix("/event").Subrouter()
	//e.HandleFunc("/source/{sourceId}/{limit:[0-9]+}", getDataBySourceHandler).Methods("GET")

	// Ping Resource
	// /api/v1/ping
	b.HandleFunc("/ping", pingHandler)

	return r
}
