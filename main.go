package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
	"log"
)

func main() {

	app := fiber.New(fiber.Config{})
	
	app.Static("/media", "./media")

	app.Get("/", func (c*fiber.Ctx) error {
		return c.Render("home", fiber.Map{})
	})

	app.Post("/", func (c *fiber.Ctx) error {
		file, err := c.FormFile("file")

		if (err != nil) {
				return c.Redirect("/")
		}

		c.SaveFile(file,  fmt.Sprintf("./media/%s", file.Filename))

		return c.Redirect("/kard/" + file.Filename)
	})

	app.Get("/kard/:filename", func (c *fiber.Ctx) error {

		filename := c.Params("filename")

		return c.Render("kard", fiber.Map{
			"filename": filename,
		})

	})

	log.Fatal(app.Listen(":9990"))
}
