package golib

import "log"

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

func LogWarn(message string) {
	log.Printf("WARN - %v", message)
}
