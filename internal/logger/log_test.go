package logger

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestGetZapLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected zapcore.Level
	}{
		{"debug level", "debug", zapcore.DebugLevel},
		{"info level", "info", zapcore.InfoLevel},
		{"warn level", "warn", zapcore.WarnLevel},
		{"error level", "error", zapcore.ErrorLevel},
		{"dpanic level", "dpanic", zapcore.DPanicLevel},
		{"panic level", "panic", zapcore.PanicLevel},
		{"fatal level", "fatal", zapcore.FatalLevel},
		{"invalid level defaults to info", "invalid", zapcore.InfoLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getZapLogLevel(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
