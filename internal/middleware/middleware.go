package middleware

import (
	"fmt"
	"log"
	"net/http"
	"pmain2/internal/controller"
	"pmain2/pkg/logger"
)

var (
	lI *log.Logger
	lR *log.Logger
)

func INFO(text string) {
	if lI == nil {
		lI, _ = logger.New("middleware", logger.INFO)
	}
	lI.Println(text)
}

func ERROR(text string) {
	if lR == nil {
		lR, _ = logger.New("middleware", logger.ERROR)
	}
	lR.Println(text)
}

func BasicAuth(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username, password, ok := r.BasicAuth()
		if ok {
			c := controller.Init()
			isAuth, err := c.User.IsAuth(username, password)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusTeapot)
				controller.ERROR(err.Error())
				return
			}
			if isAuth {
				h.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		INFO(fmt.Sprint(r.URL))
		log.Println(r.URL)
		h.ServeHTTP(w, r)
	})
}

func CorsDisable(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token, Origin")
		h.ServeHTTP(w, r)
	})
}
