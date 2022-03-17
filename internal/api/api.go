package api

import (
	"log"
	"pmain2/pkg/logger"
)

var (
	lI *log.Logger
	lR *log.Logger
)

func INFO(text string) {
	if lI == nil {
		lI, _ = logger.New("api", logger.INFO)
	}
	lI.Println(text)
}

func ERROR(text string) {
	if lR == nil {
		lR, _ = logger.New("api", logger.ERROR)
	}
	lR.Println(text)
}
