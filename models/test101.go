package models

//* สร้าง Model สำหรับตาราง `product_`
type Test struct {
	Test01  string  `gorm:"column:test01"`
	Test02  string  `gorm:"column:imagetest02"`
	Test03  int     `gorm:"column:test03"`
	Test04  int     `gorm:"column:test04"`
	Test05  float64  `gorm:"column:test05;type:DECIMAL"`
	Test06  float64  `gorm:"column:test06;type:DECIMAL"`
}

//* ฟังก์ชันสำหรับตั้งชื่อคอลัมน์ในฐานข้อมูล
func (Test) TableName() string {
	return "test_101" //* ชื่อตารางในฐานข้อมูล
}
