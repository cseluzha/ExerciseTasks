package models

import (
	uuid "github.com/google/uuid"
)

type Status struct {
	Id   uuid.UUID `json:"IdStatus" db:"Id" sql:",type:uuid"`
	Name string    `json:"Name" db:"Name"`
}
