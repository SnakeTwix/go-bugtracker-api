package domain

import (
	"time"
)

type UserProject struct {
	UserID      uint64 `gorm:"primaryKey"`
	ProjectID   uint64 `gorm:"primaryKey"`
	AccessLevel *int32

	CreatedAt time.Time
}
