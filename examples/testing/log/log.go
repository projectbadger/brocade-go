package log

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/projectbadger/brocade-go/examples/testing/config"

	"github.com/sirupsen/logrus"
)

var (
	Log *Logger
	// serviceName string
	// configReference *config.ConfigApp
)

type Logger struct {
	Config  config.ConfigLog
	loggers []*logrus.Logger
	// maskedFields []string
}

func (l *Logger) FormatToString(v ...interface{}) string {
	return fmt.Sprint(v...)
}

func (l *Logger) LogOut(v ...interface{}) {
	log.Println(v...)
}

// LogWithFIeldsOut logs to all outputs at the specified
// level the provided formatted message.
func (l *Logger) LogWithFieldsOut(level logrus.Level, fields logrus.Fields, msg string, v ...interface{}) {
	if l.loggers == nil || len(l.loggers) < 1 {
		return
	}
	for _, logger := range l.loggers {
		switch level {
		case logrus.PanicLevel:
			logger.WithFields(fields).Panicf(msg, v...)
		case logrus.FatalLevel:
			logger.WithFields(fields).Fatalf(msg, v...)
		case logrus.ErrorLevel:
			logger.WithFields(fields).Errorf(msg, v...)
		case logrus.WarnLevel:
			logger.WithFields(fields).Warnf(msg, v...)
		case logrus.InfoLevel:
			logger.WithFields(fields).Infof(msg, v...)
		case logrus.DebugLevel:
			logger.WithFields(fields).Debugf(msg, v...)
		case logrus.TraceLevel:
			logger.WithFields(fields).Tracef(msg, v...)
		default:
			logger.WithFields(fields).Fatalf(msg, v...)
		}
	}
}

func (l *Logger) Info(ctx context.Context, msg string, v ...interface{}) {
	l.LogWithFieldsOut(logrus.InfoLevel, logrus.Fields{}, msg, v...)
}

func (l *Logger) Warn(ctx context.Context, msg string, v ...interface{}) {
	l.LogWithFieldsOut(logrus.WarnLevel, logrus.Fields{}, msg, v...)
}

func (l *Logger) Error(ctx context.Context, msg string, v ...interface{}) {
	l.LogWithFieldsOut(logrus.ErrorLevel, logrus.Fields{}, msg, v...)
}

func (l *Logger) Debug(ctx context.Context, msg string, v ...interface{}) {
	l.LogWithFieldsOut(logrus.DebugLevel, logrus.Fields{}, msg, v...)
}

func (l *Logger) Panic(ctx context.Context, msg string, v ...interface{}) {
	l.LogWithFieldsOut(logrus.PanicLevel, logrus.Fields{}, msg, v...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, v ...interface{}) {
	l.LogWithFieldsOut(logrus.FatalLevel, logrus.Fields{}, msg, v...)
}

// NewLogger returns a pointer to a new logger created
// with a LoggerConfig.
/*
Example:

  config := &LoggerConfig{Type: 1, Level: logrus.InfoLevel}
  logger := NewLogger(cfg)
*/
func NewLogger(cfg config.ConfigLog) (*Logger, error) {
	logger := &Logger{
		Config: cfg,
	}
	if cfg.Type&config.LogTypeMask <= 0 || cfg.Type <= 0 {
		// return nil, errors.LogInvalidType
		return nil, errors.New("invalid type: " + cfg.Type.String())
	}
	if cfg.Type&config.LogTypeFile > 0 {
		if cfg.FilePath == "" {
			return nil, errors.New("missing-filepath")
		}
		if cfg.ErrFilePath == "" {
			cfg.ErrFilePath = cfg.FilePath
		}
	}
	// logrus.PanicLevel == 0
	if cfg.Level <= logrus.PanicLevel {
		cfg.Level = logrus.ErrorLevel
	}
	err := logger.addLoggers()
	return logger, err
	// Output: *Logger, error
}

// SetupLogging sets up the logging outputs.
func SetupLogging() error {
	var err error
	Log, err = NewLogger(config.Cfg.Log)
	return err
}

func (l *Logger) addLoggers() error {
	if l.Config.Type&config.LogTypeStdOut > 0 {
		stdoutLogger := logrus.New()
		switch l.Config.Format {
		case "text", "txt", "":
			stdoutLogger.SetFormatter(&logrus.TextFormatter{})
		case "json", "js":
			stdoutLogger.SetFormatter(&logrus.JSONFormatter{})
		}
		stdoutLogger.SetLevel(l.Config.Level)
		stdoutLogger.Out = os.Stdout
		l.loggers = append(l.loggers, stdoutLogger)
	}
	if l.Config.Type&config.LogTypeFile > 0 {
		_, err := os.Stat(l.Config.FilePath)
		exists := false
		if err == os.ErrNotExist {
			exists = true
		}
		var f *os.File
		if exists {
			f, err = os.Open(l.Config.FilePath)
			if err != nil {
				return err
			}
		} else {
			f, err = os.Create(l.Config.FilePath)
			if err != nil {
				return err
			}
		}
		fileLogger := logrus.New()
		if l.Config.Format == "json" {
			fileLogger.SetFormatter(&logrus.JSONFormatter{})
		} else {
			fileLogger.SetFormatter(&logrus.TextFormatter{})
		}
		fileLogger.SetLevel(l.Config.Level)
		fileLogger.Out = f
		l.loggers = append(l.loggers, fileLogger)
	}
	return nil
}

func Info(msg string, v ...interface{}) {
	if Log != nil {
		Log.LogWithFieldsOut(logrus.InfoLevel, logrus.Fields{}, msg, v...)
	}
}

func Warn(msg string, v ...interface{}) {
	if Log != nil {
		Log.Warn(context.TODO(), msg, v...)
	}
}

func Error(msg string, v ...interface{}) {
	if Log != nil {
		Log.Error(context.TODO(), msg, v...)
	}
}

func Debug(msg string, v ...interface{}) {
	if Log != nil {
		Log.Debug(context.TODO(), msg, v...)
	}
}

func Trace(msg string, v ...interface{}) {
	if Log != nil {
		Log.Debug(context.TODO(), msg, v...)
	}
}

func Panic(msg string, v ...interface{}) {
	if Log != nil {
		Log.Panic(context.TODO(), msg, v...)
	}
}

func Fatal(msg string, v ...interface{}) {
	if Log != nil {
		Log.Fatal(context.TODO(), msg, v...)
	}
}
