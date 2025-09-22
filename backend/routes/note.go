package routes

import (
	"golang-notes/db"
	"golang-notes/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, dbError := db.DB()
		if dbError != nil {
			c.Error(dbError)
			return
		}

		notes, notesError := gorm.G[entities.Note](db).Find(c)

		if notesError != nil {
			println(notesError)
			c.JSON(500, gin.H{"error": "Failed to fetch notes"})
			return
		}
		c.JSON(200, notes)
	}
}

func AddNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, dbError := db.DB()
		if dbError != nil {
			c.Error(dbError)
			return
		}

		var note entities.Note
		if err := c.ShouldBindJSON(&note); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		notes := gorm.G[entities.Note](db).Create(c, &note)

		c.JSON(200, notes)
	}
}
