package log

import (
	"fmt"
	"log/syslog"
	"os"
)

// Syslogger implements the Logger interface for Syslog
type Syslogger struct {
	w    *syslog.Writer
	prio syslog.Priority
	tag  string
}

// SyslogOption is a type for functional options applied to the Syslogger
type SyslogOption func(*Syslogger)

// WithSyslogPriority sets the priority. This only makes sense if the logger is used standalone and not encapsulated by a Logger that deals with levels already
func WithSyslogPriority(prio syslog.Priority) SyslogOption {
	return func(s *Syslogger) {
		s.prio = prio
		return
	}
}

// WithSyslogTag allows to set the tag to somenthing other than the program's name (default behavior)
func WithSyslogTag(tag string) SyslogOption {
	return func(s *Syslogger) {
		s.tag = tag
		return
	}
}

// NewSyslogger ...
func NewSyslogger(opts ...SyslogOption) (*Syslogger, error) {
	var (
		err error
		s   = &Syslogger{prio: syslog.LOG_DEBUG, tag: os.Args[0]} // the chattiest priority is taken as the overlying logger will control severity levels. If the logger is used standalone, the option should be used to set the priority
	)

	// open new syslog writer
	s.w, err = syslog.New(s.prio, s.tag)
	if err != nil {
		return nil, err
	}

	// apply options
	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}

// Debug prints messages with level LOG_DEBUG to syslog
func (s *Syslogger) Debug(args ...interface{}) {
	s.Debugf("%s", args...)
}

// Debugf is formatted Debug
func (s *Syslogger) Debugf(format string, args ...interface{}) {
	s.writeMessage(s.w.Debug, format, args...)
}

// Error prints messages with level LOG_ERR to syslog
func (s *Syslogger) Error(args ...interface{}) {
	s.Errorf("%s", args...)
}

// Errorf is formatted Error
func (s *Syslogger) Errorf(format string, args ...interface{}) {
	s.writeMessage(s.w.Err, format, args...)
}

// Info prints messages with level LOG_INFO to syslog
func (s *Syslogger) Info(args ...interface{}) {
	s.Infof("%s", args...)
}

// Infof is formatted Info
func (s *Syslogger) Infof(format string, args ...interface{}) {
	s.writeMessage(s.w.Info, format, args...)
}

// Warn prints messages with level LOG_WARN to syslog
func (s *Syslogger) Warn(args ...interface{}) {
	s.Warnf("%s", args...)
}

// Warnf is formatted Warn
func (s *Syslogger) Warnf(format string, args ...interface{}) {
	s.writeMessage(s.w.Warning, format, args...)
}

// Close closes the connection to the local syslog daemon
func (s *Syslogger) Close() error {
	return s.w.Close()
}

func (s *Syslogger) writeMessage(writeFunc func(string) error, format string, args ...interface{}) {
	var err error

	// attempt reconnect in case logger is nil
	if s.w == nil {
		s.w, err = syslog.New(s.prio, s.tag)
		if err != nil {
			return
		}
	}

	// write message
	err = writeFunc(fmt.Sprintf(format, args...))

	if err != nil {
		s.w, _ = syslog.New(s.prio, s.tag)
	}
}
