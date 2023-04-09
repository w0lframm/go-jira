package connector

import (
	"GoJira/pkg/structure"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var connectorConfig structure.ConnectorConfig
var projects []structure.Project
var issuesMap = make(map[structure.Project][]structure.Issue)
var mutex sync.Mutex

// DownloadProjects downloads projects from Jira
// @Summary Downloads projects from Jira
// @Description Downloads projects from Jira
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /projects [get]
func DownloadProjects() {
	f, _ := os.ReadFile("resources/config.yaml")

	_ = yaml.Unmarshal(f, &connectorConfig)

	resp, err := http.Get(connectorConfig.JiraURL + "/project")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &projects)

	for i := 0; i < 5; i++ {
		DownloadIssues(projects[i])
		fmt.Println(len(issuesMap[projects[i]]))
	}

	PushDataToDb()
}

// DownloadIssues downloads issues for a given project from Jira
// @Summary Downloads issues for a given project from Jira
// @Description Downloads issues for a given project from Jira
// @Param projectName query string true "Project name"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /issues [get]
func DownloadIssues(project structure.Project) {
	resp, err := http.Get(connectorConfig.JiraURL + "/search?jql=project=\"" + project.Name +
		"\"&expand=changelog&startAt=0&maxResults=" + strconv.Itoa(connectorConfig.IssuesCountInRequest))
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	var issueresp = new(structure.IssuesResponse)
	_ = json.Unmarshal(body, &issueresp)

	issuesMap[project] = append(issuesMap[project], issueresp.Issues...)

	var requestsCount = issueresp.IssuesCount / connectorConfig.IssuesCountInRequest
	for i := 0; i < requestsCount/connectorConfig.ThreadCount; i++ {
		var wg sync.WaitGroup
		wg.Add(connectorConfig.ThreadCount)
		for j := 0; j < connectorConfig.ThreadCount; j++ {
			var startAt = (i*connectorConfig.ThreadCount + j + 1) * connectorConfig.IssuesCountInRequest
			go func() {
				DownloadIssuesThread(startAt, project)
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
			DownloadIssuesThread(startAt, project)
			wg.Done()
		}()
	}
	wg.Wait()
}

func DownloadIssuesThread(startAt int, project structure.Project) {
	resp, err := http.Get(connectorConfig.JiraURL + "/search?jql=project=\"" + project.Name +
		"\"&expand=changelog&startAt=" + strconv.Itoa(startAt) + "&maxResults=" +
		strconv.Itoa(connectorConfig.IssuesCountInRequest))
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)

	var issueresp = new(structure.IssuesResponse)
	_ = json.Unmarshal(body, &issueresp)

	mutex.Lock()
	issuesMap[project] = append(issuesMap[project], issueresp.Issues...)
	mutex.Unlock()
}

func PushDataToDb() {
	OpenDBConnection()
	for i := 0; i < len(projects); i++ {
		InsertIssuesIntoDB(issuesMap[projects[i]], InsertProjectIntoDB(projects[i]))
	}
}
