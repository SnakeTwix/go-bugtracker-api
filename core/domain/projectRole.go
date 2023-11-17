package domain

import "time"

type ProjectRole struct {
	ID          uint64
	Name        *string
	Color       *string
	AccessLevel *uint64

	ProjectID uint64
	Project   Project

	CreatedAt time.Time
	UpdatedAt time.Time
}
