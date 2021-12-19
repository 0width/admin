package logger

import (
	"os"

	"git.xios.club/xios/gc"
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Level  string `value:"${log.level}"`
	Format string `value:"${log.format}"`
	File   string `value:"${log.file:=}"`
	Caller bool   `value:"${log.caller:=false}"`
}

func init() {
	gc.RegisterBeanFn(func(config LoggerConfig) *logrus.Logger {
		levels := map[string]logrus.Level{
			"panic": logrus.PanicLevel,
			"fatal": logrus.FatalLevel,
			"error": logrus.ErrorLevel,
			"warn":  logrus.WarnLevel,
			"info":  logrus.InfoLevel,
			"debug": logrus.DebugLevel,
			"trace": logrus.TraceLevel,
		}

		formatters := map[string]logrus.Formatter{
			"json": new(logrus.JSONFormatter),
			"text": new(logrus.TextFormatter),
		}

		logger := logrus.New()
		if config.File == "" {
			logger.SetOutput(os.Stdout)
		} else {
			file, err := os.OpenFile(config.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				panic("不能打开日志文件: " + err.Error())
			}
			logger.SetOutput(file)
		}

		if level, ok := levels[config.Level]; ok {
			logger.SetLevel(level)
		} else {
			logger.SetLevel(logrus.InfoLevel)
		}

		logger.ReportCaller = config.Caller

		if formatter, ok := formatters[config.Format]; ok {
			logger.SetFormatter(formatter)
		} else {
			logger.SetFormatter(&logrus.JSONFormatter{})
		}

		return logger
	})
}
