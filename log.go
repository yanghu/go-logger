// Package levelLog provides a logger with different levels
package levelLog

import (
	"log"
)

type LogLevel int

const (
	LevelTrace LogLevel = 1 << iota
	LevelInfo
	LevelWarning
	LevelError
)

type levelLog struct {
	Level       LogLevel
	trace       *log.Logger
	info        *log.Logger
	warning     *log.Logger
	errorLogger *log.Logger
}

// a global logger provide singleton
var logger levelLog
