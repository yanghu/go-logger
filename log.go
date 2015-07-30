// Package levelLog provides a logger with different levels
package levelLog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type LogLevel int

const (
	LevelTraceOnly LogLevel = 1 << iota
	LevelInfoOnly
	LevelWarningOnly
	LevelErrorOnly
	LevelTrace   = LevelTraceOnly | LevelInfoOnly | LevelWarningOnly | LevelErrorOnly
	LevelInfo    = LevelInfoOnly | LevelWarningOnly | LevelErrorOnly
	LevelWarning = LevelWarningOnly | LevelErrorOnly
	LevelError   = LevelErrorOnly
)

type levelLog struct {
	Level   LogLevel
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

// a global logger provide singleton
var logger levelLog

func turnOnLogging(level LogLevel, fileHandle io.Writer) {
	traceHandle := ioutil.Discard
	infoHandle := ioutil.Discard
	warningHandle := ioutil.Discard
	errorHandle := ioutil.Discard

	if level&LevelTraceOnly != 0 {
		if fileHandle != nil {
			traceHandle = io.MultiWriter(fileHandle, os.Stdout)
		} else {
			traceHandle = os.Stdout
		}
	}

	if level&LevelInfoOnly != 0 {
		if fileHandle != nil {
			infoHandle = io.MultiWriter(fileHandle, os.Stdout)
		} else {
			infoHandle = os.Stdout
		}
	}

	if level&LevelWarningOnly != 0 {
		if fileHandle != nil {
			warningHandle = io.MultiWriter(fileHandle, os.Stdout)
		} else {
			warningHandle = os.Stdout
		}
	}

	if level&LevelErrorOnly != 0 {
		if fileHandle != nil {
			errorHandle = io.MultiWriter(fileHandle, os.Stderr)
		} else {
			errorHandle = os.Stderr
		}
	}

	logger.Trace = log.New(traceHandle, "TRACE:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info = log.New(infoHandle, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warning = log.New(warningHandle, "WARNING:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = log.New(errorHandle, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Level = level
}
