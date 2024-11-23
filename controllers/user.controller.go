package controllers

import (
	"crud_api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) CreateUser(ctx *fiber.Ctx) error {
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

	result := uc.DB.Create(&newUser)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": newUser})
}

func (uc *UserController) UpdateUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	var payload *models.User

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var user models.User
	result := uc.DB.First(&user, "id = ?", userId)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "User not found"})
	}

	user.Username = payload.Username
	user.Email = payload.Email
	user.Password = payload.Password
	user.UpdatedAt = time.Now()

	uc.DB.Save(&user)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": user})
}

func (uc *UserController) FindUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	var user models.User
	result := uc.DB.First(&user, "id = ?", userId)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "User not found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": user})
}

func (uc *UserController) FindUsers(ctx *fiber.Ctx) error {
	var users []models.User
	result := uc.DB.Find(&users)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": users})
}

func (uc *UserController) DeleteUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	result := uc.DB.Delete(&models.User{}, "id = ?", userId)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "User not found"})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
