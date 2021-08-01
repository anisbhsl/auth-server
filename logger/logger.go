package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var logger *logging

type Config struct {
	Service  string
	Level    string
	OutPath  string
	Encoding string
}

type logging struct {
	Log *zap.Logger
}

var level = map[string]zapcore.Level{
	"info":  zap.InfoLevel,
	"debug": zap.DebugLevel,
	"error": zap.ErrorLevel,
	"fatal": zap.FatalLevel,
	"panic": zap.PanicLevel,
	"warn":  zap.WarnLevel,
}

func NewLogger(config Config) {
	logLevel, ok := level[strings.ToLower(config.Level)]
	if !ok {
		logLevel = level["debug"]
	}

	if config.OutPath == "" {
		config.OutPath = "stdout"
	}

	if config.Encoding == "" {
		config.Encoding = "json"
	}

	logConfig := zap.Config{
		OutputPaths: []string{config.OutPath},
		Encoding:    config.Encoding, //options: console/json
		Level:       zap.NewAtomicLevelAt(logLevel),
		InitialFields: map[string]interface{}{
			"service": config.Service,
		},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	Log, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	logger = &logging{Log: Log}
	logger.Log.Info(fmt.Sprintf("Log level set to: %s | File output set to: %s | Log Encoding set to: %s", logLevel, config.OutPath, config.Encoding))
}

func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func Debug(msg string, tags ...zap.Field) {
	logger.Log.Debug(msg, tags...)
	logger.Log.Sync()
}

func Info(msg string, tags ...zap.Field) {
	logger.Log.Info(msg, tags...)
	logger.Log.Sync()
}

func Warn(msg string, tags ...zap.Field) {
	logger.Log.Warn(msg, tags...)
	logger.Log.Sync()
}

func Error(msg string, tags ...zap.Field) {
	logger.Log.Error(msg, tags...)
	logger.Log.Sync()
}

func Fatal(msg string, tags ...zap.Field) {
	logger.Log.Fatal(msg, tags...)
	logger.Log.Sync()
}

func Panic(msg string, tags ...zap.Field) {
	logger.Log.Panic(msg, tags...)
	logger.Log.Sync()
}

func TraceRequestWithContext(ctx context.Context) zap.Field{
	traceID:=ctx.Value("traceID")
	return zap.String("traceID",fmt.Sprintf("%v",traceID))
}
