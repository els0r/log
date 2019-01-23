package log

// DevNullLogger provides the no-op logger
type DevNullLogger struct{}

// NewDevNullLogger creates a new DevNullLogger
func NewDevNullLogger() (*DevNullLogger, error) {
	return &DevNullLogger{}, nil
}

// Debug does nothing
func (d *DevNullLogger) Debug(msg string) { return }

// Debugf does nothing
func (d *DevNullLogger) Debugf(format string, args ...interface{}) { return }

// Err does nothing
func (d *DevNullLogger) Err(msg string) { return }

// Errf does nothing
func (d *DevNullLogger) Errf(format string, args ...interface{}) { return }

// Info does nothing
func (d *DevNullLogger) Info(msg string) { return }

// Infof does nothing
func (d *DevNullLogger) Infof(format string, args ...interface{}) { return }

// Warn does nothing
func (d *DevNullLogger) Warn(msg string) { return }

// Warnf does nothing
func (d *DevNullLogger) Warnf(format string, args ...interface{}) { return }

// Close does nothing
func (d *DevNullLogger) Close() error { return nil }
