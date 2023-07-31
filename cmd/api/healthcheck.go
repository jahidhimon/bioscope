package main

import (
	"encoding/json"
	"net/http"
)

type Healthcheck struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	res := &Healthcheck{
		Status:      "available",
		Environment: app.config.env,
		Version:     version,
	}
	js, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
