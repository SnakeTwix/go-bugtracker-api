package domain

import "time"

type Project struct {
	ID          uint64
	Name        *string `gorm:"not null"`
	Description *string
	Icon        *string
	AccessLevel *int32 `gorm:"not null"`

	CreatorId uint64 `gorm:"not null"`
	Creator   User
	Members   []*User `gorm:"many2many:user_projects"`
	GroupID   uint64
	Tags      []*TicketTag

	CreatedAt time.Time
	UpdateAt  time.Time
}
