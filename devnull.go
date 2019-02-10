package log

// DevNullLogger provides the no-op logger
type DevNullLogger struct{}

// NewDevNullLogger creates a new DevNullLogger
func NewDevNullLogger() *DevNullLogger {
	return &DevNullLogger{}
}

// Debug does nothing
func (d *DevNullLogger) Debug(args ...interface{}) { return }

// Debugf does nothing
func (d *DevNullLogger) Debugf(format string, args ...interface{}) { return }

// Error does nothing
func (d *DevNullLogger) Error(args ...interface{}) { return }

// Errorf does nothing
func (d *DevNullLogger) Errorf(format string, args ...interface{}) { return }

// Info does nothing
func (d *DevNullLogger) Info(args ...interface{}) { return }

// Infof does nothing
func (d *DevNullLogger) Infof(format string, args ...interface{}) { return }

// Warn does nothing
func (d *DevNullLogger) Warn(args ...interface{}) { return }

// Warnf does nothing
func (d *DevNullLogger) Warnf(format string, args ...interface{}) { return }

// Close does nothing
func (d *DevNullLogger) Close() error { return nil }
