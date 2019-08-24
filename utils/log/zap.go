package utils

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"sync"
	"utils/config"
)

var (
	l                              *Logger
	sp                             = string(filepath.Separator)
	errWS, warnWS, infoWS, debugWS zapcore.WriteSyncer       // IO输出
	debugConsoleWS                 = zapcore.Lock(os.Stdout) // 控制台标准输出
	errorConsoleWS                 = zapcore.Lock(os.Stderr)
)

// Options 配置项
type Options struct {
	zap.Config
	LogFileDir    string `json:"log_file_dir"`
	AppName       string `json:"appName"`
	ErrorFileName string `json:"errorFileName"`
	WarnFileName  string `json:"warnFileName"`
	InfoFileName  string `json:"infoFileName"`
	DebugFileName string `json:"debugFileName"`
	MaxSize       int    `json:"maxSize"` // megabytes
	MaxBackups    int    `json:"maxBackups"`
	MaxAge        int    `json:"maxAge"` // days
}

type Logger struct {
	*zap.Logger
	once      sync.Once
	Opts      *Options
	zapConfig zap.Config
}

func init() {
	l = &Logger{
		Opts: &Options{},
	}
}

func InitLogger() {
	l.once.Do(func() {
		l.loadCfg()
		l.init()
		l.Info("[initLogger] zap logger initializing completed")
	})
}

func (l *Logger) init() {
	l.setSyncers()
	var err error
	l.Logger, err = l.zapConfig.Build(l.cores())
	if err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}

func (l *Logger) MakeErrFields(err error, fields ...zap.Field) []zap.Field {
	return append([]zap.Field{zap.Error(err)}, fields...)
}

func (l *Logger) loadCfg() {
	err := config.C().Path(l.Opts, "zap")
	if err != nil {
		panic(err)
	}
	data, _ := json.Marshal(l.Opts)
	fmt.Println(string(data))
	if l.Opts.Development {
		l.zapConfig = zap.NewDevelopmentConfig()
	} else {
		l.zapConfig = zap.NewProductionConfig()
		l.zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	// application log output path
	if l.Opts.OutputPaths == nil || len(l.Opts.OutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stdout"}
	}

	//  error of zap-self log
	if l.Opts.ErrorOutputPaths == nil || len(l.Opts.ErrorOutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stderr"}
	}

	// 默认输出到程序运行目录的logs子目录
	if l.Opts.LogFileDir == "" {
		l.Opts.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		l.Opts.LogFileDir += sp + "logs" + sp
	}

	if l.Opts.AppName == "" {
		l.Opts.AppName = "app"
	}

	if l.Opts.ErrorFileName == "" {
		l.Opts.ErrorFileName = "error.log"
	}

	if l.Opts.WarnFileName == "" {
		l.Opts.WarnFileName = "warn.log"
	}

	if l.Opts.InfoFileName == "" {
		l.Opts.InfoFileName = "info.log"
	}

	if l.Opts.DebugFileName == "" {
		l.Opts.DebugFileName = "debug.log"
	}

	if l.Opts.MaxSize == 0 {
		l.Opts.MaxSize = 50
	}

	if l.Opts.MaxBackups == 0 {
		l.Opts.MaxBackups = 3
	}

	if l.Opts.MaxAge == 0 {
		l.Opts.MaxAge = 30
	}
}

func (l *Logger) setSyncers() {
	f := func(fN string) zapcore.WriteSyncer {
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.Opts.LogFileDir + sp + l.Opts.AppName + "-" + fN,
			MaxSize:    l.Opts.MaxSize,
			MaxBackups: l.Opts.MaxBackups,
			MaxAge:     l.Opts.MaxAge,
			Compress:   true,
			LocalTime:  true,
		})
	}

	errWS = f(l.Opts.ErrorFileName)
	warnWS = f(l.Opts.WarnFileName)
	infoWS = f(l.Opts.InfoFileName)
	debugWS = f(l.Opts.DebugFileName)

	return
}

func (l *Logger) cores() zap.Option {

	fileEncoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(l.zapConfig.EncoderConfig)

	errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zapcore.WarnLevel && zapcore.WarnLevel-l.zapConfig.Level.Level() > -1
	})
	warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel && zapcore.WarnLevel-l.zapConfig.Level.Level() > -1
	})
	infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel && zapcore.InfoLevel-l.zapConfig.Level.Level() > -1
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && zapcore.DebugLevel-l.zapConfig.Level.Level() > -1
	})
	var cores []zapcore.Core
	if l.Opts.Development {
		cores = []zapcore.Core{
			// region 控制台

			// error
			zapcore.NewCore(consoleEncoder, errorConsoleWS, errPriority),

			// warning
			zapcore.NewCore(consoleEncoder, debugConsoleWS, warnPriority),

			// info
			zapcore.NewCore(consoleEncoder, debugConsoleWS, infoPriority),

			// debug
			zapcore.NewCore(consoleEncoder, debugConsoleWS, debugPriority),

			// endregion
		}
	} else {
		cores = []zapcore.Core{
			// region 日志文件

			// error 及以上
			zapcore.NewCore(fileEncoder, errWS, errPriority),

			// warn
			zapcore.NewCore(fileEncoder, warnWS, warnPriority),

			// info
			zapcore.NewCore(fileEncoder, infoWS, infoPriority),

			// debug
			zapcore.NewCore(fileEncoder, debugWS, debugPriority),

			// endregion
		}
	}

	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}

func GetLoger() *Logger {
	return l
}
