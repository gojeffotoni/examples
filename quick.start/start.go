package main

import "github.com/jeffotoni/quick"

func main() {
	app := quick.New()
	app.Get("/v1/user", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(200).SendString("Quick em ação com Cors❤️!")
	})
	app.Listen("0.0.0.0:8080")
}
