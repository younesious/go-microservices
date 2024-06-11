package main

import (
	"log"
	"net/http"

	"github.com/younesious/logger-service/log/data"
)

const authServiceURL = "http://authentication-service:8083/authenticate"

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		log.Println(err)
		_ = app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	_ = app.writeJSON(w, http.StatusAccepted, resp)
}
