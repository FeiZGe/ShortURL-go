package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v3"
	
	"github.com/FeiZGe/ShortURL-go/config"
	"github.com/FeiZGe/ShortURL-go/routes"
)

func main() {
	
	// โหลดค่าจาก .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// เริ่มต้นเชื่อมต่อฐานข้อมูล
	config.InitDB()

	// สร้างแอป Fiber
	app := fiber.New()

	// กำหนดเส้นทาง API
	routes.SetupRoutes(app)

	// เริ่มต้นเซิร์ฟเวอร์ที่พอร์ต 3000
	log.Println("Server is running on http://localhost:3000")
	app.Listen(":3000")
}