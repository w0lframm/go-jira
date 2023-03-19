package connector

import (
	"GoJira/pkg/structure"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var config map[string]string
var projects []structure.Project
var issues []structure.Issue

func DownloadProjects() {
	f, _ := os.ReadFile("resources/config.yaml")

	_ = yaml.Unmarshal(f, &config)

	resp, err := http.Get(config["jiraURL"] + "/project")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &projects)

	DownloadIssues(projects[0].Name)
	//for i := 0; i < len(projects); i++ {
	//	DownloadIssues(projects[i].Name)
	//}
}

func DownloadIssues(projectName string) {
	resp, err := http.Get(config["jiraURL"] + "/search?jql=project=\"" + projectName +
		"\"&expand=changelog&startAt=0&maxResults=" + config["issuesCountInRequest"])
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	var issueresp = new(structure.IssuesResponse)
	_ = json.Unmarshal(body, &issueresp)

	issues = append(issues, issueresp.Issues...)

	countInRequest, _ := strconv.Atoi(config["issuesCountInRequest"])
	for i := 1; i <= issueresp.IssuesCount/countInRequest; i++ {
		resp, err := http.Get(config["jiraURL"] + "/search?jql=project=\"" + projectName +
			"\"&expand=changelog&startAt=" + strconv.Itoa(i*countInRequest) + "&maxResults=" +
			config["issuesCountInRequest"])
		if err != nil {
			log.Fatal(err)
		}
		body, _ := io.ReadAll(resp.Body)

		var issueresp = new(structure.IssuesResponse)
		_ = json.Unmarshal(body, &issueresp)

		issues = append(issues, issueresp.Issues...)
	}
}
