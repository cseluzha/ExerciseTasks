package models

import (
	uuid "github.com/google/uuid"
	pq "github.com/lib/pq"
	"time"
)

type Task struct {
	Id          uuid.UUID   `json:"IdTask" db:"Id" sql:",type:uuid"`
	Title       string      `json:"Title" db:"Title"`
	Description string      `json:"Description" db:"Description"`
	CreatedOn   time.Time   `json:"CreatedOn" db:"CreatedOn"`
	UpdatedOn   pq.NullTime `json:"UpdatedOn" db:"UpdatedOn"`
	Status      Status      `json:"Status,omitempty" db:"StatusId"`
	Active      bool        `json:"Active"  db:"Active"`
}
