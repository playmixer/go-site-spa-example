package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pmain2/internal/database"
	"pmain2/internal/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	model := models.SprDoctModel{Db: conn.DB}
	data, err := model.Get(613)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	//LoggerInfo.Println("test")
	fmt.Fprintf(w, string(res))
}
