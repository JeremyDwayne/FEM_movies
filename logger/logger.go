package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	file        *os.File
}

func NewLogger(logFilePath string) (*Logger, error) {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		file:        file,
	}, nil
}

func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

func (l *Logger) Error(message string, err error) {
	l.errorLogger.Printf("%s: %v\n", message, err)
}

func (l *Logger) Close() error {
	return l.file.Close()
}
