package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	LogCategoryDefault = "default"
	LogCategorySystem  = "application.system"
	LogCategoryLogic   = "application.logic"
	LogCategoryDebug   = "application.debug"
)

func NewLogger() LoggerInterface {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	return &Logger{}
}

type LoggerInterface interface {
	Info(category string, message string)

	Warning(category string, message string)

	Error(category string, message string)

	Debug(category string, message string)
}

type Logger struct{}

func (l *Logger) Info(category string, message string) {
	log.Info().Str("category", category).Msg("[Metadata Server] " + message)
}

func (l *Logger) Warning(category string, message string) {
	log.Warn().Str("category", category).Msg("[Metadata Server] " + message)
}

func (l *Logger) Error(category string, message string) {
	log.Error().Str("category", category).Msg("[Metadata Server] " + message)
}

func (l *Logger) Debug(category string, message string) {
	log.Debug().Str("category", category).Msg("[Metadata Server] " + message)
}
