package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init() {
	log_format := os.Getenv("LOG_FORMAT")
	if log_format == "" {
		log_format = "console"
	}

	log_level := os.Getenv("LOG_LEVEL")
	if log_level == "" {
		log_level = "info"
	}

	var zapConfig zap.Config
	if log_format == "json" {
		zapConfig = zap.NewProductionConfig()
		zapConfig.EncoderConfig.TimeKey = "timestamp"
		zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapConfig.Level = zap.NewAtomicLevelAt(getZapLogLevel(log_level))

	createdLogger, err := zapConfig.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		// Hmm...
		logger = zap.NewExample()
		logger.Error("Failed to initialize logger, using example logger", zap.Error(err))
	} else {
		logger = createdLogger
	}
}

// Info logs an informational message with optional key-value pairs.
func Info(msg string, keysAndValues ...any) {
	logger.Sugar().Infow(msg, keysAndValues...)
}

// Error logs an error message with optional key-value pairs.
func Error(msg string, keysAndValues ...any) {
	logger.Sugar().Errorw(msg, keysAndValues...)
}

// Debug logs a debug message with optional key-value pairs.
func Debug(msg string, keysAndValues ...any) {
	logger.Sugar().Debugw(msg, keysAndValues...)
}

// Warn logs a warning message with optional key-value pairs.
func Warn(msg string, keysAndValues ...any) {
	logger.Sugar().Warnw(msg, keysAndValues...)
}

// Fatal logs a fatal message with optional key-value pairs and exits the application.
func Fatal(msg string, keysAndValues ...any) {
	logger.Sugar().Fatalw(msg, keysAndValues...)
	os.Exit(1)
}

func GetLogger() *zap.Logger {
	return logger
}

// Sync flushes any buffered log entries.
func Sync() {
	_ = logger.Sync()
}

func getZapLogLevel(level string) zapcore.Level {
	level = strings.ToLower(level)

	// Implement the logic to convert string level to zap.AtomicLevel
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
