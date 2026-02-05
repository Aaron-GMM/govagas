package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	write   io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, ">[DEBUG]:", logger.Flags()),
		info:    log.New(writer, "$[INFO]:", logger.Flags()),
		warning: log.New(writer, "![WARNING]:", logger.Flags()),
		error:   log.New(writer, "X[ERROR]:", logger.Flags()),
		write:   writer,
	}
}

// create no formatted logs

func (L *Logger) Debug(v ...interface{}) {
	L.debug.Println(v)
}
func (L *Logger) Info(v ...interface{}) {
	L.debug.Println(v)
}
func (L *Logger) Warning(v ...interface{}) {
	L.warning.Println(v)
}
func (L *Logger) Error(v ...interface{}) {
	L.error.Println(v)
}

//create formatted logs

func (L *Logger) Debugf(format string, v ...interface{}) {
	L.debug.Printf(format, v)
}
func (L *Logger) InforF(format string, v ...interface{}) {
	L.info.Printf(format, v)
}
func (L *Logger) WarningF(format string, v ...interface{}) {
	L.warning.Printf(format, v)
}
func (L *Logger) ErrorF(format string, v ...interface{}) {
	L.error.Printf(format, v)
}
