package routes

import (
	controllers "belajar_golang/controller"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	app := gin.Default()

	app.GET("/", controllers.TestIndex)
	app.GET("/test2", controllers.TestIndex2)
	app.POST("/test_post", controllers.TestPost)
	app.POST("/test_post_form", controllers.TestPostForm)
	app.POST("/test_post_form2", controllers.TestPostForm2)

	return app
}
