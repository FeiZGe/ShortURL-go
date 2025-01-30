package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/FeiZGe/ShortURL-go/handlers"
)

// SetupRoutes ตั้งค่าเส้นทาง API
func SetupRoutes(app *fiber.App) {
	app.Post("/shorten", handlers.CreateShortURL)  // API สำหรับย่อลิงก์
	app.Get("/:short", handlers.RedirectURL)       // API สำหรับ Redirect
	app.Get("/stats/:short", handlers.GetClickStats) // API สำหรับดูสถิติ
}
