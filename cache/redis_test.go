package cache

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var redis_host string = "192.168.99.100"
var redis_port int = 32773
var redis_db int = 0

func TestRedisGetSet(t *testing.T) {

	r := NewRedis(redis_host, redis_port, redis_db)

	Convey("Given a Redis service", t, func() {

		Convey("When you Set some values against a key", func() {
			r.Set("string-test", []byte("stringy"), time.Minute*5)

			Convey("Then the same values should be obtained back via Get", func() {
				v, err := r.Get("string-test")
				So(err, ShouldBeNil)
				So(string(v), ShouldEqual, "stringy")
			})
		})
	})
}
