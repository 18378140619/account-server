package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if viper.GetBool("mode.develop") {
		logMode = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), getWriteSyncer(), logMode)
	//core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)
	return zap.New(core).Sugar()
}

// 配置日志格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

// log地址
func getWriteSyncer() zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	strRootDir, _ := os.Getwd()
	stLogFilePath := strRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".txt"

	luberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"),
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"),
		Compress:   false, // disabled by default
	}

	return zapcore.AddSync(luberjackSyncer)
}
