package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 1.
	// sugar := zap.NewExample().Sugar()
	// sugar.Info("Hell %s, age:%d ", "TipsGo", 40)

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipsGo"), zap.Int("age", 40))


	// 2.
	// logger := zap.NewExample()
	// logger.Info("hello")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// 3.

	encoder := getEncoderLog()
	sync := getWrite()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// milisecond -> time YYYY-MM-DD...
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> Time
	encodeConfig.TimeKey = "time"

	//from info Info
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//"caller":"cli/main.log.go"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)

}

// ghi vao vi tri nao
func getWrite() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE | os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}