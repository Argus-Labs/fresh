package runner

import (
	"fmt"
	logPkg "log"
	"time"

	"github.com/mattn/go-colorable"
)

type logFunc func(string, ...interface{})

var logger = logPkg.New(colorable.NewColorableStderr(), "", 0)

func newLogFunc() func(string, ...interface{}) {
	black := fmt.Sprintf("\033[%sm", colors["grey"])
	green := fmt.Sprintf("\033[%sm", colors["green"])
	reset := fmt.Sprintf("\033[%sm", colors["reset"])

	return func(format string, v ...interface{}) {
		now := time.Now()
		timeString := now.Format(time.Kitchen)
		format = fmt.Sprintf("%s%s%s INF %s%s", black, timeString, green, reset, format)
		logger.Printf(format, v...)
	}
}

func newLogFuncPlain() func(string, ...interface{}) {
	return func(format string, v ...interface{}) {
		logger.Printf(format, v...)
	}
}

func fatal(err error) {
	logger.Fatal(err)
}

type appLogWriter struct{}

func (a appLogWriter) Write(p []byte) (n int, err error) {
	appLog(string(p))

	return len(p), nil
}
