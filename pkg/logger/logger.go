package logger

import (
	"log"
	"os"
	"sync"
	"time"
)

var (
	logFile *os.File
	once    sync.Once
)

func initLogger() {
	var err error
	logFile, err = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)
}

// ensureLogger ensures logger is initialized once
func ensureLogger() {
	once.Do(initLogger)
}

func Info(msg string) {
	ensureLogger()
	log.Println(formatLog("INFO", msg))
}

func Warn(msg string) {
	ensureLogger()
	log.Println(formatLog("WARN", msg))
}

func Error(msg string) {
	ensureLogger()
	log.Println(formatLog("ERROR", msg))
}

func Fatal(msg string) {
	ensureLogger()
	log.Fatalln(formatLog("FATAL", msg))
}

func formatLog(level string, msg string) string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "] [" + level + "] " + msg
}
