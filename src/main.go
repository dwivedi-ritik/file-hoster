package main

// Never ever panic in the routing the request

import (
	"log"

	"github.com/dwivedi-ritik/filehost-go/db"
	"github.com/gofiber/fiber/v2"
)

type UserPost struct {
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

func home(c *fiber.Ctx) error {

	allEntry, err := db.GetAllRow()

	if err != nil {
		return c.SendStatus(400)
	}

	return c.JSON(allEntry)
}

func homePost(c *fiber.Ctx) error {
	p := new(UserPost)
	if c.GetReqHeaders()["Content-Type"] != "application/json" {
		return c.SendStatus(400)
	}

	err := c.BodyParser(&p)
	if err != nil {
		return c.SendStatus(400)

	}

	return c.JSON(p)
}

func main() {
	app := fiber.New()
	app.Get("/api", home).Post("/api", homePost)
	log.Fatal(app.Listen(":3000"))

}
