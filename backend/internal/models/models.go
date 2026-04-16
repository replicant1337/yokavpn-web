package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	IsAdmin   bool           `gorm:"default:false" json:"is_admin"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Subscription struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	UserID         uint           `gorm:"not null" json:"user_id"`
	User           User           `json:"user"`
	PlanName       string         `json:"plan_name"`
	RemnaUserID    string         `json:"remna_user_id"`
	RemnaSubLink   string         `json:"remna_sub_link"`
	ShortID        string         `gorm:"uniqueIndex" json:"short_id"`
	AuthKey        string         `gorm:"uniqueIndex;not null" json:"auth_key"`
	TrafficTotal   int64          `json:"traffic_total"`
	TrafficUsed    int64          `json:"traffic_used"`
	ExpiresAt      time.Time      `json:"expires_at"`
	Status         string         `gorm:"default:'active'" json:"status"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
