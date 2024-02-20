package utils

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStruct2map(t *testing.T) {
	convey.Convey("test regular struct to map", t, func() {
		type User struct {
			Name string
			Id   int
		}

		u := User{Name: "wo", Id: 1}
		m, err := Struct2map(u)
		So(err, ShouldBeNil)
		val, _ := m["Name"]
		So(val, ShouldEqual, "wo")

	})

	convey.Convey("test struct to map with json  tag", t, func() {
		type User struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		}

		u := User{Name: "wo", Id: 1}
		m, err := Struct2map(u)
		So(err, ShouldBeNil)
		val, _ := m["name"]
		So(val, ShouldEqual, "wo")

	})

}
