package wgolib

import (
	"testing"

	"github.com/changmu/wgolib/wlog"
	"github.com/smartystreets/goconvey/convey"
)

func TestJSONEncode(t *testing.T) {
	convey.Convey("", t, func() {
		convey.Convey("test nil", func() {
			convey.So(JSONEncode(nil), convey.ShouldEqual, "null")
			log := wlog.NewLogger()
			log.Errorf("hello wlog")
		})
	})
}
