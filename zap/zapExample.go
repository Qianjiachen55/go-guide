package main

import (
	"fmt"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//logger := zap.NewExample()
	//
	//defer logger.Sync()
	//
	//sugar := logger.Sugar()
	//
	//sugar.Infof("hhhhhhh")

	//var (
	//rawJSON := []byte(`{
	//  "level": "debug",
	//  "encoding": "json",
	//  "outputPaths": ["stdout", "/tmp/logs"],
	//  "errorOutputPaths": ["stderr"],
	//  "initialFields": {"foo": "bar"},
	//  "encoderConfig": {
	//    "messageKey": "message",
	//    "levelKey": "level",
	//    "levelEncoder": "lowercase"
	//  }
	//}`)
	//
	//var cfg zap.Config
	//
	//if err:=json.Unmarshal(rawJSON,&cfg);err != nil{
	//	panic(err)
	//}
	//)
	var cfg = zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "",
			StacktraceKey:  "",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths:      []string{"stdout", "log/log.txt"},
		ErrorOutputPaths: []string{"stderr", "log/error.txt"},
		InitialFields:    nil,
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("logger init success")
	var (
		b int
		a int
	)
	a = 5
	b = 0
	defer errFun(logger)

	i := a / b

	fmt.Println(i)
	fmt.Println("qwefqw")

}

func errFun(logger *zap.Logger) {
	if err := recover();err != nil{
		//logger.Error(cast.ToString(err))

		//logger.Panic(cast.ToString(err))
		logger.DPanic(cast.ToString(err))
		fmt.Println("end!!!")
	}
}
