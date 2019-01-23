package log

import "fmt"

// Option is any function operating on or modifying a Log object. A function with this signature can be passed to the constructors
type Option func(*Log) error

// WithOutput sets a logger to something other than the default output. If possible, the logging level of `out` is inherited
func WithOutput(out Logger) Option {
	return func(l *Log) error {

		// do a type assertion to check if the Logger supports log levels. If so, inherit the log level
		if o, ok := out.(*Log); ok {
			l.level = o.level
		}

		l.l = out
		return nil
	}
}

// WithLevel sets the log level of the logger
func WithLevel(level Level) func(*Log) error {
	return func(l *Log) error {
		// check if level is in range
		if !(ERR <= level && level <= DEBUG) {
			return fmt.Errorf("Unknown log level provided")
		}
		l.level = level

		return nil
	}
}
