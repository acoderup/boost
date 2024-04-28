package safe

import (
	"context"
	"log"
	"time"
)

type Handler struct {
	Func func(err error)
}

var defaultHandler = &Handler{
	Func: func(err error) {
		log.Println(err)
	},
}

func Default() *Handler {
	return defaultHandler
}

func (*Handler) Do(a any) {
	if err := Do(a); err != nil {
		defaultHandler.Func(err)
	}
}

func (*Handler) DoWithContext(ctx context.Context, a any) {
	if err := DoWithContext(ctx, a); err != nil {
		defaultHandler.Func(err)
	}
}

func (*Handler) DoWithTimeout(d time.Duration, a any) {
	if err := DoWithTimeout(d, a); err != nil {
		defaultHandler.Func(err)
	}
}

func Wrap(a any) func() {
	return func() {
		defaultHandler.Do(a)
	}
}

func (h *Handler) Wrap(a any) func() {
	return func() {
		h.Do(a)
	}
}

func WrapWithContext(ctx context.Context, a any) func() {
	return func() {
		defaultHandler.DoWithContext(ctx, a)
	}
}

func (h *Handler) WrapWithContext(ctx context.Context, a any) func() {
	return func() {
		h.DoWithContext(ctx, a)
	}
}

func WrapWithTimeout(d time.Duration, a any) func() {
	return func() {
		defaultHandler.DoWithTimeout(d, a)
	}
}

func (h *Handler) WrapWithTimeout(d time.Duration, a any) func() {
	return func() {
		h.DoWithTimeout(d, a)
	}
}
