package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

type (
	logger struct {
		started          bool
		ended            bool
		time             *time.Time
		logLevel         logLevel
		verbose          bool
		file             bool
		std              bool
		filePath         string
		useCollores      bool
		broadCastChannel chan interface{}
		duration         *time.Duration
		prefixString     string
	}

	Logger interface {
		// Start the logger operation
		doLog(level logLevel, a ...interface{})

		End()
		Prefix(format ...string) *logger
		GetCaller() *logger

		// Log logs a message at log level
		Logln(a ...interface{}) LogResult
		// Logln logs a message at log level to new line
		Log(a ...interface{}) LogResult
		// LogF logs a message at log level with string formater
		LogF(format string, a ...interface{}) LogResult

		// Alert logs a message at log level
		Alertln(a ...interface{}) LogResult
		// Alertln logs a message at log level to new line
		Alert(a ...interface{}) LogResult
		// AlertF logs a message at log level with string formater
		AlertF(format string, a ...interface{}) LogResult

		// Error logs a message at log level
		Error(a ...interface{}) LogResult
		// Errorln logs a message at log level to new line
		Errorln(a ...interface{}) LogResult
		// ErrorF logs a message at log level with string formater
		ErrorF(format string, a ...interface{}) LogResult

		// Highlight logs a message at log level
		Highlight(a ...interface{}) LogResult
		// Highlightln logs a message at log level to new line
		Highlightln(a ...interface{}) LogResult
		// HighlightF logs a message at log level with string formater
		HighlightF(format string, a ...interface{}) LogResult

		// Inform logs a message at log level
		Inform(a ...interface{}) LogResult
		// Informln logs a message at log level to new line
		Informln(a ...interface{}) LogResult
		// InformF logs a message at log level with string formater
		InformF(format string, a ...interface{}) LogResult

		// Trace logs a message at log level
		Trace(a ...interface{}) LogResult
		// Traceln logs a message at log level to new line
		Traceln(a ...interface{}) LogResult
		// TraceF logs a message at log level with string formater
		TraceF(format string, a ...interface{}) LogResult

		// Warn logs a message at log level
		Warn(a ...interface{}) LogResult
		// Warnln logs a message at log level to new line
		Warnln(a ...interface{}) LogResult
		// WarnF logs a message at log level with string formater
		WarnF(format string, a ...interface{}) LogResult
	}
	logResult struct {
		logger *logger
	}
	LogResult interface {
		TraceStack()
	}
	logLevel      int
	LoggerOptions struct {
		LogLevel    logLevel
		Verbose     bool
		File        bool
		FilePath    string
		Std         bool
		UseCollores bool
		Writer      io.Writer
	}
)

const (
	Nothing logLevel = iota
	Alert
	Error
	Warn
	Highlight
	Inform
	Log
	Trace
)

func CreateLogger(LoggerOpts *LoggerOptions) Logger {
	c := make(chan interface{})
	if LoggerOpts.LogLevel > Trace {
		LoggerOpts.LogLevel = Trace
	}
	// Cuz alert and error are in 1 level
	if LoggerOpts.LogLevel == Alert {
		LoggerOpts.LogLevel = Error
	}
	l := &logger{
		logLevel:         LoggerOpts.LogLevel,
		verbose:          LoggerOpts.Verbose,
		file:             LoggerOpts.File,
		filePath:         LoggerOpts.FilePath,
		std:              LoggerOpts.Std,
		useCollores:      LoggerOpts.UseCollores,
		broadCastChannel: c,
	}
	l.started = true
	t := time.Now()
	l.time = &t
	go l.wStd(l.broadCastChannel)
	l.broadCastChannel <- "BEGIN :" + "\n"
	return l
}

func (l *logger) Prefix(format ...string) *logger {
	l.prefixString = strings.Join(format, ": ")
	return l
}

func (l *logger) End() {
	l.ended = true
	d := time.Since(
		*l.time,
	)
	l.duration = &d
	l.broadCastChannel <- "\nEND : " + l.duration.String() + "\n"
	l.broadCastChannel <- nil // to gratefull close the channel
}

func (l *logger) getCaller() string {
	fpcs := make([]uintptr, 1)

	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return ""
	}
	caller := runtime.FuncForPC(fpcs[0] - 2)
	return caller.Name() + "()"
}

func (l *logger) doLog(level logLevel, a ...interface{}) {
	if l.prefixString != "" {
		l.broadCastChannel <- l.prefixString + ": "
	}
	for _, v := range a {
		l.broadCastChannel <- v
	}
}

func (lr *logResult) TraceStack() {
	stackSlice := make([]byte, 512)
	s := runtime.Stack(stackSlice, false)
	lr.logger.LogF("\n%s", stackSlice[0:s])
}

func (l *logger) GetCaller() *logger {
	l.LogF("%s :: ", l.getCaller())
	return l
}

func (l *logger) wStd(c chan interface{}) {
	if l.std {
		for msg := range c {
			if msg != nil {
				fmt.Fprint(
					os.Stdout,
					msg,
				)
			} else {
				close(c)
			}
		}

	}
}
