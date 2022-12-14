package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        2,
	}))
	log.Fatal(app.Listen(":3000"))
}
