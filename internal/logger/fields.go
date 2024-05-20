package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func StringField(key string, val string) zapcore.Field {
	return zap.String(key, val)
}

func IntField(key string, val int) zapcore.Field {
	return zap.Int(key, val)
}

func DurationField(key string, val time.Duration) zapcore.Field {
	return zap.Duration(key, val)
}
