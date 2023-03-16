package connector

import (
	"fmt"
	_ "io"

	jira "github.com/andygrunwald/go-jira"
)

func Foobar() {
	// jiraClient, _ := jira.NewClient(nil, "https://issues.apache.org/jira/")
	// issue, _, _ := jiraClient.Issue.Get("MESOS-3325", nil)

	// fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	// fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	// fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)

	// MESOS-3325: Running mesos-slave@0.23 in a container causes slave to be lost after a restart
	// Type: Bug
	// Priority: Critical

	jiraClient, _ := jira.NewClient(nil, "https://issues.apache.org/jira/")
	//var projectList jira.ProjectList
	projectList, _, _ := jiraClient.Project.GetList()

	fmt.Printf("%+v\n", projectList)
}
