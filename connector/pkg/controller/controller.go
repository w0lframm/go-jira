package controller

import (
	"GoJira/pkg/connector"
	"GoJira/pkg/converter"
	"GoJira/pkg/structure"
	"GoJira/pkg/utils"
	_ "github.com/swaggo/swag"
	"net/http"
	"strconv"
	"strings"
)

// DownloadAllProjects downloads all projects from Jira
// @Summary Downloads all projects from Jira
// @Description Downloads all projects from Jira
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /updateAll [post]
var DownloadAllProjects = func(w http.ResponseWriter, r *http.Request) {
	connector.DownloadProjects()
}

// DownloadProject downloads project with given key from Jira
// @Summary Downloads project with given key from Jira
// @Description Downloads project with given key from Jira
// @Param key query string true "Project's key"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /updateProject [post]
var DownloadProject = func(w http.ResponseWriter, r *http.Request) {
	connector.DownloadProject(r.FormValue("key"))
}

// GetProjects retrieving all projects
// @Summary Retrieving all projects
// @Description Retrieving all projects
// @Param limit query integer true "Projects count on one page"
// @Param page query integer true "Number of page"
// @Param search query string true "Search for project name"
// @Success 200 {object} structure.RestProjects
// @Router /getProjects [get]
var GetProjects = func(w http.ResponseWriter, r *http.Request) {
	var limit, _ = strconv.Atoi(r.FormValue("limit"))
	var page, _ = strconv.Atoi(r.FormValue("page"))
	var search = r.FormValue("search")
	var projects = connector.GetProjects()

	if search != "" {
		var temp []structure.Project
		for i := 0; i < len(projects); i++ {
			if strings.HasPrefix(strings.ToLower(projects[i].Name), strings.ToLower(search)) {
				temp = append(temp, projects[i])
			}
		}
		projects = temp
	}
	var pageCount = len(projects)/limit + 1

	if page*limit > len(projects) {
		projects = projects[(page-1)*limit:]
	} else {
		projects = projects[(page-1)*limit : page*limit]
	}

	var result = converter.ConvertProjectsToRestProjects(projects)
	result.PageCount = pageCount

	utils.Respond(w, result)
}
