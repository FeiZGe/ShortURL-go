package utils

import (
	"math/rand"
	"os"
	"time"
)

// GenerateShortCode สุ่มสร้าง Short URL ยาว 6 ตัวอักษร
func GenerateShortCode() string {
	// สร้างตัว generator ที่ใช้ seed ใหม่ในแต่ละรอบการเรียก
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// สุ่มรหัส 6 ตัว
	shortCode := make([]byte, 6)
	for i := range shortCode {
		shortCode[i] = os.Getenv("CHARSET")[r.Intn(len(os.Getenv("CHARSET")))]
	}

	return string(shortCode)
}
