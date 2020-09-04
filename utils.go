package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	formatSpecifier          = "%(\\d+\\$)?([-#+ 0,(\\<]*)?(\\d+)?(\\.\\d+)?([tT])?([a-zA-Z%])"
	maxExceededString string = "Error, exceeded the max limit as defined in config"
)

// Helper function for encoding things for returning from REST calls
func encode(i interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	err := enc.Encode(i)
	// Problems encoding
	if err != nil {
		loggingClient.Error("Error encoding the data: " + err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
}

// Printing function purely for debugging purposes
// Print the body of a request to the console
func printBody(r io.ReadCloser) {
	body, err := ioutil.ReadAll(r)
	bodyString := string(body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bodyString)
}

// Test if the service is working
func pingHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	_, err := w.Write([]byte("pong"))
	if err != nil {
		loggingClient.Error("Error writing pong: " + err.Error())
	}
}
