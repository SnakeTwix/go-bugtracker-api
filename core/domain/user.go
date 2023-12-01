package domain

import "time"

type User struct {
	ID           uint64  `json:"id"`
	Username     *string `json:"username" validate:"required" gorm:"unique;not null"`
	Password     *string `json:"-" validate:"required" gorm:"not null"`
	Icon         *string `json:"icon,omitempty"`
	Status       *string `json:"status" gorm:"not null"`
	Privilege    *string `json:"privilege" gorm:"not null"`
	RefreshToken *string `gorm:"unique"`

	Projects []*Project `json:"projects,omitempty" gorm:"many2many:user_projects"`
	Tickets  []*Ticket  `json:"tickets,omitempty" gorm:"many2many:ticket_assignees"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetUser struct {
	ID        uint64  `json:"id"`
	Username  *string `json:"username" validate:"required"`
	Icon      *string `json:"icon,omitempty"`
	Status    *string `json:"status"`
	Privilege *string `json:"privilege"`

	Projects []*Project `json:"projects,omitempty" swaggerignore:"true"`
	Tickets  []*Ticket  `json:"tickets,omitempty" swaggerignore:"true"`
}

type CreateUser struct {
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password,omitempty" validate:"required"`
}

type LoginUser struct {
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password,omitempty" validate:"required"`
}
