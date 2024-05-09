package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
)

type Logger struct {
	*zerolog.Logger
}

func NewConsoleAndHookLogger(logLevel string, hookWriters ...io.Writer) (*Logger, error) {
	if logLevel == "" {
		logLevel = "trace"
	}
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	consoleWriter := NewConsoleWriter(level)
	consoleWriter.FormatTimestamp = func(i interface{}) string {
		return ""
	}
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("\n    %s: ", i)
	}
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	l := zerolog.Logger{}
	var hookIsNil bool
	var multi zerolog.LevelWriter
	if len(hookWriters) == 0 {
		multi = zerolog.MultiLevelWriter(consoleWriter, l.Hook(&DummyHook{}))
		hookIsNil = true
	} else {
		writers := make([]io.Writer, 0, len(hookWriters)+1)
		writers = append(writers, consoleWriter)
		for _, writer := range hookWriters {
			if writer != nil {
				writers = append(writers, writer)
			}
		}
		multi = zerolog.MultiLevelWriter(writers...)
	}
	l3 := zerolog.New(multi).Level(level).With().Logger()
	if hookIsNil {
		l3.Warn().Msg("empty logger hook")
	}
	return &Logger{&l3}, nil
}
