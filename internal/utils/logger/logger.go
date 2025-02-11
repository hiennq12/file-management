package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func New() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
	}
}

func (l *Logger) Info(args ...interface{}) {
	l.infoLogger.Println(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.errorLogger.Println(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.errorLogger.Fatal(args...)
}
