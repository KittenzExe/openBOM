package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("disclaimer.md")
	})

	app.Get("/v1/location/search", func(c *fiber.Ctx) error {
		query := c.Query("q")
		return locationSearch(c, query)
	})

	app.Get("/v1/weather/daily/direct", func(c *fiber.Ctx) error {
		x_query := c.Query("x")
		y_query := c.Query("y")
		timezone_query := c.Query("t")
		return weatherDailyDirect(c, x_query, y_query, timezone_query)
	})

	app.Get("/v1/weather/daily/query", func(c *fiber.Ctx) error {
		query := c.Query("q")
		return weatherDailyQuery(c, query)
	})

	app.Get("/v1/weather/hourly/direct", func(c *fiber.Ctx) error {
		x_query := c.Query("x")
		y_query := c.Query("y")
		timezone_query := c.Query("t")
		return weatherHourlyDirect(c, x_query, y_query, timezone_query)
	})

	app.Get("/v1/weather/hourly/query", func(c *fiber.Ctx) error {
		query := c.Query("q")
		return weatherHourlyQuery(c, query)
	})

	app.Get("/v1/weather/astro/direct", func(c *fiber.Ctx) error {
		longitude_query := c.Query("lon")
		latitude_query := c.Query("lat")
		return weatherAstroDirect(c, longitude_query, latitude_query)
	})

	app.Get("/v1/weather/astro/query", func(c *fiber.Ctx) error {
		query := c.Query("q")
		return weatherAstroQuery(c, query)
	})

	log.Fatal(app.Listen(":3838"))
}
