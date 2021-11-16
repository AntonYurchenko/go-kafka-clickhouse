package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func reciever(events chan<- Event, entrypoint string) {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		switch {
		case r.Method != "POST":
			fmt.Fprint(rw, "Only POST method is supported")
			return
		case r.Header.Get("Content-Type") != "application/json":
			fmt.Fprint(rw, "Only JSON format of events is supported")
			return
		}
		defer r.Body.Close()

		var event Event
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			rw.WriteHeader(400)
		}

		events <- event
		rw.WriteHeader(200)
	})

	if err := http.ListenAndServe(entrypoint, nil); err != nil {
		panic(err)
	}
}
