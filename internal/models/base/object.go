package base

import "time"

type AuditFields struct {
	IsActive  bool      `gorm:"default:1" json:"is_active"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	DeletedBy string    `json:"deleted_by"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime:milli" json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
