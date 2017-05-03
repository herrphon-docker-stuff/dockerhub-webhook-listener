package server

import (
	"./api"
)

type Handler interface {
	Call(api.HubMessage)
}

type HandlerRegistry struct {
	entries []func(api.HubMessage)
}

func (r *HandlerRegistry) Add(h func(msg api.HubMessage)) {
	r.entries = append(r.entries, h)
	return
}

func (r *HandlerRegistry) Call(msg api.HubMessage) {
	for _, h := range r.entries {
		go h(msg)
	}
}
