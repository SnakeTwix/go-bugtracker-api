package domain

import (
	"gorm.io/datatypes"
	"time"
)

type Ticket struct {
	ID       uint64          `json:"id"`
	Priority *string         `json:"priority"`
	Title    *string         `json:"title"`
	Body     *datatypes.JSON `json:"body"`
	Status   *string         `json:"status"`

	CreatorID uint64 `json:"creatorId"`
	User      *User  `gorm:"foreignKey:CreatorID"`
	ProjectID uint64 `json:"projectId"`
	Project   *Project
	Tags      []*TicketTag `gorm:"many2many:ticket_tags"`
	Users     []*User      `gorm:"many2many:ticket_assignees"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
