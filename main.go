package main

import (
	_ "blog_service/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define routes for your API
	app.Get("/posts", handlers.GetPosts)
	app.Get("/posts/:id", handlers.GetPostByID)
	app.Post("/posts", handlers.CreatePost)
	app.Put("/posts/:id", handlers.UpdatePost)
	app.Delete("/posts/:id", handlers.DeletePost)

	app.Listen(":8080")
}
