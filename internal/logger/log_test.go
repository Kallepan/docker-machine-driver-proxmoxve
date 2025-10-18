package logger_test

import (
	"testing"

	"github.com/deepshore/docker-machine-driver-proxmoxve/internal/logger"
)

func TestInitCreatesLogger(t *testing.T) {
	// Ensure fresh start
	logger.Logger = nil

	logger.Init()
	if logger.Logger == nil {
		t.Fatal("expected logger to be initialized, got nil")
	}
}

func TestInitIsIdempotent(t *testing.T) {
	logger.Logger = nil
	logger.Init()
	first := logger.Logger

	logger.Init() // should not change the instance
	second := logger.Logger

	if first != second {
		t.Fatal("expected Init() to be idempotent, but got different instances")
	}
}
