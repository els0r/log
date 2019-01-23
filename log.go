package log

import (
	"fmt"
	"strings"
	"sync"
)

// Logger specifies the general interface for plugging in loggers from third-party logging packages
type Logger interface {
	Debug(msg string)
	Debugf(format string, args ...interface{})
	Err(msg string)
	Errf(format string, args ...interface{})
	Info(msg string)
	Infof(format string, args ...interface{})
	Warn(msg string)
	Warnf(format string, args ...interface{})

	Close() error // some logger implementations will need this
}

// Log wraps an attached Logger
type Log struct {
	sync.Mutex
	l     Logger
	level Level
}

// Level is a wrapper for the criticality level
type Level int

// Enumeration for loggers using different criticality levels (such as syslog)
const (
	ERR Level = iota + 1
	WARN
	INFO
	DEBUG
)

// LoggerImplementation enumerates the Logger implementations provided by this package
type LoggerImplementation int

const (
	// DevNull makes sure that messages go nowhere
	DevNull LoggerImplementation = iota + 1

	// Console writes output to terminal
	Console

//	Syslog uses the standard syslog facilities for logging
//	Json encodes the log messages as machine-readable json structures
)

// GetLoggerImplementation returns the enumeration value for a logger implementation provided as string. Both lower and upper case work. In case the string does not specify a valid implementation, 0 is returned.
func GetLoggerImplementation(s string) LoggerImplementation {
	s = strings.ToUpper(s)
	if i, ok := loggerFromStrings[s]; ok {
		return i
	}
	return 0
}

// String returns the string representation of the LoggerImplemenation
func (l LoggerImplementation) String() string {
	if !(0 <= int(l) && int(l) < len(loggerImplementationToStrings)) {
		l = LoggerImplementation(0)
	}
	return loggerImplementationToStrings[l]
}

// GetLevel returns the enumeration value for a level provided as string. Both lower and upper case work. In case the string does not specify a valid level, 0 is returned.
func GetLevel(s string) Level {
	s = strings.ToUpper(s)
	if level, ok := fromStrings[s]; ok {
		return level
	}
	return 0
}

// String returns the string representation of the Level
func (l Level) String() string {
	if !(0 <= int(l) && int(l) < len(toStrings)) {
		l = Level(0)
	}
	return toStrings[l]
}

// NewFromString creates a new log object based on the string identifiers for a supported logger implementation. Options can be applied.
func NewFromString(id string, opts ...Option) (*Log, error) {
	var err error

	// default level is info
	l := &Log{level: INFO}

	// call the constructor and look if the the implemenation is support
	switch GetLoggerImplementation(id) {
	case DevNull:
		l.l, err = NewDevNullLogger()
	case Console:
		l.l, err = NewConsoleLogger()
		// case Syslog:
	//	l.l, err = NewSyslogLogger()
	default:
		return nil, fmt.Errorf("Unable to find logger implementation '%s'", id)
	}
	if err != nil {
		return nil, err
	}

	// apply options
	for _, opt := range opts {
		err = opt(l)
		if err != nil {
			return nil, err
		}
	}

	return l, nil
}

// New creates a new log object. Options allow to set the destination of the log output
func New(opts ...Option) (*Log, error) {
	var err error

	// by default, Log writes to the console with level INFO
	l := &Log{level: INFO}
	l.l, err = NewConsoleLogger()
	if err != nil {
		return nil, err
	}

	// apply options
	for _, opt := range opts {
		err = opt(l)
		if err != nil {
			return nil, err
		}
	}

	return l, nil
}

// Debug prints messages on level DEBUG
func (l *Log) Debug(msg string) {
	l.Lock()
	defer l.Unlock()

	if l.ignoreLine(DEBUG) {
		return
	}

	l.l.Debug(msg)
}

// Debugf is formatted Debug
func (l *Log) Debugf(format string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.ignoreLine(DEBUG) {
		return
	}

	l.l.Debugf(format, args...)
}

// Err prints messages on level ERROR
func (l *Log) Err(msg string) {
	l.Lock()
	defer l.Unlock()

	if l.ignoreLine(ERR) {
		return
	}

	l.l.Err(msg)
}

// Errf is formatted Err
func (l *Log) Errf(format string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.ignoreLine(ERR) {
		return
	}

	l.l.Errf(format, args...)
}

// Info prints messages on level INFO
func (l *Log) Info(msg string) {
	l.Lock()
	defer l.Unlock()

	if l.ignoreLine(INFO) {
		return
	}

	l.l.Info(msg)
}

// Infof is formatted Info
func (l *Log) Infof(format string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.ignoreLine(INFO) {
		return
	}

	l.l.Infof(format, args...)
}

// Warn prints messages on level WARN
func (l *Log) Warn(msg string) {
	l.Lock()
	defer l.Unlock()

	if l.ignoreLine(WARN) {
		return
	}

	l.l.Warn(msg)
}

// Warnf is formatted Warn
func (l *Log) Warnf(format string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.ignoreLine(WARN) {
		return
	}

	l.l.Warnf(format, args...)
}

// Close is used to close any open objects the underlying logger may use
func (l *Log) Close() error {
	l.Lock()
	defer l.Unlock()

	return l.l.Close()
}

const unknown = "UNKNOWN"

// helper to convert a string to a LoggerImplementation
var loggerFromStrings = map[string]LoggerImplementation{
	"DEVNULL": DevNull,
	"CONSOLE": Console,
	//	"SYSLOG":  Syslog,
}

// helper for string method
var loggerImplementationToStrings = [...]string{
	unknown,
	"DEVNULL",
	"CONSOLE",
	//"SYSLOG",
}

// helper for string method
var toStrings = [...]string{
	unknown,
	"ERR",
	"WARN",
	"INFO",
	"DEBUG",
}

// helper to convert a string to a Level
var fromStrings = map[string]Level{
	"ERR":   ERR,
	"WARN":  WARN,
	"INFO":  INFO,
	"DEBUG": DEBUG,
}

func (l *Log) ignoreLine(level Level) bool {
	return l.level < level
}

/*
if level > l.level ignore
if level == l.level print
if level < l.level print
*/