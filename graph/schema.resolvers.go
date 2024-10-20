package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
	"strings"
	"test-server/dbmodel"
	"test-server/graph/model"
	"test-server/graph/services"
)

// CreateVehicle is the resolver for the createVehicle field.
func (r *mutationResolver) CreateVehicle(ctx context.Context, input model.NewVehicle) (*model.Vehicle, error) {
	Addvehicle := dbmodel.Vehicle{
		CargoCapacity:        *input.CargoCapacity,
		Consumables:          *input.Consumables,
		CostInCredits:        *input.CostInCredits,
		Crew:                 *input.Crew,
		Length:               *input.Length,
		Manufacturer:         *input.Manufacturer,
		MaxAtmospheringSpeed: *input.MaxAtmospheringSpeed,
		Model:                *input.Model,
		Name:                 *input.Name,
		Passengers:           *input.Passengers,
		VehicleClass:         *input.VehicleClass,
	}

	if err := r.Database.Create(&Addvehicle).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return services.DBVehicleToGraphQLVehicle(&Addvehicle), nil
}

// Residents is the resolver for the residents field.
func (r *planetResolver) Residents(ctx context.Context, obj *model.Planet) ([]*model.Character, error) {
	var residents []*model.Character

	for _, residentURL := range obj.ResidentURLs {
		parts := strings.Split(*residentURL, "/")
		id := parts[len(parts)-2]
		character, _ := services.FetchCharacterById(id)
		residents = append(residents, character)
	}
	return residents, nil
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	var character, err = services.FetchCharacterById(id)
	if err != nil {
		return nil, err
	}
	return character, nil
}

// Planet is the resolver for the planet field.
func (r *queryResolver) Planet(ctx context.Context, id string) (*model.Planet, error) {
	var planet, err = services.FetchPlanetById(id)
	if err != nil {
		return nil, err
	}
	return planet, nil
}

// Vehicle is the resolver for the vehicle field.
func (r *queryResolver) Vehicle(ctx context.Context, id string) (*model.Vehicle, error) {
	vehicle := model.Vehicle{}

	if err := r.Database.Find(&vehicle, id).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &vehicle, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Planet returns PlanetResolver implementation.
func (r *Resolver) Planet() PlanetResolver { return &planetResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type planetResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
