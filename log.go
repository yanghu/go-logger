// Package logger provides a logger with different logging level support and redis support.
package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	LevelTraceOnly Level = 1 << iota
	LevelInfoOnly
	LevelWarningOnly
	LevelErrorOnly
	LevelTrace   = LevelTraceOnly | LevelInfoOnly | LevelWarningOnly | LevelErrorOnly
	LevelInfo    = LevelInfoOnly | LevelWarningOnly | LevelErrorOnly
	LevelWarning = LevelWarningOnly | LevelErrorOnly
	LevelError   = LevelErrorOnly
)

type Level int

// LevelLogger stores loggers for each level
type LevelLogger struct {
	LogLevel Level
	Trace    *log.Logger
	Info     *log.Logger
	Warning  *log.Logger
	Error    *log.Logger
}

// a global logger provide singleton service. It is called within actual loggin actions
// like Info(), Error(), etc.
var logger LevelLogger

// It initializes the logger with specified logging level. By default, all
// levels of loggin goes to stdout. if writer is not nil, then the message
// also goes to the writer. It could be a file handle, or a redis writer
// implemented in logger/redis package.
func turnOnLogging(level Level, writer io.Writer) {
	traceHandle := ioutil.Discard
	infoHandle := ioutil.Discard
	warningHandle := ioutil.Discard
	errorHandle := ioutil.Discard

	if level&LevelTraceOnly != 0 {
		if writer != nil {
			traceHandle = io.MultiWriter(writer, os.Stdout)
		} else {
			traceHandle = os.Stdout
		}
	}

	if level&LevelInfoOnly != 0 {
		if writer != nil {
			infoHandle = io.MultiWriter(writer, os.Stdout)
		} else {
			infoHandle = os.Stdout
		}
	}

	if level&LevelWarningOnly != 0 {
		if writer != nil {
			warningHandle = io.MultiWriter(writer, os.Stdout)
		} else {
			warningHandle = os.Stdout
		}
	}

	if level&LevelErrorOnly != 0 {
		if writer != nil {
			errorHandle = io.MultiWriter(writer, os.Stderr)
		} else {
			errorHandle = os.Stderr
		}
	}

	logger.Trace = log.New(traceHandle, "TRACE:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info = log.New(infoHandle, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warning = log.New(warningHandle, "WARNING:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = log.New(errorHandle, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.LogLevel = level
}
