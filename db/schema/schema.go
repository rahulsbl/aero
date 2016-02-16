package schema

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thejackrabbit/aero/panik"
	"github.com/thejackrabbit/aero/str"
)

var Dbo *gorm.DB

func Build(delExisting bool, models ...interface{}) {
	panik.If(Dbo == nil, "Dbo reference is nil")
	Dbo.LogMode(true)

	// create tables
	for _, model := range models {
		tbl := NewTable(model)

		// delete tables
		if delExisting && tbl.exists() {
			tbl.drop()
		}

		// main migration
		Dbo.AutoMigrate(model)

		// updated_at trigger (except on history tables)
		updAt := tbl.field("updated_at")
		if !tbl.isHistory() && updAt != nil {
			if !strings.Contains(strings.ToLower(updAt.Extra), strings.ToLower("on update current_timestamp")) {
				sql := fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s ON UPDATE CURRENT_TIMESTAMP", tbl.name, updAt.info())
				sqlExec(sql)
			}
		}
	}

	// add foreign keys
	for _, model := range models {
		tbl := NewTable(model)
		if !tbl.isHistory() {
			// add foreign keys
			mt := reflect.TypeOf(model).Elem()
			num := mt.NumField()
			for i := 0; i < num; i++ {
				fld := mt.FieldByIndex([]int{i})
				tag := fld.Tag
				if len(tag.Get("fk")) > 0 {
					fk := str.SnakeCase(fld.Name)
					fmt.Println(tag.Get("fk"), fk)
					Dbo.Model(model).AddForeignKey(fk, tag.Get("fk"), "RESTRICT", "RESTRICT")
				}
			}
		}
	}

	// create history tables
	for _, model := range models {
		tbl := NewTable(model)
		if tbl.isHistory() {
			// drop auto_increments & primary key
			tbl.readyHistoryTable()
		}
	}

	// triggers to send data to history tables
	for _, model := range models {
		tbl := NewTable(model)
		if !tbl.isHistory() {
			// setup triggers
			if tbl.exists() && tbl.historyExists() {
				tbl.setupHistoryTriggers()
			}
		}
	}
}

func sqlHasRows(sql string) bool {
	rows, err := Dbo.Raw(sql).Rows()
	panik.On(err)
	defer rows.Close()
	return rows.Next()
}

func sqlExec(sql string) {
	panik.On(Dbo.Exec(sql).Error)
}
