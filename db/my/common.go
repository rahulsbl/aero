package my

import (
	"time"

	"github.com/rightjoin/aero/db"
)

type IDKey struct {
	ID uint `sql:"auto_increment;not null;primary_key" json:"id" insert:"no" update:"no"`
}

type Timed struct {
	CreatedAt time.Time `sql:"not null;DEFAULT:current_timestamp" json:"created_at" insert:"no" update:"no"`
	UpdatedAt time.Time `sql:"not null;DEFAULT:current_timestamp" json:"updated_at" insert:"no" update:"no"`
}

type Persistent struct {
	Deleted   uint8      `sql:"TYPE:tinyint unsigned;not null;DEFAULT:'0'" json:"deleted" insert:"no"`
	DeletedAt *time.Time `sql:"null;" json:"deleted_at" insert:"no" update:"no"`
}

type WWW struct {
	URLWeb       string   `sql:"TYPE:varchar(256);not null" json:"url_web" unique:"true"`
	URLWebOld    *db.JArr `sql:"TYPE:json;" json:"-" insert:"no" update:"no"`
	MetaTitle    string   `sql:"TYPE:varchar(512);not null;DEFAULT:''" json:"meta_title"`
	MetaDesc     string   `sql:"TYPE:varchar(512);not null;DEFAULT:''" json:"meta_desc"`
	MetaKeywords string   `sql:"TYPE:varchar(512);not null;DEFAULT:''" json:"meta_keywords"`
	Sitemap      uint8    `sql:"TYPE:tinyint unsigned;not null;DEFAULT:'1'" json:"sitemap"`
}

type Tagged struct {
	Tags *db.JArr `sql:"TYPE:json;" json:"tags"`
}

type Ordered struct {
	Sequence uint16 `sql:"TYPE:smallint unsigned;not null;DEFAULT:'1'" json:"sequence"`
}

type Modifier struct {
	Actor *db.JDoc `sql:"TYPE:json" json:"actor"`
}

type AuditTrail struct {
	Action     string    `sql:"TYPE:varchar(6);not null;DEFAULT:'insert'"`
	ActionedAt time.Time `sql:"not null;DEFAULT:current_timestamp"`
}

type DynamicFields struct {
	Info *db.JDoc `sql:"TYPE:json" json:"info"`
}

type PopulateDB interface {
	InitRecords() []interface{}
}

type Triggers interface {
	InitTriggers() []string
}
