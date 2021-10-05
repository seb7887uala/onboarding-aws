package logger

import (
	"go.uber.org/zap"
)

var Instance *zap.Logger

func Init() {
	Instance, _ = zap.NewProduction()
}
