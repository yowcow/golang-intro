package mylogging

import (
	"fmt"
	"io"
	"log"
)

type Logger struct {
	callerDepth int
	logger      *log.Logger
}

func New(callerDepth int, w io.Writer, prefix string, flag int) *Logger {
	logger := log.New(w, prefix, flag)
	return &Logger{
		callerDepth,
		logger,
	}
}

func (l Logger) Debug(v ...interface{}) {
	v = append([]interface{}{"[DEBUG] "}, v...)
	l.Output(l.callerDepth+1, v...)
}

func (l Logger) Debugln(v ...interface{}) {
	v = append([]interface{}{"[DEBUG]"}, v...)
	l.Outputln(l.callerDepth+1, v...)
}

func (l Logger) Debugf(format string, v ...interface{}) {
	a := append([]interface{}{"[DEBUG]"}, fmt.Sprintf(format, v...))
	l.Outputln(l.callerDepth+1, a...)
}

func (l Logger) Info(v ...interface{}) {
	v = append([]interface{}{"[INFO] "}, v...)
	l.Output(l.callerDepth+1, v...)
}

func (l Logger) Infoln(v ...interface{}) {
	v = append([]interface{}{"[INFO]"}, v...)
	l.Outputln(l.callerDepth+1, v...)
}

func (l Logger) Infof(format string, v ...interface{}) {
	a := append([]interface{}{"[INFO]"}, fmt.Sprintf(format, v...))
	l.Outputln(l.callerDepth+1, a...)
}

func (l Logger) Error(v ...interface{}) {
	v = append([]interface{}{"[ERROR] "}, v...)
	l.Output(l.callerDepth+1, v...)
}

func (l Logger) Errorln(v ...interface{}) {
	v = append([]interface{}{"[ERROR]"}, v...)
	l.Outputln(l.callerDepth+1, v...)
}

func (l Logger) Errorf(format string, v ...interface{}) {
	a := append([]interface{}{"[ERROR]"}, fmt.Sprintf(format, v...))
	l.Outputln(l.callerDepth+1, a...)
}

func (l Logger) Output(callerDepth int, v ...interface{}) {
	l.logger.Output(callerDepth, fmt.Sprint(v...))
}

func (l Logger) Outputln(callerDepth int, v ...interface{}) {
	l.logger.Output(callerDepth, fmt.Sprintln(v...))
}
