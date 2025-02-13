package logging

import (
	"os"

	"github.com/bnb-chain/greenfield-challenger/config"

	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// Logger instance for quick declarative logging levels
	Logger = logging.MustGetLogger("inscription-challenger")
	// log levels that are available
	levels = map[string]logging.Level{
		"CRITICAL": logging.CRITICAL,
		"ERROR":    logging.ERROR,
		"WARNING":  logging.WARNING,
		"NOTICE":   logging.NOTICE,
		"INFO":     logging.INFO,
		"DEBUG":    logging.DEBUG,
	}
)

// InitLogger initialises the logger.
func InitLogger(config *config.LogConfig) {
	backends := make([]logging.Backend, 0)

	if config.UseConsoleLogger {
		consoleFormat := logging.MustStringFormatter(`%{time:2006-01-02 15:04:05} %{level} %{shortfunc} %{message}`)
		consoleLogger := logging.NewLogBackend(os.Stdout, "", 0)
		consoleFormatter := logging.NewBackendFormatter(consoleLogger, consoleFormat)
		consoleLoggerLeveled := logging.AddModuleLevel(consoleFormatter)
		consoleLoggerLeveled.SetLevel(levels[config.Level], "")
		backends = append(backends, consoleLoggerLeveled)
	}

	if config.UseFileLogger {
		fileLogger := logging.NewLogBackend(&lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxFileSizeInMB,              // MaxSize is the maximum size in megabytes of the log file
			MaxBackups: config.MaxBackupsOfLogFiles,         // MaxBackups is the maximum number of old log files to retain
			MaxAge:     config.MaxAgeToRetainLogFilesInDays, // MaxAge is the maximum number of days to retain old log files
			Compress:   config.Compress,
		}, "", 0)
		fileFormat := logging.MustStringFormatter(`%{time:2006-01-02 15:04:05} %{level} %{shortfunc} %{message}`)
		fileFormatter := logging.NewBackendFormatter(fileLogger, fileFormat)
		fileLoggerLeveled := logging.AddModuleLevel(fileFormatter)
		fileLoggerLeveled.SetLevel(levels[config.Level], "")
		backends = append(backends, fileLoggerLeveled)
	}

	logging.SetBackend(backends...)
}
