package routes

import (
	controllers "belajar_golang/controller"

	"github.com/gin-gonic/gin"

	"github.com/gofiber/fiber/v2"
)

func Routes() *gin.Engine {
	app := gin.Default()

	app.GET("/", controllers.TestIndex)
	app.GET("/test2", controllers.TestIndex2)
	app.POST("/test_post", controllers.TestPost)
	app.POST("/test_post_form", controllers.TestPostForm)
	app.POST("/test_post_form2", controllers.TestPostForm2)
	app.POST("/upload_minio_file", controllers.TestPostFormUploadMinio)

	return app
}

func RoutesV2Fiber() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		controllers.TestIndexFiber(c)
		return nil
	})
	app.Get("/test2", func(c *fiber.Ctx) error {
		controllers.TestIndex2F(c)
		return nil
	})

	app.Get("/test2_gagal", func(c *fiber.Ctx) error {
		controllers.TestIndex2FGagal(c)
		return nil
	})

	app.Post("/test_post", func(c *fiber.Ctx) error {
		controllers.TestPostFiber(c)
		return nil
	})
	app.Post("/test_post_form", func(c *fiber.Ctx) error {
		controllers.TestPostFormFiber(c)
		return nil
	})
	app.Post("/test_post_form2", func(c *fiber.Ctx) error {
		controllers.TestPostForm2Fiber(c)
		return nil
	})
	app.Post("/upload_minio_file", func(c *fiber.Ctx) error {
		controllers.TestPostFormUploadMinioFiber(c)
		return nil
	})

	return app
}
