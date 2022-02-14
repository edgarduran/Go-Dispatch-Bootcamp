package service

import (
	"github.com/edgarduran/Go-Dispatch-Bootcamp/models"
)

var vehicles = []models.Vehicle{
	{Id: 1, Make: "ford", Model: "mustang"},
	{Id: 2, Make: "kia", Model: "sorento"},
}

func GetAllVehicles() ([]models.Vehicle, error) {
	// TODO: replace with csv data
	return vehicles, nil
}
