package ilogger

import (
	"log"
	"os"
)

type FileLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

// NewFileLogger creates a new FileLogger writing to the given file path.
func NewFileLogger(filePath string) (*FileLogger, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileLogger{
		infoLogger:  log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(file, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile),
	}, nil
}

func (l *FileLogger) Info(msg string) {
	l.infoLogger.Println(msg)
}

func (l *FileLogger) Error(msg string) {
	l.errorLogger.Println(msg)
}

func (l *FileLogger) Debug(msg string) {
	l.debugLogger.Println(msg)
}
