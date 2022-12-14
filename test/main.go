package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	pr := fmt.Println
	app := fiber.New(fiber.Config{
		// Prefork: true,
		// CaseSensitive: true,
		// StrictRouting: true,
		// ServerHeader:  "Fiber",
		AppName: "Testing Fiber 1.2.3",
		GETOnly: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world")
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		return c.SendString("params is : " + c.Params("id", "apple is a fruit"))
	})

	app.Get("/members/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
		}
		return c.SendString("Where is smith?")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("POST method")
	})

	if !fiber.IsChild() {
		pr("this is parent")
	} else {
		pr("i am child process")
	}

	app.Get("/api/posts", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
			"posts":   "this is posts",
		})
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "custom message is here!")
	})

	app.Static("/page", "./public")

	//match any request
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	// match request starting with /api
	app.Use("/abc", func(c *fiber.Ctx) error {
		return c.Next()
	})

	// mount
	// /john/doe
	micro := fiber.New()
	app.Mount("/john", micro)
	micro.Get("/doe", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	//

	app.Listen(":3000")
}
