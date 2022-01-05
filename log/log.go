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
	TypeSystem      = "system"
	TypeInit        = "init"
)

func NewLogger(component string) *Logger {
	return &Logger{logger: zlog.With().Str("component", component).Logger()}
}

func (l *Logger) Fatal(lt Type, where, obj, msgf string, v ...interface{}) {
	l.logger.Fatal().Str(logType, string(lt)).Str(logWhere, where).Str(logObj, obj).Msgf(msgf, v...)
}

func (l *Logger) Error(lt Type, where, obj, msgf string, v ...interface{}) {
	l.logger.Error().Str(logType, string(lt)).Str(logWhere, where).Str(logObj, obj).Msgf(msgf, v...)
}

func (l *Logger) Warn(lt Type, where, obj, msgf string, v ...interface{}) {
	l.logger.Warn().Str(logType, string(lt)).Str(logWhere, where).Str(logObj, obj).Msgf(msgf, v...)
}

func (l *Logger) Info(lt Type, where, obj, msgf string, v ...interface{}) {
	l.logger.Info().Str(logType, string(lt)).Str(logWhere, where).Str(logObj, obj).Msgf(msgf, v...)
}

func (l *Logger) Debug(lt Type, where, obj, msgf string, v ...interface{}) {
	l.logger.Debug().Str(logType, string(lt)).Str(logWhere, where).Str(logObj, obj).Msgf(msgf, v...)
}

func (l *Logger) Trace(lt Type, where, obj, msgf string, v ...interface{}) {
	l.logger.Trace().Str(logType, string(lt)).Str(logWhere, where).Str(logObj, obj).Msgf(msgf, v...)
}
