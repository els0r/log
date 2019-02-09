package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

// ConsoleLogger provides logging facilities to standard out and standard err. Anything >= Info is written to Stdout, anything below to Stderr
type ConsoleLogger struct{}

var (
	errPrefix   = color.New(color.Bold, color.FgRed).Sprintf("[ ERR]")
	warnPrefix  = color.New(color.Bold, color.FgYellow).Sprintf("[WARN]")
	infoPrefix  = color.New(color.Bold, color.FgWhite).Sprintf("[INFO]")
	debugPrefix = color.New(color.Bold, color.FgGreen).Sprintf("[DEBG]")
)

// NewConsoleLogger creates a new ConsoleLogger. Options allow to set the level. It is `Info` by default
func NewConsoleLogger() (*ConsoleLogger, error) {
	return &ConsoleLogger{}, nil
}

// Debug prints messages with level DEBUG to Stdout
func (c *ConsoleLogger) Debug(args ...interface{}) {
	c.Debugf("%s", fmt.Sprint(args...))
}

// Debugf is formatted Debug
func (c *ConsoleLogger) Debugf(format string, args ...interface{}) {
	c.writeLine(os.Stdout, debugPrefix, fmt.Sprintf(format, args...))
}

// Error prints messages with level ERR to Stderr
func (c *ConsoleLogger) Error(args ...interface{}) {
	c.Errorf("%s", fmt.Sprint(args...))
}

// Errorf is formatted Error
func (c *ConsoleLogger) Errorf(format string, args ...interface{}) {
	c.writeLine(os.Stderr, errPrefix, fmt.Sprintf(format, args...))
}

// Info prints messages with level INFO to Stdout
func (c *ConsoleLogger) Info(args ...interface{}) {
	c.Infof("%s", fmt.Sprint(args...))
}

// Infof is formatted Info
func (c *ConsoleLogger) Infof(format string, args ...interface{}) {
	c.writeLine(os.Stdout, infoPrefix, fmt.Sprintf(format, args...))
}

// Warn prints messages with level WARN to Stderr
func (c *ConsoleLogger) Warn(args ...interface{}) {
	c.Warnf("%s", fmt.Sprint(args...))
}

// Warnf is formatted Warn
func (c *ConsoleLogger) Warnf(format string, args ...interface{}) {
	c.writeLine(os.Stderr, warnPrefix, fmt.Sprintf(format, args...))
}

// helper function to filter out deselected criticality levels
func (c *ConsoleLogger) writeLine(output io.Writer, prefix, msg string) {
	fmt.Fprintf(output, "%s %s %s\n", prefix, time.Now().Local().Format("Mon Jan 2 15:04:05 2006"), msg)
}

// Close is a no-op function to fulfil the Logger interface
func (c *ConsoleLogger) Close() error { return nil }
