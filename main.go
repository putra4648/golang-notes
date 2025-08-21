package main

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"golang-notes/entities"

	"github.com/google/uuid"
)

func main() {
	db, err := gorm.Open((postgres.New(postgres.Config{
		DSN:                  "host=localhost user=user password=user123 dbname=notes_db port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	})), &gorm.Config{})

	println("Connected to database", db.Name())

	if err != nil {
		panic("failed to connect to database")
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		notes, err := gorm.G[entities.Note](db).Find(c)

		if err != nil {
			println(err)
			c.JSON(500, gin.H{"error": "Failed to fetch notes"})
			return
		}
		c.JSON(200, notes)
	})

	router.GET("/add", func(c *gin.Context) {
		note := entities.Note{
			ID:      uuid.NewString(),
			Title:   "Sample Note",
			Content: "This is a sample note content.",
			Created: time.Now(),
			Updated: time.Now(),
		}

		if err := db.Create(&note).Error; err != nil {
			println(err)
			c.JSON(500, gin.H{"error": "Failed to create note"})
			return
		}

		c.JSON(200, gin.H{"message": "Note created successfully", "note": note})
	})

	router.Run()

}
