package routes

import (
	controllers "belajar_golang/controller"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	services "belajar_golang/service"
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

	//struct for dbConnection
	type DBConnection struct {
		Host string
		Port string
		User string
		Pass string
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println(err.Error())
			return nil
		},
	})

	app.Use(recover.New())

	mainDB, err_db_main := services.MainDBServiceConnection()

	if err_db_main != nil {
		log.Fatal(err_db_main)
	}

	app.Use(func(c *fiber.Ctx) error {
		//Set Local MainDB
		services.SetLocal[services.TestMainDBService](c, "mainDB", *mainDB)
		return c.Next()
	})

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
	app.Post("/test_post_form_fiber", func(c *fiber.Ctx) error {
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

	app.Get("/test_db", func(c *fiber.Ctx) error {
		controllers.TestDB1(c)
		return nil
	})

	app.Get("/test_db2", func(c *fiber.Ctx) error {
		controllers.TestDBSQLx(c)
		return nil
	})

	return app
}
