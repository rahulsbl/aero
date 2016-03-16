package orm

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestModelMustInsert(t *testing.T) {

	type Abc struct {
		Alphabet string `insert:"must"`
	}

	Convey("Missing insert-field must give an error", t, func() {
		ok, errs := Insertable(Abc{}, map[string]string{"some": "thing"})
		So(ok, ShouldBeFalse)
		So(len(errs), ShouldEqual, 1)

		Convey("Present insert-field should not give error", func() {
			ok, errs := Insertable(Abc{}, map[string]string{"alphabet": "thing"})
			So(ok, ShouldBeTrue)
			So(errs, ShouldBeNil)
		})
	})

	// type Def struct {
	// 	Alphabet string `insert:"no"`
	// }

	// Convey("Missing insert field must give an error", t, func() {
	// 	ok, errs := Insertable(Def{}, map[string]string{"alphabet": ""})
	// 	So(ok, ShouldBeFalse)
	// 	So(len(errs), ShouldEqual, 1)
	// })

}
