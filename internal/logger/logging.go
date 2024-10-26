package logger

import (
	"log"
	"os"
)

var (
	infoLogger    = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger   = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	fatalLogger   = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Info(msg string) {
	infoLogger.Println(msg)
}

func Warning(msg string) {
	warningLogger.Println(msg)
}

func Error(msg string) {
	errorLogger.Println(msg)
}

func Fatal(msg string) {
	fatalLogger.Fatalln(msg)
}
