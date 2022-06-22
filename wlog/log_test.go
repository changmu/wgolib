package wlog

import (
	"testing"
	"time"

	"github.com/changmu/wgolib/wlog/consts"
	"github.com/changmu/wgolib/wlog/opt"
	"github.com/smartystreets/goconvey/convey"
)

func TestNewLogger(t *testing.T) {
	convey.Convey("", t, func() {
		convey.Convey("test logger", func() {
			logger = NewLogger(
				opt.WithLogLevel(consts.LogLevelTrace),
				opt.WithFileName("test"),
				opt.WithLogType(consts.LogTypeLogrus),
				opt.WithKeepDays(2),
			)

			logger.Tracef("hello logrus")
			time.Sleep(10 * time.Millisecond)
			logger.Debugf("hello logrus")
			logger.Infof("hello logrus")
			time.Sleep(20 * time.Millisecond)
			logger.Errorf("hello logrus")
		})
	})
}
