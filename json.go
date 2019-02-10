package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// JSONLogger uses the logrus logging framework to write JSON formatted log messages
type JSONLogger struct {
	*logrus.Logger
	out io.Writer
}

// JSONLoggerOption is for functional options applied to the JSONLogger
type JSONLoggerOption func(*JSONLogger)

// WithJSONOutputRouting sets the destination for the encoded messages
func WithJSONOutputRouting(w io.Writer) JSONLoggerOption {
	return func(j *JSONLogger) {
		j.out = w
	}
}

// NewJSONLogger returns a JSON logger writing to standard out by default
func NewJSONLogger(opts ...JSONLoggerOption) *JSONLogger {

	j := &JSONLogger{logrus.New(), os.Stdout} // by default, JSON is written to Stdout
	j.SetFormatter(&logrus.JSONFormatter{})

	// apply options
	for _, opt := range opts {
		opt(j)
	}

	j.SetOutput(j.out)

	return j
}

// Close is a no-op for the JSON logger
func (j *JSONLogger) Close() error {
	return nil
}
