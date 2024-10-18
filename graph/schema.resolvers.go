package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"strings"
	"test-server/graph/model"
	"test-server/graph/services"
)

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

// Planet returns PlanetResolver implementation.
func (r *Resolver) Planet() PlanetResolver { return &planetResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type planetResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
