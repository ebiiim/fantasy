package log

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type Logger struct {
	logger zerolog.Logger
}

type Type string

const (
	logType  = "type"
	logWhere = "where"

	TypeValidation  = "validation"
	TypeIntelligent = "intelligent"
	TypeInternal    = "internal"
)

func NewLogger(component string) *Logger {
	return &Logger{logger: zlog.With().Str("component", component).Logger()}
}

func (l *Logger) Error(lt Type, who string, msg string) {
	l.logger.Error().Str(logType, string(lt)).Str(logWhere, who).Msg(msg)
}

func (l *Logger) Info(lt Type, who string, msg string) {
	l.logger.Info().Str(logType, string(lt)).Str(logWhere, who).Msg(msg)
}

func (l *Logger) Debug(lt Type, who string, msg string) {
	l.logger.Debug().Str(logType, string(lt)).Str(logWhere, who).Msg(msg)
}
