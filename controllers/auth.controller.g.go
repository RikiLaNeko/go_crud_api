package controllers

import (
	"crud_api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) SignUpUser(ctx *fiber.Ctx) error {
	var payload *models.User

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	now := time.Now()
	newUser := models.User{
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ac.DB.Create(&newUser)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": newUser})
}

func (ac *AuthController) SignInUser(ctx *fiber.Ctx) error {
	var payload *models.User

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var user models.User
	result := ac.DB.Where("email = ? AND password = ?", payload.Email, payload.Password).First(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid email or password"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": user})
}

func (ac *AuthController) RefreshAccessToken(ctx *fiber.Ctx) error {
	// Implement token refresh logic here
	return ctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"status": "fail", "message": "Not implemented"})
}

func (ac *AuthController) LogoutUser(ctx *fiber.Ctx) error {
	// Implement logout logic here
	return ctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"status": "fail", "message": "Not implemented"})
}
