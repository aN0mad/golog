package log

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

// LLogger is a charmbracelet logger type redefinition
type LLogger = log.Logger

// Logger is this package level logger
var Logger *LLogger

type ctxLogger struct{}

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l *log.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) *log.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*log.Logger); ok {
		return l
	}
	return Logger
}

func init() {
	styles := log.DefaultStyles()
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["err"] = lipgloss.NewStyle().Bold(true)

	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
	})
	Logger.SetStyles(styles)
	Logger.SetLevel(log.InfoLevel)
}

// EnableDebug enabled debug logging and caller reporting
func EnableDebug() {
	Logger.SetLevel(log.DebugLevel)
	Logger.SetReportCaller(true)
}

// EnableSilence will silence most logs, except this written with Print
func EnableSilence() {
	Logger.SetLevel(log.FatalLevel + 100)
}

// Debug logs debug messages
func Debug(msg string, keyvals ...interface{}) {
	Logger.Helper()
	Logger.Debug(msg, keyvals...)
}

// Debugf prints a debug message with formatting.
func Debugf(format string, args ...interface{}) {
	Logger.Helper()
	Logger.Debug(fmt.Sprintf(format, args...))
}

// Info logs info messages
func Info(msg string, keyvals ...interface{}) {
	Logger.Helper()
	Logger.Info(msg, keyvals...)
}

// Infof prints an info message with formatting.
func Infof(format string, args ...interface{}) {
	Logger.Helper()
	Logger.Info(fmt.Sprintf(format, args...))
}

// Warn logs warning messages
func Warn(msg string, keyvals ...interface{}) {
	Logger.Helper()
	Logger.Warn(msg, keyvals...)
}

// Warnf prints a warning message with formatting.
func Warnf(format string, args ...interface{}) {
	Logger.Helper()
	Logger.Warn(fmt.Sprintf(format, args...))
}

// Error logs error messages
func Error(msg string, keyvals ...interface{}) {
	Logger.Helper()
	Logger.Error(msg, keyvals...)
}

// Errorf prints an error message with formatting.
func Errorf(format string, args ...interface{}) {
	Logger.Helper()
	Logger.Error(fmt.Sprintf(format, args...))
}

// Fatal logs fatal messages and panics
func Fatal(msg string, keyvals ...interface{}) {
	Logger.Helper()
	Logger.Fatal(msg, keyvals...)
}

// Fatalf prints a fatal message with formatting and exits.
func Fatalf(format string, args ...interface{}) {
	Logger.Helper()
	Logger.Fatal(fmt.Sprintf(format, args...))
	os.Exit(1)
}

// Print logs messages regardless of level
func Print(msg string, keyvals ...interface{}) {
	Logger.Helper()
	Logger.Print(msg, keyvals...)
}

// Printf prints a message with no level and formatting.
func Printf(format string, args ...interface{}) {
	Logger.Helper()
	Logger.Print(fmt.Sprintf(format, args...))
}

// With returns a sublogger with a prefix
func With(keyvals ...interface{}) *LLogger {
	return Logger.With(keyvals...)
}
