package api

import (
	"log"
	"net/http"

	"github.com/amebalabs/nocreep/app/model"

	"github.com/go-chi/chi"
)

func (a *API) getEvents(w http.ResponseWriter, r *http.Request) {
	id := model.DeviceID(chi.URLParam(r, "id"))

	events, err := a.GetDeviceEvents(id)
	if err != nil {
		log.Println(err)
		a.respondWithError(w, http.StatusNotFound, "No events for this device found")
		return
	}
	a.respondWithJSON(w, http.StatusOK, events)
}
