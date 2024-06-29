package zerolog

import "github.com/rs/zerolog/log"

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l Logger) Errorf(format string, args ...interface{}) {
	log.Error().Msgf(format, args...)
}

func (l Logger) Warnf(format string, args ...interface{}) {
	log.Warn().Msgf(format, args...)
}

func (l Logger) Infof(format string, args ...interface{}) {
	log.Info().Msgf(format, args...)
}

func (l Logger) Debugf(format string, args ...interface{}) {
	log.Debug().Msgf(format, args...)
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	log.Fatal().Msgf(format, args...)
}
