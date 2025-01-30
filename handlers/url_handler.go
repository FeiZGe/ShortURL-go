package handlers

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/FeiZGe/ShortURL-go/config"
	"github.com/FeiZGe/ShortURL-go/models"
	"github.com/FeiZGe/ShortURL-go/utils"
)

// CreateShortURL สร้าง Short URL
func CreateShortURL(c fiber.Ctx) error {
	type Request struct {
		LongURL    string `json:"long_url"`
		ExpiryDays int    `json:"expiry_days"`
	}
	var req Request

	// อ่านข้อมูลจาก Request
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// กำหนดวันหมดอายุ
	expiryDate := time.Time{}
	if req.ExpiryDays > 0 {
		expiryDate = time.Now().Add(time.Duration(req.ExpiryDays) * 24 * time.Hour)
	}

	// สร้าง Short URL
	shortCode := utils.GenerateShortCode()
	newURL := models.URL{
		ShortURL:   shortCode,
		LongURL:    req.LongURL,
		ExpiryDate: expiryDate,
	}

	// บันทึกลงฐานข้อมูล
	config.DB.Create(&newURL)

	return c.JSON(fiber.Map{"short_url": "http://localhost:3000/" + shortCode})
}

// RedirectURL ใช้สำหรับ Redirect ไปยัง Long URL
func RedirectURL(c fiber.Ctx) error {
	shortCode := c.Params("short")
	var url models.URL

	// ค้นหา URL
	if err := config.DB.Where("short_url = ?", shortCode).First(&url).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "URL not found"})
	}

	// ตรวจสอบ Expiry Date
	if !url.ExpiryDate.IsZero() && url.ExpiryDate.Before(time.Now()) {
		config.DB.Delete(&url)
		return c.Status(410).JSON(fiber.Map{"error": "This short URL has expired"})
	}

	// เพิ่มจำนวนคลิก
	config.DB.Model(&url).Update("ClickCount", url.ClickCount+1)

	// Redirect ไปยัง URL ต้นฉบับ
	return c.Redirect().To(url.LongURL)

}

// GetClickStats ดึงข้อมูลสถิติการคลิก
func GetClickStats(c fiber.Ctx) error {
	shortCode := c.Params("short")
	var url models.URL

	// ค้นหา URL
	if err := config.DB.Where("short_url = ?", shortCode).First(&url).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "URL not found"})
	}

	return c.JSON(fiber.Map{
		"short_url":   "http://localhost:3000/" + url.ShortURL,
		"long_url":    url.LongURL,
		"click_count": url.ClickCount,
		"expiry_date": url.ExpiryDate,
	})
}
