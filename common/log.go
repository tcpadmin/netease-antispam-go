package common

import (
	"context"
	"fmt"
)

type LogLevel int

const (
	// Silent silent log level
	Silent LogLevel = iota + 1
	// Error error log level
	Error
	// Warn warn log level
	Warn
	// Info info log level
	Info
)

// ILogger 实现该接口的日志；类似 gorm
type ILogger interface {
	SetLogLevel(LogLevel)
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
}

type DefaultLog struct {
	level LogLevel
}

func (d *DefaultLog) SetLogLevel(level LogLevel) {
	d.level = level
}

func (d *DefaultLog) Info(ctx context.Context, s string, i ...interface{}) {
	if d.level >= Info {
		fmt.Println(s, i)
	}
}

func (d *DefaultLog) Warn(ctx context.Context, s string, i ...interface{}) {
	if d.level >= Warn {
		fmt.Println(s, i)
	}
}

func (d *DefaultLog) Error(ctx context.Context, s string, i ...interface{}) {
	if d.level >= Error {
		fmt.Println(s, i)
	}
}
