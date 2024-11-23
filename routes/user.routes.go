package routes

import (
	"crud_api/controllers"
	"crud_api/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg fiber.Router) {
	router := rg.Group("/users")
	router.Use(middleware.DeserializeUser())
	router.Post("/", uc.userController.CreateUser)
	router.Get("/", uc.userController.FindUsers)
	router.Put("/:userId", uc.userController.UpdateUser)
	router.Get("/:userId", uc.userController.FindUserById)
	router.Delete("/:userId", uc.userController.DeleteUser)
}
