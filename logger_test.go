package utils

import (
	"bytes"
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupLoggerTest() (*SimpleTaggedLog, *bytes.Buffer) {
	var buf bytes.Buffer
	logger := WrapLogger(log.New(io.Writer(&buf), "", 0), true)
	return &logger, &buf
}

func Test_NewLogger(t *testing.T) {
	logger := NewLogger()
	require.Equal(t, log.Default(), logger.logger)
	require.False(t, logger.Debug)
}

func Test_NewDebugLogger(t *testing.T) {
	logger := NewDebugLogger()
	require.Equal(t, log.Default(), logger.logger)
	require.True(t, logger.Debug)
}

func Test_Debugf(t *testing.T) {
	logger, buf := setupLoggerTest()
	logger.Debugf("test %d", 1)
	require.Equal(t, "[DEBUG] test 1\n", buf.String())
}

func Test_Debugf_Disabled(t *testing.T) {
	var buf bytes.Buffer
	logger := WrapLogger(log.New(io.Writer(&buf), "", 0), false)
	logger.Debugf("test %d", 1)
	require.Empty(t, buf)
}

func Test_Errorf(t *testing.T) {
	logger, buf := setupLoggerTest()
	logger.Errorf("test %d", 1)
	require.Equal(t, "[ERROR] test 1\n", buf.String())
}

func Test_Infof(t *testing.T) {
	logger, buf := setupLoggerTest()
	logger.Infof("test %d", 1)
	require.Equal(t, "[INFO] test 1\n", buf.String())
}

func Test_Panicf(t *testing.T) {
	logger, buf := setupLoggerTest()
	require.Panics(t, func() { logger.Panicf("test %d", 1) })
	require.Equal(t, "[PANIC] test 1\n", buf.String())
}

func Test_LoggerPanicIfError(t *testing.T) {
	logger, buf := setupLoggerTest()
	require.Panics(t, func() { logger.PanicIfError(assert.AnError) })
	require.Equal(t, "[PANIC] "+assert.AnError.Error()+"\n", buf.String())
	require.NotPanics(t, func() { logger.PanicIfError(nil) })
}

func Test_StringTaggedLogger(t *testing.T) {
	logger := NewStringTaggedLogger()
	logger.Debugf("test %d", 1)
	require.Equal(t, "[DEBUG] test 1\n", logger.sb.String())
	logger.Errorf("test %d", 2)
	require.Equal(t, "[DEBUG] test 1\n[ERROR] test 2\n", logger.sb.String())
	logger.Infof("test %d", 3)
	require.Equal(t, "[DEBUG] test 1\n[ERROR] test 2\n[INFO] test 3\n",
		logger.sb.String())
	require.Panics(t, func() { logger.Panicf("test %d", 4) })
	require.Equal(t,
		"[DEBUG] test 1\n[ERROR] test 2\n[INFO] test 3\n[PANIC] test 4\n",
		logger.sb.String())
}

func Test_StringTaggedLogger_panics_if_error(t *testing.T) {
	logger := NewStringTaggedLogger()
	require.Panics(t, func() { logger.PanicIfError(assert.AnError) })
	require.Equal(t, "[PANIC] "+assert.AnError.Error()+"\n", logger.sb.String())
}

func Test_StringTaggedLogger_does_not_panic(t *testing.T) {
	logger := NewStringTaggedLogger()
	require.NotPanics(t, func() { logger.PanicIfError(nil) })
	require.Empty(t, logger.sb.String())
}
