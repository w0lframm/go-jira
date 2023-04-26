package connector

import (
	"GoJira/pkg/converter"
	"GoJira/pkg/structure"
	"github.com/go-pg/pg/v10"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var dbConfig structure.DbConfig
var db *pg.DB

func openDBConnection() {
	if (structure.DbConfig{}) == dbConfig {
		f, _ := os.ReadFile("resources/config.yaml")

		_ = yaml.Unmarshal(f, &dbConfig)
	}

	if db == nil {
		db = pg.Connect(&pg.Options{
			Addr:     dbConfig.PostgresHost + ":" + dbConfig.PostgresPort,
			User:     dbConfig.PostgresUser,
			Password: dbConfig.PostgresPassword,
			Database: dbConfig.DbName,
		})
	}
}

func InsertProjectIntoDB(project structure.Project, issues []structure.Issue) {
	openDBConnection()
	_, err := db.Model(&project).Returning("id").Insert()
	if err != nil {
		log.Fatal(err)
	}
	insertIssuesIntoDB(issues, project)
}

func insertIssuesIntoDB(issues []structure.Issue, project structure.Project) {
	for i := 0; i < len(issues); i++ {
		creator := insertAuthorIntoDB(issues[i].Fields.Creator)
		assignee := insertAuthorIntoDB(issues[i].Fields.Assignee)
		issue := converter.ConvertIssueToIssueDB(issues[i])
		issue.ProjectId = project.Id
		issue.CreatorId = creator.Id
		issue.AssigneeId = assignee.Id
		_, err := db.Model(&issue).Returning("id").Insert()
		if err != nil {
			log.Fatal(err)
		}
		for j := 0; j < len(issues[i].ChangeLog.Histories); j++ {
			statusChange := converter.ConvertHistoryToStatusChangeDB(issues[i].ChangeLog.Histories[j])
			if (structure.StatusChangeDB{}) != statusChange {
				statusChange.IssueId = issue.Id
				statusChange.AuthorId = insertAuthorIntoDB(issues[i].ChangeLog.Histories[j].Author).Id
				_, err := db.Model(&statusChange).Insert()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func insertAuthorIntoDB(author structure.Person) structure.Person {
	var result structure.Person
	err := db.Model(&result).Where("key='" + author.Key + "'").Select()
	if err != nil {
		_, err := db.Model(&result).Returning("id").Insert()
		if err != nil {
			log.Fatal(err)
		}
		return result
	}
	return result
}

func GetProjects() []structure.Project {
	openDBConnection()
	var result []structure.Project
	err := db.Model(&result).Select()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func ClearProject(project structure.Project) {
	openDBConnection()
	_, err := db.Model(&project).Where("key='" + project.Key + "'").Delete()
	if err != nil {
		log.Fatal(err)
	}
}
