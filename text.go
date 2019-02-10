package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

// TextLogger provides logging facilities to standard out and standard err. By default, anything >= Info is written to Stdout, anything below to Stderr. Other destinations can be set via the logger's options.
type TextLogger struct {
	wOut, wErr io.Writer // message routing for info and error messages
}

var (
	errPrefix   = color.New(color.Bold, color.FgRed).Sprintf("[ ERR]")
	warnPrefix  = color.New(color.Bold, color.FgYellow).Sprintf("[WARN]")
	infoPrefix  = color.New(color.Bold, color.FgWhite).Sprintf("[INFO]")
	debugPrefix = color.New(color.Bold, color.FgGreen).Sprintf("[DEBG]")
)

// TextLoggerOption is for functional options applied to the TextLogger
type TextLoggerOption func(*TextLogger)

// WithTextOutputRouting sets the destination for informational messages as well as for errors.
func WithTextOutputRouting(info io.Writer, err io.Writer) TextLoggerOption {
	return func(t *TextLogger) {
		t.wOut, t.wErr = info, err
	}
}

// NewTextLogger creates a new TextLogger.
func NewTextLogger(opts ...TextLoggerOption) (*TextLogger, error) {
	t := &TextLogger{os.Stdout, os.Stderr}

	// apply options
	for _, opt := range opts {
		opt(t)
	}

	return t, nil
}

// Debug prints messages with level DEBUG to Stdout
func (t *TextLogger) Debug(args ...interface{}) {
	t.Debugf("%s", fmt.Sprint(args...))
}

// Debugf is formatted Debug
func (t *TextLogger) Debugf(format string, args ...interface{}) {
	t.writeLine(t.wOut, debugPrefix, fmt.Sprintf(format, args...))
}

// Error prints messages with level ERR to Stderr
func (t *TextLogger) Error(args ...interface{}) {
	t.Errorf("%s", fmt.Sprint(args...))
}

// Errorf is formatted Error
func (t *TextLogger) Errorf(format string, args ...interface{}) {
	t.writeLine(t.wErr, errPrefix, fmt.Sprintf(format, args...))
}

// Info prints messages with level INFO to Stdout
func (t *TextLogger) Info(args ...interface{}) {
	t.Infof("%s", fmt.Sprint(args...))
}

// Infof is formatted Info
func (t *TextLogger) Infof(format string, args ...interface{}) {
	t.writeLine(t.wOut, infoPrefix, fmt.Sprintf(format, args...))
}

// Warn prints messages with level WARN to Stderr
func (t *TextLogger) Warn(args ...interface{}) {
	t.Warnf("%s", fmt.Sprint(args...))
}

// Warnf is formatted Warn
func (t *TextLogger) Warnf(format string, args ...interface{}) {
	t.writeLine(t.wErr, warnPrefix, fmt.Sprintf(format, args...))
}

// helper function to filter out deselected criticality levels
func (t *TextLogger) writeLine(output io.Writer, prefix, msg string) {
	fmt.Fprintf(output, "%s %s %s\n", prefix, time.Now().Local().Format("Mon Jan 2 15:04:05 2006"), msg)
}

// Close is a no-op function to fulfil the Logger interface
func (t *TextLogger) Close() error { return nil }
