package orm

import (
	"time"

	"github.com/thejackrabbit/aero/db"
)

type Trail struct {
	Action     string    `sql:"TYPE:varchar(6);not null;DEFAULT:'insert'"`
	ActionedAt time.Time `sql:"not null;DEFAULT:current_timestamp"`
}

type Actor struct {
	ActorID uint   `sql:"not null;"`
	ActorIP string `sql:"TYPE:varchar(15);not null"`
}

type Timed struct {
	CreatedAt time.Time `sql:"not null;DEFAULT:current_timestamp"`
	UpdatedAt time.Time `sql:"not null;DEFAULT:current_timestamp"`
}

type Persistent struct {
	Deleted   uint8      `sql:"TYPE:tinyint unsigned;not null;DEFAULT:0;"`
	DeletedAt *time.Time `sql:"null;"`
}

type WebResource struct {
	URLWeb string   `sql:"TYPE:text;not null;"`
	URLApp string   `sql:"TYPE:text;not null;"`
	URLs   *db.JDoc `sql:"TYPE:json;"`
}

type Tagged struct {
	Tags *db.JArr `sql:"TYPE:json;"`
}
