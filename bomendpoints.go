package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func locationSearch(c *fiber.Ctx, query string) error {
	body, err := proxyBomRequest("https://api.bom.gov.au/apikey/v1/locations/places/autocomplete?name=" + query + "&limit=10&website-sort=true&include-states=true&include-districts=true")
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", "application/json")
	return c.Send(body)
}

func weatherDailyDirect(c *fiber.Ctx, x_query string, y_query string, timezone string) error {
	body, err := proxyBomRequest("https://api.bom.gov.au/apikey/v1/forecasts/daily/" + x_query + "/" + y_query + "?timezone=" + timezone)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", "application/json")
	return c.Send(body)
}

func weatherDailyQuery(c *fiber.Ctx, query string) error {
	resp, err := preprocessQueryForGridcells(c, query)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	forecast := resp["forecast"].(Gridcell)
	timezone := resp["timezone"].(string)

	return weatherDailyDirect(c, fmt.Sprintf("%d", forecast.X), fmt.Sprintf("%d", forecast.Y), timezone)
}

func weatherHourlyDirect(c *fiber.Ctx, x_query string, y_query string, timezone string) error {
	body, err := proxyBomRequest("https://api.bom.gov.au/apikey/v1/forecasts/1hourly/" + x_query + "/" + y_query + "?timezone=" + timezone)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", "application/json")
	return c.Send(body)
}

func weatherHourlyQuery(c *fiber.Ctx, query string) error {
	resp, err := preprocessQueryForGridcells(c, query)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	forecast := resp["forecast"].(Gridcell)
	timezone := resp["timezone"].(string)

	return weatherHourlyDirect(c, fmt.Sprintf("%d", forecast.X), fmt.Sprintf("%d", forecast.Y), timezone)
}

func weatherAstroDirect(c *fiber.Ctx, longitude string, latitude string) error {
	body, err := proxyBomRequest("https://api.bom.gov.au/apikey/v1/forecasts/astro/" + longitude + "/" + latitude)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", "application/json")
	return c.Send(body)
}

func weatherAstroQuery(c *fiber.Ctx, query string) error {
	resp, err := preprocessQueryForCoordinates(c, query)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	longitude := resp["longitude"].(float64)
	latitude := resp["latitude"].(float64)

	return weatherAstroDirect(c, fmt.Sprintf("%v", longitude), fmt.Sprintf("%v", latitude))
}
