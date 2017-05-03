package handler

import (
	"../api"
	"log"
)

type Logger struct{}

func (l *Logger) Call(msg api.HubMessage) {
	log.Print(msg)
}
