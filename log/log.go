package log

import (
	"log"
)

func Error(err string) {
	log.Println("[ERROR] - MAILTRAP " + err)
}

func Info(err string) {
	log.Println("[INFO] - MAILTRAP " + err)
}
