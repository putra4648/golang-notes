package routes

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		var result = map[string]string{"message": "You see this? Then server is running"}
		context.JSON(200, result)
	})

	router.GET("/notes", GetNotes())
	router.POST("/notes", AddNote())

	return router
}
