package models

import (
	uuid "github.com/google/uuid"
	"time"
)

type Task struct {
	Id          uuid.UUID `json:"IdTask" db:"Id" sql:",type:uuid"`
	Title       string    `json:"Title" db:"Title"`
	Description string    `json:"Description" db:"Description"`
	CreatedOn   time.Time `json:"CreatedOn" db:"CreatedOn"`
	UpdatedOn   time.Time `json:"UpdatedOn" db:"UpdatedOn"`
	Status      Status    `json:"Status" db:"StatusId"`
}
