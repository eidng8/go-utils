package utils

import (
	lg "log"
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
