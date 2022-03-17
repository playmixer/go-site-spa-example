package controller

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
		lI, _ = logger.New("controller", logger.INFO)
	}
	lI.Println(text)
}

func ERROR(text string) {
	if lR == nil {
		lR, _ = logger.New("controller", logger.ERROR)
	}
	lR.Println(text)
}

type Controller struct {
	User *user
}

func Init() *Controller {
	return &Controller{
		User: createUser(),
	}
}
