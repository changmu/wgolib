package wlog

import (
	"testing"
	"time"

	"github.com/changmu/wgolib/wlog/consts"
	"github.com/changmu/wgolib/wlog/opt"
	log "github.com/sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
)

func TestNewLogger(t *testing.T) {
	convey.Convey("", t, func() {
		convey.Convey("test log", func() {
			InitLogger(
				opt.WithLogLevel(consts.LogLevelTrace),
				opt.WithFileName("test"),
				opt.WithLogType(consts.LogTypeLogrus),
				opt.WithKeepDays(2),
			)

			log.Tracef("hello logrus")
			time.Sleep(10 * time.Millisecond)
			log.Debugf("hello logrus")
			log.Infof("hello logrus")
			time.Sleep(20 * time.Millisecond)
			log.Errorf("hello logrus")
		})
	})
}
