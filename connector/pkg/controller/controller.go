package controller

import (
	"GoJira/pkg/connector"
	"GoJira/pkg/converter"
	"GoJira/pkg/structure"
	"GoJira/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

var DownloadAllProjects = func(w http.ResponseWriter, r *http.Request) {
	connector.DownloadProjects()
}

var DownloadProject = func(w http.ResponseWriter, r *http.Request) {
	connector.DownloadProject(r.FormValue("key"))
}

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
