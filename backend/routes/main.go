package routes

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/notes", GetNotes())
	router.POST("/note", AddNote())

	return router
}
