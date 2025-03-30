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
// package main

// import (
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"log"
// 	"os"
// )

// var db *gorm.DB // ประกาศตัวแปร db ให้เป็น global

// func connectDB() {
// 	dsn := os.Getenv("DATABASE_URL") // ดึงค่าจาก Environment Variable
// 	var err error
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // ใช้ตัวแปร db

// 	if err != nil {
// 		log.Fatal("❌ Failed to connect to database:", err)
// 	}

// 	log.Println("✅ Connected to Neon PostgreSQL!")
// }

//	func main() {
//		connectDB()
//	}
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

	// ใช้ Router
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("❌ Failed to get *sql.DB from *gorm.DB:", err)
	}
	r := routes.Router(sqlDB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	fmt.Println("🚀 Server is running on port:", port)
	r.Run(":" + port)
}
