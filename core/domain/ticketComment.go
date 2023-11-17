package domain

import (
	"gorm.io/datatypes"
	"time"
)

type TicketComment struct {
	Body *datatypes.JSON

	PosterID uint64
	Poster   *User `gorm:"foreignKey:PosterID"`
	TicketID uint64
	Ticket   *Ticket

	CreatedAt time.Time
	UpdatedAt time.Time
}
