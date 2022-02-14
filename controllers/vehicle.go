package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edgarduran/Go-Dispatch-Bootcamp/usecase"
)

type vehicleController struct {
}

func New() vehicleController {
	return vehicleController{}
}

func (c vehicleController) GetAllVehicles(w http.ResponseWriter, rtr *http.Request) {
	vehicles, err := usecase.GetAllVehicles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting all vehicles: %v\n", err)

		return
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
