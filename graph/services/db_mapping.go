package services

import (
	"fmt"
	"test-server/dbmodel"
	"test-server/graph/model"
)

func DBVehicleToGraphQLVehicle(dbVehicle *dbmodel.Vehicle) *model.Vehicle {
	return &model.Vehicle{
		ID:                   fmt.Sprint(dbVehicle.ID), // Convert uint to string
		CargoCapacity:        &dbVehicle.CargoCapacity,
		Consumables:          &dbVehicle.Consumables,
		CostInCredits:        &dbVehicle.CostInCredits,
		Crew:                 &dbVehicle.Crew,
		Length:               &dbVehicle.Length,
		Manufacturer:         &dbVehicle.Manufacturer,
		MaxAtmospheringSpeed: &dbVehicle.MaxAtmospheringSpeed,
		Model:                &dbVehicle.Model,
		Name:                 &dbVehicle.Name,
		Passengers:           &dbVehicle.Passengers,
		VehicleClass:         &dbVehicle.VehicleClass,
	}
}
