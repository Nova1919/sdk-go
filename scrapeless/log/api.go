package log

import "fmt"

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	msg := fmt.Sprint(args...)
	logger.Trace().CallerSkipFrame(1).Msg(msg)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	logger.Trace().CallerSkipFrame(1).Msgf(format, args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	msg := fmt.Sprint(args...)
	logger.Debug().CallerSkipFrame(1).Msg(msg)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	logger.Debug().CallerSkipFrame(1).Msgf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	msg := fmt.Sprint(args...)
	logger.Info().CallerSkipFrame(1).Msg(msg)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	logger.Info().CallerSkipFrame(1).Msgf(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	msg := fmt.Sprint(args...)
	logger.Warn().CallerSkipFrame(1).Msg(msg)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	logger.Warn().CallerSkipFrame(1).Msgf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	msg := fmt.Sprint(args...)
	logger.Error().CallerSkipFrame(1).Msg(msg)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	logger.Error().CallerSkipFrame(1).Msgf(format, args...)
}
