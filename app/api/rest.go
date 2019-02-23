package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amebalabs/nocreep/app/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//API service
type API struct {
	store.DataInterface
	db *store.BoltDB
}

func requestTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is: %s\n", time.Now())
}

func (a *API) requestBolt(w http.ResponseWriter, r *http.Request) {
	devices, err := a.GetDevices()
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
func (a *API) Run(connection *store.BoltDB) {
	fmt.Println("Starting server on port :3000")
	a.db = connection
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/time", requestTime)
	r.Route("/say", func(r chi.Router) {
		r.Get("/{name}", requestSay)
		r.Get("/", requestSay)
	})
	r.Route("/devices", func(r chi.Router) {
		r.Get("/{deviceid}", a.requestBolt)
		r.Get("/", a.requestBolt)
	})

	r.Route("/events", func(r chi.Router) {
		r.Get("/{id}", a.getEvents)
	})

	r.Post("/event", a.recordEvent)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}

func (a *API) respondWithError(w http.ResponseWriter, code int, message string) {
	a.respondWithJSON(w, code, map[string]string{"error": message})
}

func (a *API) respondWithSuccess(w http.ResponseWriter, code int, message string) {
	a.respondWithJSON(w, code, map[string]string{"success": message})
}

func (a *API) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data := map[string]interface{}{}
	data["results"] = payload
	out := map[string]interface{}{}
	out["status"] = "ok"
	out["message"] = "ok"
	out["data"] = data
	response, _ := json.Marshal(out)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
