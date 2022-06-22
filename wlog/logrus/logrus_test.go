package logrus

import (
	"testing"
	"time"

	"github.com/changmu/wgolib/wlog/consts"
	"github.com/changmu/wgolib/wlog/opt"
	"github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	convey.Convey("", t, func() {
		convey.Convey("test logger", func() {
			logger := New(opt.NewLogOpt(
				opt.WithLogLevel(consts.LogLevelTrace),
			))

			logger.Tracef("hello logrus")
			time.Sleep(10 * time.Millisecond)
			logger.Debugf("hello logrus")
			logger.Infof("hello logrus")
			time.Sleep(20 * time.Millisecond)
			logger.Errorf("hello logrus")
		})
	})
}
