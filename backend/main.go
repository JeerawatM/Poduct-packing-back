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
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB // Global DB Connection

func connectDB() {
	dsn := "host=ep-crimson-firefly-a15c9507-pooler.ap-southeast-1.aws.neon.tech user=neondb_admin password=npg_FsqMG6bgpRo0 dbname=neondb port=10000 sslmode=disable" // ดึงค่าจาก Environment Variable
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // ใช้ตัวแปร db

	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Connected to Neon PostgreSQL!")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 🔹 Example Route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Backend!"})
	})

	return router
}

func main() {
	connectDB() // เชื่อมต่อ DB

	// 🔹 อ่าน PORT จาก Environment ถ้าไม่มีให้ใช้ค่า Default
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	router := setupRouter()
	log.Printf("🚀 Server is running on port %s", port)
	router.Run(":" + port) // Start Server
}
