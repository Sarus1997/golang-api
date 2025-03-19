package models

import (
	"time"
)

// สร้าง Model สำหรับตาราง `product_`
type Product struct {
	ProductID  string    `gorm:"primaryKey;column:product_id"`
	ImageURL   string    `gorm:"column:image_url"`
	ProductName string   `gorm:"column:product_name"`
	Price      float64   `gorm:"column:price"`
	Brand      string    `gorm:"column:brand"`
	Status     string    `gorm:"column:status"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

// ฟังก์ชันสำหรับตั้งชื่อคอลัมน์ในฐานข้อมูล
func (Product) TableName() string {
	return "product_" // ชื่อตารางในฐานข้อมูล
}
