package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"pmain2/internal/api"
	"pmain2/internal/config"
	"pmain2/internal/middleware"
	"pmain2/internal/server"
)

func indexServe(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(path)
	tmpl.Execute(w, nil)
}

func routesFrontend(router *mux.Router, handle http.HandlerFunc) {
	path, _ := os.Getwd()
	file, _ := os.Open(path + "./.froutes")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		router.HandleFunc(url, handle).Methods(http.MethodGet)
	}
}

func main() {
	conf, err := config.Create()
	if err != nil {
		fmt.Println(err)
	}

	server := server.Create(conf)
	apiRouter := server.Router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.BasicAuth)
	apiRouter.Use(middleware.Logging)
	apiRouter.Use(middleware.CorsDisable)
	apiRouter.HandleFunc("/user/{id}/", api.GetUser).Methods(http.MethodGet)

	webRouter := server.Router.PathPrefix("/").Subrouter()
	webRouter.PathPrefix("/static/").Handler(http.FileServer(http.Dir("./"))).Methods(http.MethodGet)
	webRouter.Use(middleware.Logging)

	routesFrontend(webRouter, indexServe)

	err = server.Run()
	if err != nil {
		fmt.Println(err)
	}
}
