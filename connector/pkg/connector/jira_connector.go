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
	"sync"
)

var connectorConfig structure.ConnectorConfig
var mutex sync.Mutex

func DownloadProjects() {
	f, _ := os.ReadFile("resources/config.yaml")
	_ = yaml.Unmarshal(f, &connectorConfig)

	resp, err := http.Get(connectorConfig.JiraURL + "/project")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	var projects []structure.Project

	_ = json.Unmarshal(body, &projects)

	for i := 0; i < 2; i++ {
		var issues = DownloadIssues(projects[i])
		PushProjectToDb(projects[i], issues)
	}
}

func DownloadProject(projectKey string) {
	f, _ := os.ReadFile("resources/config.yaml")
	_ = yaml.Unmarshal(f, &connectorConfig)

	resp, err := http.Get(connectorConfig.JiraURL + "/project/" + projectKey)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	var project structure.Project
	_ = json.Unmarshal(body, &project)

	var issues = DownloadIssues(project)
	PushProjectToDb(project, issues)
}

func DownloadIssues(project structure.Project) []structure.Issue {
	resp, err := http.Get(connectorConfig.JiraURL + "/search?jql=project=\"" + project.Name +
		"\"&expand=changelog&startAt=0&maxResults=" + strconv.Itoa(connectorConfig.IssuesCountInRequest))
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	var issueresp = new(structure.IssuesResponse)
	_ = json.Unmarshal(body, &issueresp)

	var issues = issueresp.Issues

	var requestsCount = issueresp.IssuesCount / connectorConfig.IssuesCountInRequest
	for i := 0; i < requestsCount/connectorConfig.ThreadCount; i++ {
		var wg sync.WaitGroup
		wg.Add(connectorConfig.ThreadCount)
		for j := 0; j < connectorConfig.ThreadCount; j++ {
			var startAt = (i*connectorConfig.ThreadCount + j + 1) * connectorConfig.IssuesCountInRequest
			go func() {
				var result = DownloadIssuesThread(startAt, project)
				mutex.Lock()
				issues = append(issues, result...)
				mutex.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()
	}
	var wg sync.WaitGroup
	wg.Add(requestsCount % connectorConfig.ThreadCount)
	for i := 1; i <= requestsCount%connectorConfig.ThreadCount; i++ {
		var startAt = (requestsCount - requestsCount%connectorConfig.ThreadCount + i) *
			connectorConfig.IssuesCountInRequest
		go func() {
			var result = DownloadIssuesThread(startAt, project)
			mutex.Lock()
			issues = append(issues, result...)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return issues
}

func DownloadIssuesThread(startAt int, project structure.Project) []structure.Issue {
	resp, err := http.Get(connectorConfig.JiraURL + "/search?jql=project=\"" + project.Name +
		"\"&expand=changelog&startAt=" + strconv.Itoa(startAt) + "&maxResults=" +
		strconv.Itoa(connectorConfig.IssuesCountInRequest))
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	var issueresp = new(structure.IssuesResponse)
	_ = json.Unmarshal(body, &issueresp)

	return issueresp.Issues
}

func PushProjectToDb(project structure.Project, issues []structure.Issue) {
	ClearProject(project)
	InsertProjectIntoDB(project, issues)
}
