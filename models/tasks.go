package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
