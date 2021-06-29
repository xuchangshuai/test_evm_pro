package example

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestExample(t *testing.T) {
	Convey("TestTestExampleEqual a == b ", t, func() {
		a := 1
		b := 1
		So(a, ShouldEqual, b)
	})
}
