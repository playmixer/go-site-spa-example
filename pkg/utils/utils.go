package utils

import (
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

func ToUTF8(text string) (string, error) {
	sr := strings.NewReader(text)
	tr := transform.NewReader(sr, charmap.Windows1251.NewDecoder())
	buf, err := ioutil.ReadAll(tr)
	if err != err {
		return "", err
	}

	text = string(buf) // строка в UTF-8
	return text, nil
}
