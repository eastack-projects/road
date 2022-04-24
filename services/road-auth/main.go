package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"golang.org/x/crypto/argon2"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Username: "EaStack", Password: "Helloworld"})

	var user User
	db.First(&user, 1)
	db.First(&user, "username = ?", "EaStack")

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.POST("/auth/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8082")
}
