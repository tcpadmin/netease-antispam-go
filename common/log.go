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

type LogFieldDto struct {
	Key   string
	Value interface{}
}

func LogField(key string, val interface{}) LogFieldDto {
	return LogFieldDto{key, val}
}

// ILogger 实现该接口的日志；类似 gorm
type ILogger interface {
	SetLogLevel(LogLevel)
	Info(context.Context, string, ...LogFieldDto)
	Warn(context.Context, string, ...LogFieldDto)
	Error(context.Context, string, ...LogFieldDto)
}

type DefaultLog struct {
	level LogLevel
}

func (d *DefaultLog) SetLogLevel(level LogLevel) {
	d.level = level
}

func (d *DefaultLog) Info(ctx context.Context, s string, i ...LogFieldDto) {
	if d.level >= Info {
		data := make([]interface{}, 0, len(i)*2)
		for _, item := range i {
			data = append(data, item.Key, item.Value)
		}
		fmt.Println(s, data)
	}
}

func (d *DefaultLog) Warn(ctx context.Context, s string, i ...LogFieldDto) {
	if d.level >= Warn {
		data := make([]interface{}, 0, len(i)*2)
		for _, item := range i {
			data = append(data, item.Key, item.Value)
		}
		fmt.Println(s, data)
	}
}

func (d *DefaultLog) Error(ctx context.Context, s string, i ...LogFieldDto) {
	if d.level >= Error {
		data := make([]interface{}, 0, len(i)*2)
		for _, item := range i {
			data = append(data, item.Key, item.Value)
		}
		fmt.Println(s, data)
	}
}
