package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	instance *zap.Logger
	once     sync.Once
)

func Setup() *zap.Logger {
	once.Do(func() {
		instance, _ = zap.NewProduction()
	})
	return instance
}
