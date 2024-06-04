package log

/*
import (
	"io"
)

var (
	_std = New()
)

type LoggerConfig struct {
	Out             io.Writer
	Level           string
	TimestampFormat string
	FullTimestamp   bool
}

func InitLogger(conf *LoggerConfig) {
	_std.SetLevel(conf.Level)
	_std.logger.SetOutput(conf.Out)
	_std.logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: conf.TimestampFormat,
		FullTimestamp:   conf.FullTimestamp,
	})
}

func SetOutput(out io.Writer) {
	_std.SetOutput(out)
}

func Output() io.Writer {
	return _std.Output()
}

func SetLevel(level string) {
	_std.SetLevel(level)
}

// exported std api
func Debugf(format string, args ...interface{})   { _std.Debugf(format, args...) }
func Infof(format string, args ...interface{})    { _std.Infof(format, args...) }
func Warningf(format string, args ...interface{}) { _std.Warningf(format, args...) }
func Errorf(format string, args ...interface{})   { _std.Errorf(format, args...) }
func Panicf(format string, args ...interface{})   { _std.Panicf(format, args...) }
func Debug(args ...interface{})                   { _std.Debug(args...) }
func Info(args ...interface{})                    { _std.Info(args...) }
func Warning(args ...interface{})                 { _std.Warning(args...) }
func Error(args ...interface{})                   { _std.Error(args...) }
func Panic(args ...interface{})                   { _std.Panic(args...) }
func Debugln(args ...interface{})                 { _std.Debugln(args...) }
func Infoln(args ...interface{})                  { _std.Infoln(args...) }
func Warningln(args ...interface{})               { _std.Warningln(args...) }
func Errorln(args ...interface{})                 { _std.Errorln(args...) }
func Panicln(args ...interface{})                 { _std.Panicln(args...) }
*/
