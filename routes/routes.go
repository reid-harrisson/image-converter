package routes

import (
	_ "image-converter/docs"
	"image-converter/handlers"
	"image-converter/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func ConfigRoutes(server *server.Server) {
	server.App.Get("/swagger/*", swagger.HandlerDefault)
	server.App.Get("/", redirectToSwagger)

	apiV1 := server.App.Group("/api/v1")

	imageHandler := handlers.CreateImageHandler()
	imageGroup := apiV1.Group("/image")
	imageGroup.Post("", imageHandler.Convert)
}

func redirectToSwagger(context *fiber.Ctx) error {
	return context.Redirect("/swagger/index.html")
}
