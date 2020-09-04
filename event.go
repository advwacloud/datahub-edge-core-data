package main

import (
	"datahub-edge-core/clients"
	"github.com/advwacloud/datahub-edge-domain-models"
	"encoding/json"
	"net/http"
)

// Handler for the event API
// Status code 404 - event not found
// Status code 413 - number of events exceeds limit
// Status code 503 - unanticipated issues
// api/v1/even
func eventHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.Method {
	// Get all events
	case "GET":
		datas, err := dbc.Datas()
		if err != nil {
			loggingClient.Error(err.Error())
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		// Check max limit
		if len(datas) > configuration.Readmaxlimit {
			http.Error(w, maxExceededString, http.StatusRequestEntityTooLarge)
			loggingClient.Error(maxExceededString)
			return
		}

		encode(datas, w)
		break
	// Post a new event
	case "POST":
		var e models.Event
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&e)

		// Problem Decoding Event
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			loggingClient.Error("Error decoding data: " + err.Error())
			return
		}
		loggingClient.Info("Posting Data: " + e.String())

		// Add the readings to the database
		if configuration.Persistdata {
			for i, reading := range e.Datas {

				reading.Device = e.Device // Update the device for the reading

				// Add the reading
				id, err := dbc.AddData(reading)
				if err != nil {
					http.Error(w, err.Error(), http.StatusServiceUnavailable)
					loggingClient.Error(err.Error())
					return
				}

				e.Datas[i].Id = id // Set the ID for referencing later
			}

			// Add the event to the database
			// id, err := dbc.AddEvent(&e)
			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusServiceUnavailable)
			// 	loggingClient.Error(err.Error())
			// 	return
			// }

			w.WriteHeader(http.StatusOK)
		} else {
			encode("unsaved", w)
		}

		//putEventOnQueue(e)                                 //
		//updateDeviceLastReportedConnected(e.Device)        // update last reported connected (device)
		//updateDeviceServiceLastReportedConnected(e.Device) // update last reported connected (device service)

		break
	// Do not update the readings
	case "PUT":
		var from models.Event
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&from)

		// Problem decoding event
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			loggingClient.Error("Error decoding the event: " + err.Error())
			return
		}

		// Check if the event exists
		to, err := dbc.DataById(from.ID.Hex())
		if err != nil {
			if err == clients.ErrNotFound {
				http.Error(w, "Event not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
			}
			loggingClient.Error(err.Error())
			return
		}

		loggingClient.Info("Updating event: " + from.ID.Hex())

		// Update the fields
		if from.Device != "" {
			to.Device = from.Device
		}

		// Update
		if err = dbc.UpdateData(to); err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			loggingClient.Error(err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("true"))
		//encode(true, w)
	}
}
