package ds

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMapEncodeDecode(t *testing.T) {

	m := map[string]interface{}{"abc": "def", "lmn": 123}

	Convey("When you encode a map", t, func() {
		b, err := ToBytes(m)
		So(err, ShouldBeNil)
		Convey("And then decode it", func() {
			Convey("Then the values should match", func() {
				d, err := ToMap(b)
				So(err, ShouldBeNil)
				So(m["abc"], ShouldEqual, d["abc"])
				So(m["lmn"], ShouldEqual, d["lmn"])
			})
		})
	})
}

type AStruct struct {
	MyString string
	MyInt    int
	MyJson   string `json:"my_json"`
}

func TestLoadStructByMap(t *testing.T) {

	var a AStruct

	Convey("When you load a map into a struct", t, func() {
		LoadStruct2(&a, map[string]interface{}{"MyString": "abc", "MyInt": 234, "my_json": "xyz", "field": "missing"})
		Convey("Then the values load accurately", func() {
			So(a.MyString, ShouldEqual, "abc")
			So(a.MyInt, ShouldEqual, 234)
			So(a.MyJson, ShouldEqual, "xyz")
		})
	})
}

func TestLoadStructByBytes(t *testing.T) {

	Convey("When you load bytes into a struct", t, func() {
		var a AStruct
		b, err := ToBytes(map[string]interface{}{"MyString": "abc", "MyInt": 234, "my_json": "xyz", "field": "missing"})
		So(err, ShouldBeNil)
		LoadStruct(&a, b)
		Convey("Then the values load accurately", func() {
			So(a.MyString, ShouldEqual, "abc")
			So(a.MyInt, ShouldEqual, 234)
			So(a.MyJson, ShouldEqual, "xyz")
		})
	})
}
