package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	logger = slog.New(NewFormatHandler(os.Stdout, slog.LevelInfo))
}

type FormatHandler struct {
	level  slog.Level
	writer io.Writer
}

func NewFormatHandler(w io.Writer, level slog.Level) *FormatHandler {
	return &FormatHandler{
		level:  level,
		writer: w,
	}
}

func (h *FormatHandler) Handle(ctx context.Context, record slog.Record) error {
	ts := record.Time.Format("2006-01-02 15:04:05")
	lvl := colorForLevel(record.Level)
	msg := record.Message

	var attrsStr string
	record.Attrs(func(a slog.Attr) bool {
		attrsStr += fmt.Sprintf(" %s=%v", a.Key, a.Value)
		return true
	})

	output := fmt.Sprintf("[%s] %s %s%s\n", ts, lvl, msg, attrsStr)
	_, err := h.writer.Write([]byte(output))
	return err
}

func (h *FormatHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *FormatHandler) WithGroup(name string) slog.Handler {
	return h
}

func (h *FormatHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func colorForLevel(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return "\033[34mDEBUG\033[0m"
	case slog.LevelInfo:
		return "\033[32mINFO\033[0m"
	case slog.LevelWarn:
		return "\033[33mWARN\033[0m"
	case slog.LevelError:
		return "\033[31mERROR\033[0m"
	default:
		return level.String()
	}
}

func Info(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}
