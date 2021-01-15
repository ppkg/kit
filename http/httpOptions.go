package http

import "time"

type httpOptions interface {
	apply(*httpHandle)
}

type funcHttpOption struct {
	f func(*httpHandle)
}

func (fdo *funcHttpOption) apply(do *httpHandle) {
	fdo.f(do)
}

func newFuncHttpOption(f func(*httpHandle)) *funcHttpOption {
	return &funcHttpOption{
		f: f,
	}
}

func defaultHttpOption() *funcHttpOption {
	return newFuncHttpOption(func(h *httpHandle) {
		h.timeout = 3 * time.Second
	})
}

func WithTimeout(duration time.Duration) *funcHttpOption {
	return newFuncHttpOption(func(h *httpHandle) {
		h.timeout = duration
	})
}
