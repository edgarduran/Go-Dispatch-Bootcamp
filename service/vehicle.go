package service

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/edgarduran/Go-Dispatch-Bootcamp/models"
)

func GetAllVehicles() ([]models.Vehicle, error) {
	f, err := os.Open("vehicles.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var vehicleList []models.Vehicle
	for _, vehicle := range data {
		var rec models.Vehicle
		for j, field := range vehicle {
			if j == 0 {
				id, _ := strconv.ParseInt(field, 10, 64)
				rec.Id = id
			} else if j == 1 {
				rec.Make = field
			} else if j == 2 {
				rec.Model = field
			}
		}

		vehicleList = append(vehicleList, rec)
	}

	return vehicleList, nil
}

func GetVehicleById(id int) ([]models.Vehicle, error) {
	return nil, nil
}
