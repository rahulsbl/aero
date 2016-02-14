package schema

import (
	"time"
)

type Auditable struct {
	Action     string    `sql:"type:varchar(6);not null;DEFAULT:insert`
	ActionedAt time.Time `sql:"not null;DEFAULT:current_timestamp"`
}

type Timable struct {
	CreatedAt time.Time `sql:"not null;DEFAULT:current_timestamp"`
	UpdatedAt time.Time `sql:"not null;DEFAULT:current_timestamp"`
}

type Trackable struct {
	Ip        string `sql:"type:varchar(15);not null"`
	UpdatedBy uint   `sql:"not null;"`
}

type Deletable struct {
	Deleted   uint16     `sql:"not null;DEFAULT:0;"`
	DeletedAt *time.Time `sql:"null;"`
}
