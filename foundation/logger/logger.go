package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func New(service string) *zerolog.Logger {
	// Setting up the logger with the desired configuration
	logger := zerolog.New(os.Stdout).
		With(). // Adding contextual fields
		Str("service", service).
		Logger()

	return &logger
}
