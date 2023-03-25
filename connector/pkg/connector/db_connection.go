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

func OpenDBConnection() {
	f, _ := os.ReadFile("resources/config.yaml")

	_ = yaml.Unmarshal(f, &dbConfig)

	db = pg.Connect(&pg.Options{
		Addr:     dbConfig.PostgresHost + ":" + dbConfig.PostgresPort,
		User:     dbConfig.PostgresUser,
		Password: dbConfig.PostgresPassword,
		Database: dbConfig.DbName,
	})
}

func InsertProjectIntoDB(project structure.Project) structure.Project {
	_, err := db.Model(&project).Returning("id").Insert()
	if err != nil {
		log.Fatal(err)
	}
	return project
}

func InsertIssuesIntoDB(issues []structure.Issue, project structure.Project) {
	for i := 0; i < len(issues); i++ {
		creator := InsertAuthorIntoDB(issues[i].Fields.Creator)
		issue := converter.Convert(issues[i])
		issue.ProjectId = project.Id
		issue.CreatorId = creator.Id
		_, err := db.Model(&issue).Insert()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func InsertAuthorIntoDB(author structure.Person) structure.Person {
	var result structure.Person
	err := db.Model(&result).Where("key='" + author.Key + "'").Select()
	if err != nil {
		_, err := db.Model(&author).Returning("id").Insert()
		if err != nil {
			log.Fatal(err)
		}
		return author
	}
	return result
}
