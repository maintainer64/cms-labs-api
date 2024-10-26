package logs

import (
	"os"

	"github.com/rs/zerolog"
)

func NewZeroLogger(name string) zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("logger", name)
	})
	return logger
}

var (
	ZeroLog zerolog.Logger = NewZeroLogger("main")
)

func ZeroLogInit(debug bool) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
