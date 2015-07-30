// Package levelLog provides ...
package levelLog

import (
	"fmt"
)

func Trace(format string, arg ...interface{}) {
	logger.Trace.Output(2, fmt.Sprintf(format, arg...))
}

func Info(format string, arg ...interface{}) {
	logger.Info.Output(2, fmt.Sprintf(format, arg...))
}

func Warning(format string, arg ...interface{}) {
	logger.Warning.Output(2, fmt.Sprintf(format, arg...))
}

func Error(format string, arg ...interface{}) {
	logger.Error.Output(2, fmt.Sprintf(format, arg...))
}