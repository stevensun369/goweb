package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"fmt"
)

func main() {
	engine := django.New("./views", ".html")

	
	
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	
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

	app.Get("/kard/:filename", func (c *fiber.Ctx) error {
		filename := c.Params("filename")

		return c.Render("kard", fiber.Map{
			"filename": filename,
		})
	})
		
}
