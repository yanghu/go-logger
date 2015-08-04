package logger

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

func Error(err error, format string, arg ...interface{}) {
	logger.Error.Output(2, fmt.Sprintf("%s:%s", err, fmt.Sprintf(format, arg...)))
}

func TraceRaw(output string) {
	logger.Trace.Output(2, output)
}

func InfoRaw(output string) {
	logger.Info.Output(2, output)
}

func WarningRaw(output string) {
	logger.Warning.Output(2, output)
}

func ErrorRaw(output string) {
	logger.Error.Output(2, output)
}
