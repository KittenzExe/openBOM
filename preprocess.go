package main

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Gridcell struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Gridcells struct {
	Forecast Gridcell `json:"forecast"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type LocationCandidate struct {
	Coordinate Coordinate `json:"coordinate"`
	Gridcells  Gridcells  `json:"gridcells"`
	Timezone   string     `json:"timezone"`
}

type LocationSearchResponse struct {
	Candidates []LocationCandidate `json:"candidates"`
}

func fetchLocationCandidate(query string) (LocationCandidate, error) {
	resp, err := proxyBomRequest("https://api.bom.gov.au/apikey/v1/locations/places/autocomplete?name=" + query + "&limit=10&website-sort=true&include-states=true&include-districts=true")
	if err != nil {
		return LocationCandidate{}, err
	}

	var data LocationSearchResponse
	if err := json.Unmarshal(resp, &data); err != nil || len(data.Candidates) == 0 {
		return LocationCandidate{}, errors.New("no location found")
	}

	return data.Candidates[0], nil
}

func preprocessQueryForGridcells(c *fiber.Ctx, query string) (fiber.Map, error) {
	candidate, err := fetchLocationCandidate(query)
	if err != nil {
		return nil, err
	}

	return fiber.Map{
		"forecast": candidate.Gridcells.Forecast,
		"timezone": candidate.Timezone,
	}, nil
}

func preprocessQueryForCoordinates(c *fiber.Ctx, query string) (fiber.Map, error) {
	candidate, err := fetchLocationCandidate(query)
	if err != nil {
		return nil, err
	}

	return fiber.Map{
		"latitude":  candidate.Coordinate.Latitude,
		"longitude": candidate.Coordinate.Longitude,
	}, nil
}
