package slogtee

import (
	"context"
	"log/slog"
)

// Compile-time check *Handler implements slog.Handler.
var _ slog.Handler = (*Handler)(nil)

// Handler is a [slog.Handler] that forwards operations to multiple other [slog.Handler]s.
type Handler struct {
	handlers []slog.Handler
}

// NewHandler creates a new [Handler].
func NewHandler(handlers ...slog.Handler) *Handler {
	return &Handler{handlers: handlers}
}

func (t *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range t.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}

	return false
}

func (t *Handler) Handle(ctx context.Context, record slog.Record) error {
	for _, handler := range t.handlers {
		err := handler.Handle(ctx, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	handlers := make([]slog.Handler, len(t.handlers))
	for i, handler := range t.handlers {
		handlers[i] = handler.WithAttrs(attrs)
	}

	return &Handler{handlers: handlers}
}

func (t *Handler) WithGroup(name string) slog.Handler {
	handlers := make([]slog.Handler, len(t.handlers))
	for i, handler := range t.handlers {
		handlers[i] = handler.WithGroup(name)
	}

	return &Handler{handlers: handlers}
}
