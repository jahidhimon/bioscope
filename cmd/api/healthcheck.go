package main

import (
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
	err := app.writeJSON(w, http.StatusOK, envelope{"Server Info": res}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
