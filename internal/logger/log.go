package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
)

var Logger *zap.Logger
var once sync.Once

func Init() {
	once.Do(func() {
		// Default to development logger if not in production
		if os.Getenv("APP_ENV") != "production" {
			Logger = zap.Must(zap.NewDevelopment())
			return
		}
		Logger = zap.Must(zap.NewProduction())
	})
}

/**
 * CallBackOnExit ensures that any buffered log entries are flushed
 * before the application exits. This is especially important for
 * asynchronous logging backends to prevent log loss.
 */
func CallBackOnExit() {
	if Logger == nil {
		return
	}

	_ = Logger.Sync()
}
