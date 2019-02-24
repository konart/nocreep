package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/amebalabs/nocreep/app/model"

	"github.com/amebalabs/nocreep/app/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//API service
type API struct {
	Version string
	store.Interface
	httpServer *http.Server
}

func requestTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is: %s\n", time.Now())
}

func (a *API) requestBolt(w http.ResponseWriter, r *http.Request) {
	devices, err := a.GetDevices(model.User{})
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
func (a *API) Run(port int) {
	fmt.Printf("Starting server on port :%d", port)
	a.fakeData()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/time", requestTime)
	r.Route("/say", func(r chi.Router) {
		r.Get("/{name}", requestSay)
		r.Get("/", requestSay)
	})
	r.Route("/devices", func(r chi.Router) {
		r.Get("/{id}", a.getDevice)
		r.Get("/", a.getDevices)
	})

	r.Route("/events", func(r chi.Router) {
		r.Get("/{id}", a.getEvents)
	})

	r.Post("/event", a.recordEvent)

	a.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      120 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	err := a.httpServer.ListenAndServe()
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

// Shutdown rest http server
func (a *API) Shutdown() {
	log.Print("[WARN] shutdown rest server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if a.httpServer != nil {
		if err := a.httpServer.Shutdown(ctx); err != nil {
			log.Printf("[DEBUG] http shutdown error, %s", err)
		}
		log.Print("[DEBUG] shutdown http server completed")
	}
}

func (a *API) fakeData() {
	for i := 1; i < 10; i++ {
		a.AddDevice(
			model.Device{
				ID: model.DeviceID(strconv.Itoa(i)),
				OS: "iOS 11.2 " + strconv.Itoa(i)})
	}
}
