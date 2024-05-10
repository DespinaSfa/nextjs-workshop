package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

func BasicAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user User
		if err := db.Where("username = ?", username).First(&user).Error; err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if user.Password != password {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func main() {
	// Verbindung zur Datenbank herstellen
	dsn := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Automatisches Migrations
	db.AutoMigrate(&User{})

	// Beispielbenutzer erstellen (optional)
	db.Create(&User{Username: "username", Password: "password"})

	// Router initialisieren
	router := gin.Default()

	// Middleware für die Basic-Auth hinzufügen
	router.Use(BasicAuth(db))

	// Beispielroute
	router.GET("/api/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "This is sensitive data!"})
	})

	// Server starten
	router.Run(":8080")
}
