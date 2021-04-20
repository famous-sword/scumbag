package logger

import (
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/foundation"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var writer *zap.Logger = zap.NewNop()

type Bootstrapper struct{}

func (_ *Bootstrapper) Bootstrap() (err error) {
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.String("logging.file"),
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel)
	writer = zap.New(core, zap.AddCaller())

	return nil
}

func NewBootstrapper() foundation.Bootable {
	return &Bootstrapper{}
}

func Writter() *zap.Logger {
	return writer
}
