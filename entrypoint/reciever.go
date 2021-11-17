package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// reciever is a function for reading events from http entrypoint.
func reciever(events chan<- *Event, entrypoint string) {

	// definition of a general handler.
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		// validation of request.
		switch {
		case r.Method != "POST":
			fmt.Fprint(rw, "Only POST method is supported")
			return
		case r.Header.Get("Content-Type") != "application/json":
			fmt.Fprint(rw, "Only JSON format of events is supported")
			return
		}
		defer r.Body.Close()

		// validation of schema.
		var event Event
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			rw.WriteHeader(400)
		}

		events <- &event
		rw.WriteHeader(200)
	})

	// start entrypoint.
	if err := http.ListenAndServe(entrypoint, nil); err != nil {
		panic(err) //todo: processing error is here
	}
}