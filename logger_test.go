package utils

import (
	"bytes"
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupLoggerTest() (*SimpleTaggedLog, *bytes.Buffer) {
	var buf bytes.Buffer
	logger := WrapLogger(log.New(io.Writer(&buf), "", 0), true)
	return &logger, &buf
}

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	assert.Equal(t, log.Default(), logger.logger)
	assert.False(t, logger.Debug)
}

func TestNewDebugLogger(t *testing.T) {
	logger := NewDebugLogger()
	assert.Equal(t, log.Default(), logger.logger)
	assert.True(t, logger.Debug)
}

func TestDebugf(t *testing.T) {
	logger, buf := setupLoggerTest()
	logger.Debugf("test %d", 1)
	assert.Equal(t, "[DEBUG] test 1\n", buf.String())
}

func TestDebugf_Disabled(t *testing.T) {
	var buf bytes.Buffer
	logger := WrapLogger(log.New(io.Writer(&buf), "", 0), false)
	logger.Debugf("test %d", 1)
	assert.Empty(t, buf)
}

func TestErrorf(t *testing.T) {
	logger, buf := setupLoggerTest()
	logger.Errorf("test %d", 1)
	assert.Equal(t, "[ERROR] test 1\n", buf.String())
}

func TestInfof(t *testing.T) {
	logger, buf := setupLoggerTest()
	logger.Infof("test %d", 1)
	assert.Equal(t, "[INFO] test 1\n", buf.String())
}

func TestPanicf(t *testing.T) {
	logger, buf := setupLoggerTest()
	assert.Panics(t, func() { logger.Panicf("test %d", 1) })
	assert.Equal(t, "[PANIC] test 1\n", buf.String())
}

func TestLoggerPanicIfError(t *testing.T) {
	logger, buf := setupLoggerTest()
	assert.Panics(t, func() { logger.PanicIfError(assert.AnError) })
	assert.Equal(t, "[PANIC] "+assert.AnError.Error()+"\n", buf.String())
	assert.NotPanics(t, func() { logger.PanicIfError(nil) })
}
