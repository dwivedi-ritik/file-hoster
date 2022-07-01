package main

// Never ever panic in the routing the request

import (
	"log"

	"github.com/dwivedi-ritik/filehost-go/db"
	"github.com/dwivedi-ritik/filehost-go/models"
	"github.com/gofiber/fiber/v2"
)

func home(c *fiber.Ctx) error {

	allEntry, err := db.GetAllRow()

	if err != nil {
		return c.SendStatus(400)
	}

	return c.JSON(allEntry)
}

func homePost(c *fiber.Ctx) error {
	// P itself a pointer
	p := new(models.Product)
	if c.GetReqHeaders()["Content-Type"] != "application/json" {
		return c.SendStatus(400)
	}

	err := c.BodyParser(&p)
	if err != nil {
		return c.SendStatus(400)

	}
	db.AddRow(p)
	return c.JSON(p)
}

func getFile(c *fiber.Ctx) error {
	return c.Download("./temp.txt")
}

func main() {
	// Make Migrations
	db.MakeMigration()

	app := fiber.New()
	app.Get("/api", home).Post("/api", homePost)
	app.Get("/api/file", getFile)
	log.Fatal(app.Listen(":3000"))

}
