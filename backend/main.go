package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}
	log.Println("✅ Connected to Neon PostgreSQL!")
}

func main() {
	connectDB()

	// ใช้ Gin เป็น Web Server
	r := gin.Default()

	// ทดสอบ API ว่าใช้งานได้
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "🚀 Backend is running!"})
	})

	// ใช้ PORT จาก Environment หรือค่าเริ่มต้น 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	fmt.Println("🌍 Server is running on port " + port)

	// Start server
	r.Run(":" + port)
}
