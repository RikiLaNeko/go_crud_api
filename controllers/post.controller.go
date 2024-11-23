package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

func (pc *PostController) CreatePost(ctx *fiber.Ctx) error {
	currentUser := ctx.Locals("currentUser").(models.User)
	var payload *models.CreatePostRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	now := time.Now()
	newPost := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Post with that title already exists"})
		}
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": newPost})
}

func (pc *PostController) UpdatePost(ctx *fiber.Ctx) error {
	postId := ctx.Params("postId")
	currentUser := ctx.Locals("currentUser").(models.User)

	var payload *models.UpdatePost
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var updatedPost models.Post
	result := pc.DB.First(&updatedPost, "id = ?", postId)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No post with that title exists"})
	}
	now := time.Now()
	postToUpdate := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      currentUser.ID,
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedPost).Updates(postToUpdate)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": updatedPost})
}

func (pc *PostController) FindPostById(ctx *fiber.Ctx) error {
	postId := ctx.Params("postId")

	var post models.Post
	result := pc.DB.First(&post, "id = ?", postId)
	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No post with that title exists"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": post})
}

func (pc *PostController) FindPosts(ctx *fiber.Ctx) error {
	page := ctx.Query("page", "1")
	limit := ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var posts []models.Post
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&posts)
	if results.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(posts), "data": posts})
}

func (pc *PostController) DeletePost(ctx *fiber.Ctx) error {
	postId := ctx.Params("postId")

	result := pc.DB.Delete(&models.Post{}, "id = ?", postId)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No post with that title exists"})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
