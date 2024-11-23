package routes

import (
	"crud_api/controllers"
	"crud_api/middleware"
	"github.com/gofiber/fiber/v2"
)

type PostRouteController struct {
	postController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}

func (pc *PostRouteController) PostRoute(rg fiber.Router) {
	router := rg.Group("/posts")
	router.Use(middleware.DeserializeUser())
	router.Post("/", pc.postController.CreatePost)
	router.Get("/", pc.postController.FindPosts)
	router.Put("/:postId", pc.postController.UpdatePost)
	router.Get("/:postId", pc.postController.FindPostById)
	router.Delete("/:postId", pc.postController.DeletePost)
}
