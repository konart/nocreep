package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amebalabs/nocreep/app/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var db *store.BoltDB

func requestTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is: %s\n", time.Now())
}

func requestBolt(w http.ResponseWriter, r *http.Request) {
	devices, err := db.GetDevice()
	if err != nil {
		fmt.Fprintf(w, "Something went wrong")
	}
	fmt.Fprintf(w, "Devices Count: %d\n", len(devices))
	fmt.Fprintf(w, "Devices list: %s\n", devices)
}

func requestSay(w http.ResponseWriter, r *http.Request) {
	val := chi.URLParam(r, "name")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!\n", val)
	} else {
		fmt.Fprintf(w, "Hello ... you.\n")
	}
}

// Run server
func Run(connection *store.BoltDB) {
	fmt.Println("Starting server on port :3000")
	db = connection
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/time", requestTime)
	r.Route("/say", func(r chi.Router) {
		r.Get("/{name}", requestSay)
		r.Get("/", requestSay)
	})
	r.Route("/devices", func(r chi.Router) {
		r.Get("/{deviceid}", requestBolt)
		r.Get("/", requestBolt)
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
