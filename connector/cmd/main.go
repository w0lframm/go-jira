package main

import (
	"GoJira/pkg/controller"
	"GoJira/pkg/structure"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/swaggo/swag"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
)

// @title JIRA_analizer
// @version	1.0
// @description	Разработка промышленного клиент-серверного приложения с применением принципов микросервисной архитектуры, языков программирования Golang, фрейморка Angular и TypeScript.
// @BasePath /
// @host localhost:8081
// @license.name MIT License
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/updateAll", controller.DownloadAllProjects).Methods("POST")
	router.HandleFunc("/updateProject", controller.DownloadProject).Methods("POST")
	router.HandleFunc("/getProjects", controller.GetProjects).Methods("GET")

	var config structure.ServerConfig
	f, _ := os.ReadFile("resources/config.yaml")
	_ = yaml.Unmarshal(f, &config)

	err := http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		fmt.Print(err)
	}
}
