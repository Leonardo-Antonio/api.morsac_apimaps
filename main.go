package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Location struct {
	Data []Data `json:"data"`
}

type Data struct {
	DisplayAddress string  `json:"display_address"`
	DisplayRegion  string  `json:"display_region"`
	Log            float64 `json:"lat"`
	Lat            float64 `json:"lon"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/mor-sac/:place", func(c *fiber.Ctx) error {
		place := c.Params("place")

		url := "https://api.mymappi.com/v2/geocoding/direct?apikey=0876a809-8921-4bb7-bd2d-f2069c9951e0&q=" + place
		res, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var location Location
		if err := json.Unmarshal(body, &location); err != nil {
			log.Fatalln(err)
		}
		return c.Status(http.StatusOK).JSON(location)
	})

	app.Listen(":" + os.Getenv("PORT"))
}
