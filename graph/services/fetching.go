package services

import (
	"encoding/json"
	"net/http"
	"test-server/graph/model"
)

func FetchCharacterById(id string) (*model.Character, error) {
	resp, err := http.Get("https://swapi.dev/api/people/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var character model.Character
	if err := json.NewDecoder(resp.Body).Decode(&character); err != nil {
		return nil, err
	}
	character.ID = &id
	return &character, nil
}

func FetchPlanetById(id string) (*model.Planet, error) {
	resp, err := http.Get("https://swapi.dev/api/planets/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var planet model.Planet
	if err := json.NewDecoder(resp.Body).Decode(&planet); err != nil {
		return nil, err
	}

	return &planet, nil
}
