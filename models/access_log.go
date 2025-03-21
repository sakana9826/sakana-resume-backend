package models

import (
	"time"

	"gorm.io/gorm"
)

type AccessLog struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	AccessCode string         `json:"access_code"`
	IP         string         `json:"ip"`
	UserAgent  string         `json:"user_agent"`
	AccessedAt time.Time      `json:"accessed_at"`
	ExpiresAt  time.Time      `json:"expires_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
