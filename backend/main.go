package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-backend/routes"
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

	// แปลง gorm.DB เป็น sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("❌ Failed to get SQL DB:", err)
	}

	// ส่ง sqlDB เข้า Router
	router := routes.Router(sqlDB)
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	log.Println("🚀 Server is running on port:", port)
	router.Run(":" + port)
}
