package golib

import (
	"fmt"
	"log"
)

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

func LogWarn(message string) {
	log.Printf("WARN - %v", message)
}

func SaveLog(l LogInsert, c string) error {
	db := connect_db(c)
	_, err := db.Exec("INSERT INTO public.logs (details, path, method, created_at) VALUES($1, $2, $3, now());", l.Details, l.Path, l.Method)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("An error occured")
	}
	defer db.Close()
	return nil
}
