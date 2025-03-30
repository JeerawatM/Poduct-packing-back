package main

import (
	"fmt"
	"go-backend/routes"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// เชื่อมต่อฐานข้อมูล PostgreSQL
func connectDB() {
	dsn := os.Getenv("DATABASE_URL") // ดึงค่าการเชื่อมต่อจากตัวแปรแวดล้อม
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Connected to Neon PostgreSQL!")
}

func main() {
	connectDB() // เชื่อมต่อฐานข้อมูล

	// สร้าง Router และเชื่อมกับ routes.go
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("❌ Failed to get *sql.DB from *gorm.DB:", err)
	}
	r := routes.Router(sqlDB)

	// ตั้งค่าหมายเลขพอร์ต (ใช้ค่าจาก ENV ถ้ามี)
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000" // ค่าเริ่มต้น
	}

	fmt.Println("🚀 Server is running on port:", port)
	r.Run(":" + port) // เริ่มเซิร์ฟเวอร์
}
