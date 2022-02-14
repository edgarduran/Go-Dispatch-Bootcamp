package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/edgarduran/Go-Dispatch-Bootcamp/models"
)

func GetAllVehicles(w http.ResponseWriter, rtr *http.Request) {
	// TODO: replace with csv data
	vehicles := []models.Vehicle{
		{Id: 1, Make: "ford", Model: "mustang"},
		{Id: 2, Make: "kia", Model: "sorento"},
	}

	jsonData, err := json.Marshal(vehicles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
