package boot

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(env Env) *zap.Logger {
	if env.AppEnv == ProdEnv {

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   env.LogSaveFile,
			MaxSize:    env.MaxLogFileSize,
			MaxBackups: env.MaxLogFileBackups,
			MaxAge:     env.MaxLogFileAge,
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
