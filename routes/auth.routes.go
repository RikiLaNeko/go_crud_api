package routes

import (
	"crud_api/controllers"
	"crud_api/middleware"
	"github.com/gofiber/fiber/v2"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg fiber.Router) {
	router := rg.Group("/auth")

	router.Post("/register", rc.authController.SignUpUser)
	router.Post("/login", rc.authController.SignInUser)
	router.Get("/refresh", rc.authController.RefreshAccessToken)
	router.Get("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
}
