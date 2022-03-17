package logger

import (
	"log"
	"os"
)

const (
	INFO    = "INFO"
	ERROR   = "ERROR"
	WARNING = "WARNING"
)

func New(name, lvl string) (*log.Logger, error) {
	path, err := os.Getwd()
	filename := path + "\\logs\\" + name + ".log"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return log.New(file, lvl+": ", log.Ldate|log.Ltime|log.Lshortfile), nil
}
