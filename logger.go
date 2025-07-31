package slogtee

import "log/slog"

// New creates a [slog.Logger] with multiple [slog.Handler]s.
func New(handlers ...slog.Handler) *slog.Logger {
	return slog.New(NewHandler(handlers...))
}
