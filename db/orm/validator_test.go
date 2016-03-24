package orm

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestModelMustInsert(t *testing.T) {

	type Abc struct {
		Alphabet string `insert:"must"`
		Numbers  string `insert:"no"`
	}

	Convey("Missing must-insert field should give an error", t, func() {
		ok, errs := Insertable(Abc{}, map[string]string{"some": "thing"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

	})
	Convey("Present must-insert and missing no-insert field should be ok", t, func() {
		ok, errs := Insertable(Abc{}, map[string]string{"alphabet": "thing"})
		So(ok, ShouldBeTrue)
		So(errs, ShouldBeNil)
	})
	Convey("Present no-insert field should give an error", t, func() {
		ok, errs := Insertable(Abc{}, map[string]string{"alphabet": "thing", "numbers": "1234"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)
	})

	type Def struct {
		Alphabet string `update:"no"`
		Numbers  string `update:"must"`
	}

	Convey("Missing must-update field should give an error", t, func() {
		ok, errs := Updatable(Def{}, map[string]string{"some": "thing"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

	})
	Convey("Present must-update and missing no-update field should be ok", t, func() {
		ok, errs := Updatable(Def{}, map[string]string{"numbers": "thing"})
		So(ok, ShouldBeTrue)
		So(errs, ShouldBeNil)
	})
	Convey("Present no-update field should give an error", t, func() {
		ok, errs := Updatable(Def{}, map[string]string{"alphabet": "thing", "numbers": "1234"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)
	})

}
