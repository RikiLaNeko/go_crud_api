package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"github.com/wpcodevo/golang-gorm-postgres/utils"
)

func DeserializeUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var accessToken string
		cookie := ctx.Cookies("access_token")

		authorizationHeader := ctx.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if cookie != "" {
			accessToken = cookie
		}

		if accessToken == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
		}

		config, _ := initializers.LoadConfig(".")
		sub, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}

		var user models.User
		result := initializers.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "The user belonging to this token no longer exists"})
		}

		ctx.Locals("currentUser", user)
		return ctx.Next()
	}
}
