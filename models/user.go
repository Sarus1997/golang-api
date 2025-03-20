package models

import (
	"time"
)

type User struct {
	ID                  string     `gorm:"primaryKey;type:char(255)"`
	Username            string     `gorm:"unique;not null"`
	Email               string     `gorm:"unique;not null"`
	PasswordHash        string     `gorm:"column:password_hash;not null"`
	FirstName           string     `gorm:"column:f_name"`
	LastName            string     `gorm:"column:l_name"`
	ProfilePicture      string     `gorm:"column:profile_picture"`
	OAuthProvider       string     `gorm:"column:oauth_provider;type:enum('email','facebook','google','apple','line');default:'email'"`
	Role                string     `gorm:"column:role;type:enum('user','admin','moderator');default:'user'"`
	OAuthID             string     `gorm:"column:oauth_id"`
	Status              string     `gorm:"column:status;default:'active'"`
	LastLoginAt         *time.Time `gorm:"column:last_login_at"`
	ResetToken          string     `gorm:"column:reset_token"`
	ResetTokenExpiresAt *time.Time `gorm:"column:reset_token_expires_at"`
	CreatedAt           *time.Time `gorm:"column:created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at"`
}

// ตั้งชื่อให้ตรงกับตารางในฐานข้อมูล
func (User) TableName() string {
	return "users"
}
