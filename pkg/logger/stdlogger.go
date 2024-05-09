package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

func NewConsoleWriter(level zerolog.Level) zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
}

type DummyHook struct {
}

func (h DummyHook) Run(e *zerolog.Event, level zerolog.Level, message string) {

}
