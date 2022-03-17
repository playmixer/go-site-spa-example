package models

import (
	"log"
	"pmain2/pkg/logger"
	. "pmain2/pkg/utils"
	"strings"
)

var (
	lI *log.Logger
	lR *log.Logger
)

func INFO(text string) {
	if lI == nil {
		lI, _ = logger.New("dbase", logger.INFO)
	}
	lI.Println(text)
}

func ERROR(text string) {
	if lR == nil {
		lR, _ = logger.New("dbase", logger.ERROR)
	}
	lR.Println(text)
}

type SprDoct struct {
	Id       int    `json:"id"`
	Lname    string `json:"lname"`
	Fname    string `json:"fname"`
	Sname    string `json:"sname"`
	Password string `json:"-"`
}

func (m *SprDoct) ToUTF8() error {
	var err error
	m.Lname, err = ToUTF8(m.Lname)
	if err != err {
		return err
	}
	m.Fname, err = ToUTF8(m.Fname)
	if err != err {
		return err
	}
	m.Sname, err = ToUTF8(m.Sname)
	if err != err {
		return err
	}
	return nil
}

func (m *SprDoct) Trim() {
	m.Lname = strings.ReplaceAll(m.Lname, " ", "")
	m.Fname = strings.ReplaceAll(m.Fname, " ", "")
	m.Sname = strings.ReplaceAll(m.Sname, " ", "")
}
