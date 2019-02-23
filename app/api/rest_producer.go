package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amebalabs/nocreep/app/model"
)

func (a *API) recordEvent(w http.ResponseWriter, r *http.Request) {
	event := model.Event{}

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		a.respondWithError(w, http.StatusBadRequest, "can't decode event payload")
		return
	}
	event.CreatedAt = time.Now()
	if err := a.db.RecordEvent(event.DeviceID, event); err != nil {
		fmt.Println(err)
		a.respondWithError(w, http.StatusBadRequest, "can't save event")
		return
	}
}
