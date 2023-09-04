package model

import "time"

type UUID string

type Common struct {
	ID        UUID
	CreatedAt time.Time
	DeletedAt time.Time
	UpdatedAt time.Time
	CreatedBy UUID
	UpdatedBy UUID
	DeletedBy UUID
}
