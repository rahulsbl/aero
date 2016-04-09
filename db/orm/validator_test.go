package orm

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/thejackrabbit/aero/db"
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
}

func TestModelMustUpdate(t *testing.T) {

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

func TestModelJsonArray(t *testing.T) {

	var ok bool
	var errs []error

	type Ghi struct {
		Array    db.JsonA
		ArrayPtr *db.JsonA
	}

	Convey("String will fail json array field", t, func() {
		ok, errs = Insertable(Ghi{}, map[string]string{"array": "thing"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

		ok, errs = Updatable(Ghi{}, map[string]string{"array": "thing"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

		ok, errs = Insertable(Ghi{}, map[string]string{"array_ptr": "thinger"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

		ok, errs = Updatable(Ghi{}, map[string]string{"array_ptr": "thinger"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)
	})

	Convey("And array will pass it", t, func() {
		ok, errs = Insertable(Ghi{}, map[string]string{"array": "[1,2,3]"})
		So(ok, ShouldBeTrue)
		So(len(errs), ShouldEqual, 0)

		ok, errs = Updatable(Ghi{}, map[string]string{"array": "[1,2,3]"})
		So(ok, ShouldBeTrue)
		So(len(errs), ShouldEqual, 0)

		ok, errs = Insertable(Ghi{}, map[string]string{"array_ptr": "[\"abc\", 1]"})
		So(ok, ShouldBeTrue)
		So(len(errs), ShouldEqual, 0)

		ok, errs = Updatable(Ghi{}, map[string]string{"array_ptr": "[\"abc\", 1]"})
		So(ok, ShouldBeTrue)
		So(len(errs), ShouldEqual, 0)
	})

}

func TestModelJsonDoc(t *testing.T) {

	var ok bool
	var errs []error

	type Jkl struct {
		Mapper    db.JsonM
		MapperPtr *db.JsonM
	}

	Convey("String will fail json mapper field", t, func() {
		ok, errs = Insertable(Jkl{}, map[string]string{"mapper": "thing"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

		ok, errs = Updatable(Jkl{}, map[string]string{"mapper": "thing"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

		ok, errs = Insertable(Jkl{}, map[string]string{"mapper_ptr": "thinger"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

		ok, errs = Updatable(Jkl{}, map[string]string{"mapper_ptr": "thinger"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

		ok, errs = Updatable(Jkl{}, map[string]string{"mapper_ptr": "[1,2,3]"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

	})

	Convey("And mapper will pass it", t, func() {
		ok, errs = Insertable(Jkl{}, map[string]string{"mapper": "{\"key\":\"value\"}"})
		So(ok, ShouldBeTrue)
		So(len(errs), ShouldEqual, 0)

		ok, errs = Updatable(Jkl{}, map[string]string{"mapper": "{\"key\":\"value\"}"})
		So(ok, ShouldBeTrue)
		So(len(errs), ShouldEqual, 0)

		ok, errs = Insertable(Jkl{}, map[string]string{"mapper_ptr": "{\"key\":\"value\"}"})
		So(ok, ShouldBeTrue)
		So(len(errs), ShouldEqual, 0)

		ok, errs = Updatable(Jkl{}, map[string]string{"mapper_ptr": "{\"key\":\"value\"}"})
		So(ok, ShouldBeTrue)
		So(len(errs), ShouldEqual, 0)
	})

}
