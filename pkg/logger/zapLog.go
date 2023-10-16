package logger

import (
	"go.uber.org/zap"
	"log"
)

type Loger interface {
	InitLog()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type apiLogger struct {
	sugar *zap.SugaredLogger
}

func NewLogger() *apiLogger {
	return &apiLogger{}
}

func (l *apiLogger) InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Println("zap logger failure: ", err.Error())
	}

	l.sugar = logger.Sugar()

	if err := l.sugar.Sync(); err != nil {
		l.sugar.Error(err)
	}
}

func (l *apiLogger) Debug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *apiLogger) Debugf(template string, args ...interface{}) {
	l.sugar.Debugf(template, args...)
}

func (l *apiLogger) Info(args ...interface{}) {
	l.sugar.Info(args...)
}

func (l *apiLogger) Infof(template string, args ...interface{}) {
	l.sugar.Infof(template, args...)
}

func (l *apiLogger) Warn(args ...interface{}) {
	l.sugar.Warn(args...)
}

func (l *apiLogger) Warnf(template string, args ...interface{}) {
	l.sugar.Warnf(template, args...)
}

func (l *apiLogger) Error(args ...interface{}) {
	l.sugar.Error(args...)
}

func (l *apiLogger) Errorf(template string, args ...interface{}) {
	l.sugar.Errorf(template, args...)
}

func (l *apiLogger) DPanic(args ...interface{}) {
	l.sugar.DPanic(args...)
}

func (l *apiLogger) DPanicf(template string, args ...interface{}) {
	l.sugar.DPanicf(template, args...)
}

func (l *apiLogger) Panic(args ...interface{}) {
	l.sugar.Panic(args...)
}

func (l *apiLogger) Panicf(template string, args ...interface{}) {
	l.sugar.Panicf(template, args...)
}

func (l *apiLogger) Fatal(args ...interface{}) {
	l.sugar.Fatal(args...)
}

func (l *apiLogger) Fatalf(template string, args ...interface{}) {
	l.sugar.Fatalf(template, args...)
}
