package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var GLog *zap.Logger

func InitLogger(level, logType, filename string, maxSize, maxBackup, maxAge int, compress bool, isLocal bool) {
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		panic("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}

	writeSyncer := getLogWriter(logType, filename, maxSize, maxBackup, maxAge, compress, isLocal)

	GLog = zap.New(
		zapcore.NewCore(getEncoder(isLocal), writeSyncer, logLevel),
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)

	// 将自定义的 logger 替换为全局的 logger，这样 zap.L().Fatal() 调用时，就会使用我们自定的 Logger
	zap.ReplaceGlobals(GLog)
}

// getEncoder 设置日志存储格式
func getEncoder(isLocal bool) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	if isLocal {
		// 终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		// 本地设置内置的 Console 解码器（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 线上环境使用 JSON 编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// getLogWriter 日志记录介质
func getLogWriter(logType, filename string, maxSize, maxBackup, maxAge int, compress bool, isLocal bool) zapcore.WriteSyncer {
	// todo 有 bug 如果配置了按照日期记录日志文件。有问题，不重启无法按日期
	if logType == "daily" {
		logName := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))
		filename = strings.ReplaceAll(filename, "logs.log", logName)
	}

	// 滚动日志，详见 config/log.go
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	// 配置输出介质
	if isLocal {
		// 本地开发终端打印和记录文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}

	// 生产环境只记录文件
	return zapcore.AddSync(lumberJackLogger)
}
