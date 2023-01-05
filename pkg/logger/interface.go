package logger

import (
	"go.uber.org/zap"

	"github.com/noovertime7/kubemanage/pkg/globalError"
)

type Logger interface {
	Info(msg interface{})
	Infof(template string, args ...interface{})
	Warn(msg interface{})
	Warnf(template string, args ...interface{})
	Error(msg interface{})
	ErrorWithCode(code int, err error)
	ErrorWithErr(msg string, err error)
}

func New(l *zap.Logger) Logger {
	return &logger{lg: l}
}

type logger struct {
	lg *zap.Logger
}

func (l *logger) Info(msg interface{}) {
	l.lg.Sugar().Info(msg)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.lg.Sugar().Infof(template, args)
}
func (l *logger) Warn(msg interface{}) {
	l.lg.Sugar().Warn(msg)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.lg.Sugar().Warnf(template, args)
}

func (l *logger) ErrorWithCode(code int, err error) {
	msg := globalError.GetErrorMsg(code)
	l.lg.Error(msg, zap.Error(err))
}

func (l *logger) Error(msg interface{}) {
	l.lg.Sugar().Error(msg)
}

func (l *logger) ErrorWithErr(msg string, err error) {
	l.lg.Error(msg, zap.Error(err))
}
