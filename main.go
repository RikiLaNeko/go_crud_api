package main

import (
	"crud_api/controllers"
	"crud_api/initializers"
	"crud_api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

var (
	server              *fiber.App
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	server = fiber.New()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	server.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080, " + config.ClientOrigin,
		AllowCredentials: true,
	}))

	server.Get("/api/healthchecker", func(ctx *fiber.Ctx) error {
		message := "Welcome to Golang with Gorm and Postgres"
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": message})
	})

	api := server.Group("/api")
	AuthRouteController.AuthRoute(api)
	UserRouteController.UserRoute(api)
	PostRouteController.PostRoute(api)

	log.Fatal(server.Listen(":" + config.ServerPort))
}
