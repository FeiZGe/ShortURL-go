package config

import (
	"log"
	"github.com/FeiZGe/ShortURL-go/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB ทำการเชื่อมต่อฐานข้อมูล SQLite
func InitDB() {
	var err error

	// เชื่อมต่อ SQLite และสร้างไฟล์ฐานข้อมูลถ้ายังไม่มี
	DB, err = gorm.Open(sqlite.Open("storage/urls.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// ทำการ migrate ตาราง URL
	DB.AutoMigrate(&models.URL{})
	log.Println("Database migrated successfully.")
}
