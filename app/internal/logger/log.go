package logger

import (
	"log"

	"gitlab.local/dhamith93/devops-playground/app/internal/color"
)

var (
	info  = color.Green
	warn  = color.Yellow
	error = color.Red
)

func Info(msg string) {
	log.Println("INFO: " + info(msg))
}

func Warn(msg string) {
	log.Println("WARN: " + warn(msg))
}

func Error(msg string) {
	log.Println("ERR: " + error(msg))
}
