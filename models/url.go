package models

import (
	"time"
)

// URL เป็นโครงสร้างข้อมูลที่ใช้เก็บ URL ที่ถูกย่อลงในฐานข้อมูล
type URL struct {
	ID         uint      `gorm:"primaryKey"` // Primary Key
	ShortURL   string    `gorm:"uniqueIndex"`
	LongURL    string
	ClickCount int       // จำนวนคลิก
	ExpiryDate time.Time // วันหมดอายุของ URL
}
