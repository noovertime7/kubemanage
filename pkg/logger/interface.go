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

func NewErrorLoger() Logger {
	return logger{}
}

type logger struct{}

func (logger) Info(msg interface{}) {
	LG.Sugar().Info(msg)
}

func (logger) Infof(template string, args ...interface{}) {
	LG.Sugar().Infof(template, args)
}
func (logger) Warn(msg interface{}) {
	LG.Sugar().Warn(msg)
}

func (logger) Warnf(template string, args ...interface{}) {
	LG.Sugar().Warnf(template, args)
}

func (logger) ErrorWithCode(code int, err error) {
	msg := globalError.GetErrorMsg(code)
	LG.Error(msg, zap.Error(err))
}

func (logger) Error(msg interface{}) {
	LG.Sugar().Error(msg)
}

func (logger) ErrorWithErr(msg string, err error) {
	LG.Error(msg, zap.Error(err))
}
