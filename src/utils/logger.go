package client_server

import (
    "time"
    "go.uber.org/zap"
    "github.com/fatih/color"
    "go.uber.org/zap/zapcore"
)

func customLevelEncoder(isServer bool) zapcore.LevelEncoder {
    var useColors bool = true
    if LoadConfig() == nil {
        useColors = COLORED_LOGS
    }

    return func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
        var levelStr string
        switch level {
            case zapcore.InfoLevel:
                levelStr = "[INFO]"
                if isServer && useColors {
                    levelStr = color.HiGreenString(levelStr)
                } else if useColors {
                    levelStr = color.HiBlueString(levelStr)
                }
            case zapcore.ErrorLevel:
                levelStr = "[ERROR]"
                if useColors {
                    levelStr = color.HiRedString(levelStr)
                }
            default:
                levelStr = level.CapitalString()
        }
        enc.AppendString(levelStr)
    }
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format("15:04:05.000"))
}

func InitLogger(isServer bool) *zap.Logger {
    config := zap.NewDevelopmentConfig()
    config.DisableCaller = true
    config.DisableStacktrace = true
    config.EncoderConfig.EncodeTime = customTimeEncoder
    config.EncoderConfig.EncodeLevel = customLevelEncoder(isServer)

    logger, err := config.Build()
    if err != nil {
        panic(err)
    }
    return logger
}
