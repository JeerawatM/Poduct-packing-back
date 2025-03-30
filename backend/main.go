// package main

// import (
// 	"fmt"
// 	"go-backend/config"
// 	"go-backend/database"
// 	"go-backend/routes"
// )

//	func main() {
//		fmt.Println("System Start")
//		config.LoadConfig()
//		db := database.InitDB()
//		defer db.Close()
//		router := routes.Router(db)
//		router.Run(":8080")
//	}
package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB // ประกาศตัวแปร db ให้เป็น global

func connectDB() {
	dsn := os.Getenv("DATABASE_URL") // ดึงค่าจาก Environment Variable
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // ใช้ตัวแปร db

	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Connected to Neon PostgreSQL!")
}

func main() {
	connectDB()
}
