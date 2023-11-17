package domain

import "time"

type TicketTag struct {
	ID   uint64
	Name *string

	ProjectID uint64
	Project   *Project
	Tickets   []*Ticket `gorm:"many2many:ticket_tags"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (TicketTag) TableName() string {
	return "project_ticket_tags"
}
