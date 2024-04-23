package logger

import (
	. "forest-run/common/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(envType EnvType, config LoggerConfig) *zap.Logger {
	if envType == ProdEnv {

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.LogSaveFile,
			MaxSize:    config.MaxLogFileSize,
			MaxBackups: config.MaxLogFileBackups,
			MaxAge:     config.MaxLogFileAge,
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zap.InfoLevel,
		)

		logger := zap.New(core)

		return logger
	}

	logger, _ := zap.NewDevelopment()
	return logger
}
