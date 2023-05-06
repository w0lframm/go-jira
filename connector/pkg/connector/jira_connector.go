package connector

import (
	"GoJira/pkg/structure"
	"context"
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var connectorConfig structure.ConnectorConfig
var mutex sync.Mutex

func GetProjects() ([]structure.Project, error) {
	f, _ := os.ReadFile("resources/config.yaml")
	_ = yaml.Unmarshal(f, &connectorConfig)

	resp, err := retriableHttpRequest(connectorConfig.JiraURL + "/project")
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)

	var projects []structure.Project

	_ = json.Unmarshal(body, &projects)
	return projects, nil
}

func DownloadProjects() error {
	f, _ := os.ReadFile("resources/config.yaml")
	_ = yaml.Unmarshal(f, &connectorConfig)

	resp, err := retriableHttpRequest(connectorConfig.JiraURL + "/project")
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)

	var projects []structure.Project

	_ = json.Unmarshal(body, &projects)

	issuesMap := make(map[structure.Project][]structure.Issue)
	for i := 0; i < 2; i++ {
		issues, tmpErr := downloadIssues(projects[i])
		if tmpErr != nil {
			return tmpErr
		}
		issuesMap[projects[i]] = issues
	}
	err = InsertProjectsIntoDB(issuesMap)
	if err != nil {
		return err
	}
	return nil
}

func DownloadProject(projectKey string) error {
	f, _ := os.ReadFile("resources/config.yaml")
	_ = yaml.Unmarshal(f, &connectorConfig)

	resp, err := retriableHttpRequest(connectorConfig.JiraURL + "/project/" + projectKey)
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)

	var project structure.Project
	_ = json.Unmarshal(body, &project)

	issues, err := downloadIssues(project)
	if err != nil {
		return err
	}
	err = InsertProjectIntoDB(project, issues)
	if err != nil {
		return err
	}
	return nil
}

func downloadIssues(project structure.Project) ([]structure.Issue, error) {
	resp, err := retriableHttpRequest(connectorConfig.JiraURL + "/search?jql=project=\"" + project.Name +
		"\"&expand=changelog&startAt=0&maxResults=" + strconv.Itoa(connectorConfig.IssuesCountInRequest))
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(resp.Body)

	var issueresp = new(structure.IssuesResponse)
	_ = json.Unmarshal(body, &issueresp)

	var issues = issueresp.Issues

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var requestsCount = issueresp.IssuesCount / connectorConfig.IssuesCountInRequest
	for i := 0; i < requestsCount/connectorConfig.ThreadCount; i++ {
		var wg sync.WaitGroup
		wg.Add(connectorConfig.ThreadCount)
		for j := 0; j < connectorConfig.ThreadCount; j++ {
			var startAt = (i*connectorConfig.ThreadCount + j + 1) * connectorConfig.IssuesCountInRequest
			go func() {
				defer wg.Done()
				var result, tmpErr = downloadIssuesThread(ctx, startAt, project)
				if tmpErr != nil {
					cancel()
					err = tmpErr
					return
				}
				mutex.Lock()
				defer mutex.Unlock()
				issues = append(issues, result...)
			}()
		}
		wg.Wait()
	}
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(requestsCount % connectorConfig.ThreadCount)
	for i := 1; i <= requestsCount%connectorConfig.ThreadCount; i++ {
		var startAt = (requestsCount - requestsCount%connectorConfig.ThreadCount + i) *
			connectorConfig.IssuesCountInRequest
		go func() {
			defer wg.Done()
			var result, tmpErr = downloadIssuesThread(ctx, startAt, project)
			if tmpErr != nil {
				cancel()
				err = tmpErr
				return
			}
			mutex.Lock()
			defer mutex.Unlock()
			issues = append(issues, result...)
		}()
	}
	wg.Wait()
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func downloadIssuesThread(ctx context.Context, startAt int, project structure.Project) ([]structure.Issue, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		resp, err := retriableHttpRequest(connectorConfig.JiraURL + "/search?jql=project=\"" + project.Name +
			"\"&expand=changelog&startAt=" + strconv.Itoa(startAt) + "&maxResults=" +
			strconv.Itoa(connectorConfig.IssuesCountInRequest))
		if err != nil {
			return nil, err
		}
		body, _ := io.ReadAll(resp.Body)

		var issueresp = new(structure.IssuesResponse)
		_ = json.Unmarshal(body, &issueresp)

		return issueresp.Issues, nil
	}
}

func retriableHttpRequest(url string) (*http.Response, error) {
	retries := getMaxRetries()
	for retries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			return resp, nil
		}
	}
	return nil, errors.New("")
}

func getMaxRetries() int {
	result := 2
	multiplier := 2
	for connectorConfig.MinWaitTime*multiplier < connectorConfig.MaxWaitTime {
		result += 1
		multiplier *= 2
	}
	return result
}
