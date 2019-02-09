package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// JSONLogger uses the logrus logging framework to write JSON formatted log messages
type JSONLogger struct{ *logrus.Logger }

// NewJSONLogger returns a JSON logger writing to standard out by default
func NewJSONLogger() *JSONLogger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})

	// by default, JSON is written to Stdout
	lr.SetOutput(os.Stdout)

	return &JSONLogger{lr}
}

// Close is a no-op for the JSON logger
func (j *JSONLogger) Close() error {
	return nil
}
