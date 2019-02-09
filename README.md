# log [![GoDoc](https://godoc.org/github.com/els0r/log?status.svg)](https://godoc.org/github.com/els0r/log) [![Go Report Card](https://goreportcard.com/badge/github.com/els0r/log)](https://goreportcard.com/report/github.com/els0r/log)
Simple logging interface and common loggers for injection into other packages
===========

This package supplies a thread-safe logger, a simple logging interface and commonly used logging destinations. Currently supported implementations are

* Console (Stdout/Stderr to terminal)
* DevNull (does nothing)
* JSON (logs JSON messages to Stdout)
* Syslog (*TODO: coming soon*)

Quick Start
------------

To use the logging interface do

```
go get github.com/els0r/log/...
```

Example on how to use the package in your code
```golang
package main

import(
  log "github.com/els0r/log"
)

func main() {
  l := log.New(WithLevel(log.INFO))

  // Do someting

  l.Infof("We have done %d things today", 10)
  l.Error("This is an error")
}
```
