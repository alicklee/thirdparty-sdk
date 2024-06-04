package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/CloudcadeSF/thirdparty-sdk/utils"

	"github.com/sirupsen/logrus"
)

var (
	logging = New()
)

type Logger struct {
	logger *logrus.Logger
}

type LoggerConfig struct {
	Out             string
	Level           string
	TimestampFormat string
	FullTimestamp   bool
}

func New() *Logger {
	return &Logger{
		logger: logrus.New(),
	}
}

func (l *Logger) Init(conf *LoggerConfig) {
	if conf.Level == "" {
		l.SetLevel(utils.GlobalObject.LogLevel)
	}

	if conf.Out == "" {
		l.logger.SetOutput(getOutput(utils.GlobalObject.LogOutput))
	}

	if conf.TimestampFormat == "" {
		l.logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
	}
}

func (l *Logger) SetOutput(out io.Writer) {
	l.logger.Out = out
	l.logger.Level = logrus.TraceLevel
}

func (l *Logger) SetLevel(logLevel string) {
	switch logLevel {
	case "panic":
		l.logger.SetLevel(logrus.PanicLevel)
	case "error":
		l.logger.SetLevel(logrus.ErrorLevel)
	case "warning":
		l.logger.SetLevel(logrus.WarnLevel)
	case "info":
		l.logger.SetLevel(logrus.InfoLevel)
	case "debug":
		l.logger.SetLevel(logrus.DebugLevel)
		return
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.logger.Warningf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.logger.Warning(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *Logger) Warningln(args ...interface{}) {
	l.logger.Warningln(args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.logger.Panicln(args...)
}

func InitLogger(conf *LoggerConfig) {
	logging.Init(conf)
}

func getOutput(outputType string) io.Writer {
	var writer io.Writer
	if outputType == "std" {
		writer = os.Stdout
		return writer
	} else {
		flag := os.O_CREATE | os.O_RDWR | os.O_APPEND
		today := time.Now().Format("20060102")
		fileName := path.Join(utils.GlobalObject.LogPath, strings.ReplaceAll(utils.GlobalObject.Name, " ", "")+"-"+today+".log")
		logFile, err := os.OpenFile(fileName, flag, 0666)
		if err != nil {
			panic(fmt.Errorf("log file `%s` cannot access.", outputType))
		}
		return logFile
	}
}

func Debugf(format string, args ...interface{})   { logging.Debugf(format, args...) }
func Infof(format string, args ...interface{})    { logging.Infof(format, args...) }
func Warningf(format string, args ...interface{}) { logging.Warningf(format, args...) }
func Errorf(format string, args ...interface{})   { logging.Errorf(format, args...) }
func Panicf(format string, args ...interface{})   { logging.Panicf(format, args...) }
func Debug(args ...interface{})                   { logging.Debug(args...) }
func Info(args ...interface{})                    { logging.Info(args...) }
func Warning(args ...interface{})                 { logging.Warning(args...) }
func Error(args ...interface{})                   { logging.Error(args...) }
func Panic(args ...interface{})                   { logging.Panic(args...) }
func Debugln(args ...interface{})                 { logging.Debugln(args...) }
func Infoln(args ...interface{})                  { logging.Infoln(args...) }
func Warningln(args ...interface{})               { logging.Warningln(args...) }
func Errorln(args ...interface{})                 { logging.Errorln(args...) }
func Panicln(args ...interface{})                 { logging.Panicln(args...) }
