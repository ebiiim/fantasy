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
	logObj   = "obj"

	TypeValidation  = "validation"
	TypeIntelligent = "intelligent"
	TypeInternal    = "internal"
)

func NewLogger(component string) *Logger {
	return &Logger{logger: zlog.With().Str("component", component).Logger()}
}

func (l *Logger) Error(lt Type, where, obj, msg string) {
	l.logger.Error().Str(logType, string(lt)).Str(logWhere, where).Str(logObj, obj).Msg(msg)
}

func (l *Logger) Info(lt Type, where, obj, msg string) {
	l.logger.Info().Str(logType, string(lt)).Str(logWhere, where).Str(logObj, obj).Msg(msg)
}

func (l *Logger) Debug(lt Type, where, obj, msg string) {
	l.logger.Debug().Str(logType, string(lt)).Str(logWhere, where).Str(logObj, obj).Msg(msg)
}
