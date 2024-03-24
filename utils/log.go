package utils

import (
	"log"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

var logLevelPrefix = []string{"DEBUG", "INFO", "WARNING", "ERROR"}

type Log struct{}

func NewLogObject() Log {
	return Log{}
}

func (l Log) SetLogLevelPrefix(prefixes []string) {
	logLevelPrefix = prefixes
}

func (l Log) Log(level LogLevel, message string) {
	if level < DebugLevel || level > ErrorLevel {
		log.Fatalf("Invalid log level: %d", level)
	}

	log.Printf("[%s]: %s", logLevelPrefix[level], message)
}

func (l Log) Debug(message string) {
	l.Log(DebugLevel, message)
}

func (l Log) Info(message string) {
	l.Log(InfoLevel, message)
}

func (l Log) Warning(message string) {
	l.Log(WarningLevel, message)
}

func (l Log) Error(message string) {
	l.Log(ErrorLevel, message)
}
