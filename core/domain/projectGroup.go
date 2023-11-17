package domain

import "time"

type ProjectGroup struct {
	ID          uint64
	Name        *string `gorm:"not null"`
	Description *string
	Icon        *string

	UserID   uint64 `gorm:"not null"`
	User     User
	Projects []*Project `gorm:"foreignKey:GroupID"`

	CreatedAt time.Time
	UpdateAt  time.Time
}
