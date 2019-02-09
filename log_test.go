package log

import (
	"testing"
)

func TestStdoutLogger(t *testing.T) {

	var (
		err error
		c   Logger
		l   *Log
	)

	// test different instantiation methods
	c, err = NewFromString("console", WithLevel(DEBUG))
	if err != nil {
		t.Fatalf("failed to instantiate console logger from string")
	}

	l, err = New(WithLogger(c))
	if err != nil {
		t.Fatalf("failed to instantiate console logger")
	}

	// call the common functions
	l.Debug("Hello debug world")
	l.Debugf("Hello %d debugf world", 1)
	l.Info("Hello info world")
	l.Infof("Hello %d infof world", 2)
	l.Warn("Hello warn world")
	l.Warnf("Hello %d warn world", 3)
	l.Error("Hello error world")
	l.Errorf("Hello %d error world", 4)

	l.Close()

}

func TestNewFromString(t *testing.T) {

	var err error

	// create the logger
	_, err = NewFromString("devnull")
	if err != nil {
		t.Fatalf("failed to instantiate devnull logger via string identifier")
	}

	// instantiate an unsupported logger
	_, err = NewFromString("I_will_Never_Be_Supported_Logger")
	if err == nil {
		t.Fatalf("expected error for invalid logger implementation identifer")
	}
}

func TestDevNullLogger(t *testing.T) {

	l, err := NewDevNullLogger()
	if err != nil {
		t.Fatalf("failed to instantiate devnull logger")

	}

	// call the common functions on the Debug
	l.Debug("Hello debug world")
	l.Debugf("Hello %d debugf world", 1)
	l.Info("Hello info world")
	l.Infof("Hello %d infof world", 2)
	l.Warn("Hello warn world")
	l.Warnf("Hello %d warn world", 3)
	l.Error("Hello error world")
	l.Errorf("Hello %d error world", 4)

	l.Close()
}

func TestJSONLogger(t *testing.T) {

	l := NewJSONLogger()
	if l == nil {
		t.Fatalf("failed to instantiate JSON logger: logger is <nil>")
	}

	// call the common functions on the Debug
	l.Debug("Hello debug world")
	l.Debugf("Hello %d debugf world", 1)
	l.Info("Hello info world")
	l.Infof("Hello %d infof world", 2)
	l.Warn("Hello warn world")
	l.Warnf("Hello %d warn world", 3)
	l.Error("Hello error world")
	l.Errorf("Hello %d error world", 4)

	l.Close()
}

func TestStringMethods(t *testing.T) {

	// test string methods
	level := ERR
	impl := Console

	if level.String() != toStrings[ERR] {
		t.Fatalf("Level String() method failed. Expected: 'ERR'; Got: '%s'", level.String())
	}

	if impl.String() != loggerImplementationToStrings[Console] {
		t.Fatalf("LoggerImplementation String() method failed. Expected: 'CONSOLE'; Got: '%s'", impl.String())
	}

	// create an unknown logger and level
	s := GetLevel("this level is not known").String()
	if s != unknown {
		t.Fatalf("calling String() on unsupported level should return '%s'. Got: '%s'", unknown, s)
	}
	s = GetLoggerImplementation("this logger is not supported").String()
	if s != unknown {
		t.Fatalf("calling String() on unsupported implementation should return '%s'. Got: '%s'", unknown, s)
	}
}
