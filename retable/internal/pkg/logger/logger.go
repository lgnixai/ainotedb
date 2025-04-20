
package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[RETABLE] ", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Error(v ...interface{}) {
	l.Printf("ERROR: %v", v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.Printf("INFO: %v", v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.Printf("DEBUG: %v", v...)
}
