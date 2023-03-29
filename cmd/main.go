package main

import (
	"GoJira/pkg/connector"

	_ "github.com/swaggo/swag"
)

//	@title			JIRA_analizer
//	@version		1.0
//	@description	Разработка промышленного клиент-серверного приложения с применением принципов микросервисной архитектуры, языков программирования Golang, фрейморка Angular и TypeScript.
//  @BasePath /

//	@license.name	MIT License

func main() {
	connector.DownloadProjects()
}
