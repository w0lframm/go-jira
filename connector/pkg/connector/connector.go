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

var connectorConfig structure.ConnectorConfig
var projects []structure.Project
var issuesMap = make(map[structure.Project][]structure.Issue)

func DownloadProjects() {
	f, _ := os.ReadFile("resources/config.yaml")

	_ = yaml.Unmarshal(f, &connectorConfig)

	resp, err := http.Get(connectorConfig.JiraURL + "/project")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &projects)

	DownloadIssues(projects[0])

	PushDataToDb()
}

func DownloadIssues(project structure.Project) {
	resp, err := http.Get(connectorConfig.JiraURL + "/search?jql=project=\"" + project.Name +
		"\"&expand=changelog&startAt=0&maxResults=" + strconv.Itoa(connectorConfig.IssuesCountInRequest))
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	var issueresp = new(structure.IssuesResponse)
	_ = json.Unmarshal(body, &issueresp)

	var issues []structure.Issue
	issues = append(issues, issueresp.Issues...)

	for i := 1; i <= issueresp.IssuesCount/connectorConfig.IssuesCountInRequest; i++ {
		resp, err := http.Get(connectorConfig.JiraURL + "/search?jql=project=\"" + project.Name +
			"\"&expand=changelog&startAt=" + strconv.Itoa(i*connectorConfig.IssuesCountInRequest) + "&maxResults=" +
			strconv.Itoa(connectorConfig.IssuesCountInRequest))
		if err != nil {
			log.Fatal(err)
		}
		body, _ := io.ReadAll(resp.Body)

		var issueresp = new(structure.IssuesResponse)
		_ = json.Unmarshal(body, &issueresp)

		issues = append(issues, issueresp.Issues...)
	}
	issuesMap[project] = issues
}

func PushDataToDb() {
	OpenDBConnection()
	for i := 0; i < len(projects); i++ {
		InsertIssuesIntoDB(issuesMap[projects[i]], InsertProjectIntoDB(projects[i]))
	}
}
