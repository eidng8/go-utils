package utils

import (
	"fmt"
	lg "log"
	"strings"
)

type TaggedLogger interface {
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	PanicIfError(err error)
}

type SimpleTaggedLog struct {
	logger *lg.Logger
	Debug  bool
}

func NewLogger() SimpleTaggedLog {
	return WrapLogger(lg.Default(), false)
}

func NewDebugLogger() SimpleTaggedLog {
	return WrapLogger(lg.Default(), true)
}

func WrapLogger(logger *lg.Logger, debug bool) SimpleTaggedLog {
	return SimpleTaggedLog{logger: logger, Debug: debug}
}

func (l SimpleTaggedLog) Debugf(format string, args ...interface{}) {
	if l.Debug {
		l.logger.Printf("[DEBUG] "+format, args...)
	}
}

func (l SimpleTaggedLog) Errorf(format string, args ...interface{}) {
	l.logger.Printf("[ERROR] "+format, args...)
}

func (l SimpleTaggedLog) Infof(format string, args ...interface{}) {
	l.logger.Printf("[INFO] "+format, args...)
}

func (l SimpleTaggedLog) Panicf(format string, args ...interface{}) {
	l.logger.Panicf("[PANIC] "+format, args...)
}

func (l SimpleTaggedLog) PanicIfError(err error) {
	if err != nil {
		l.logger.Panicf("[PANIC] " + err.Error())
	}
}

type StringTaggedLogger struct {
	sb *strings.Builder
}

func NewStringTaggedLogger() StringTaggedLogger {
	return StringTaggedLogger{sb: &strings.Builder{}}
}

func (m StringTaggedLogger) Debugf(format string, args ...interface{}) {
	m.sb.WriteString(fmt.Sprintf("[DEBUG] "+format+"\n", args...))
}

func (m StringTaggedLogger) Errorf(format string, args ...interface{}) {
	m.sb.WriteString(fmt.Sprintf("[ERROR] "+format+"\n", args...))
}

func (m StringTaggedLogger) Infof(format string, args ...interface{}) {
	m.sb.WriteString(fmt.Sprintf("[INFO] "+format+"\n", args...))
}

func (m StringTaggedLogger) Panicf(format string, args ...interface{}) {
	s := fmt.Sprintf("[PANIC] "+format+"\n", args...)
	m.sb.WriteString(s)
	panic(s)
}

func (m StringTaggedLogger) PanicIfError(err error) {
	if err != nil {
		m.Panicf(err.Error())
	}
}

var _ TaggedLogger = SimpleTaggedLog{}
var _ TaggedLogger = StringTaggedLogger{}
