package connector

import (
	"GoJira/pkg/converter"
	"GoJira/pkg/structure"
	"github.com/go-pg/pg/v10"
	"gopkg.in/yaml.v2"
	"os"
)

var dbConfig structure.DbConfig
var db *pg.DB

func InsertProjectsIntoDB(issuesMap map[structure.Project][]structure.Issue) error {
	openDBConnection()
	err := db.RunInTransaction(db.Context(), func(tx *pg.Tx) error {
		for key, value := range issuesMap {
			tmpErr := insertProjectIntoDB(tx, key, value)
			if tmpErr != nil {
				return tmpErr
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func InsertProjectIntoDB(project structure.Project, issues []structure.Issue) error {
	openDBConnection()
	err := db.RunInTransaction(db.Context(), func(tx *pg.Tx) error {
		tmpErr := clearProject(project, tx)
		if tmpErr != nil {
			return tmpErr
		}
		_, tmpErr = tx.Model(&project).Returning("id").Insert()
		if tmpErr != nil {
			return tmpErr
		}
		tmpErr = insertIssuesIntoDB(issues, project, tx)
		if tmpErr != nil {
			return tmpErr
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func insertProjectIntoDB(tx *pg.Tx, project structure.Project, issues []structure.Issue) error {
	err := clearProject(project, tx)
	if err != nil {
		return err
	}
	_, err = tx.Model(&project).Returning("id").Insert()
	if err != nil {
		return err
	}
	err = insertIssuesIntoDB(issues, project, tx)
	if err != nil {
		return err
	}
	return nil
}

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

func insertIssuesIntoDB(issues []structure.Issue, project structure.Project, tx *pg.Tx) error {
	for i := 0; i < len(issues); i++ {
		creator, err := insertAuthorIntoDB(issues[i].Fields.Creator, tx)
		if err != nil {
			return err
		}
		assignee, err := insertAuthorIntoDB(issues[i].Fields.Assignee, tx)
		if err != nil {
			return err
		}
		issue := converter.ConvertIssueToIssueDB(issues[i])
		issue.ProjectId = project.Id
		issue.CreatorId = creator.Id
		issue.AssigneeId = assignee.Id
		_, err = tx.Model(&issue).Returning("id").Insert()
		if err != nil {
			return err
		}
		for j := 0; j < len(issues[i].ChangeLog.Histories); j++ {
			statusChange := converter.ConvertHistoryToStatusChangeDB(issues[i].ChangeLog.Histories[j])
			if (structure.StatusChangeDB{}) != statusChange {
				statusChange.IssueId = issue.Id
				author, err := insertAuthorIntoDB(issues[i].ChangeLog.Histories[j].Author, tx)
				if err != nil {
					return err
				}
				statusChange.AuthorId = author.Id
				_, err = tx.Model(&statusChange).Insert()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func insertAuthorIntoDB(author structure.Person, tx *pg.Tx) (structure.Person, error) {
	var result structure.Person
	err := tx.Model(&result).Where("key='" + author.Key + "'").Select()
	if err != nil {
		_, err := tx.Model(&author).Returning("id").Insert()
		if err != nil {
			return structure.Person{}, err
		}
		return author, nil
	}
	return result, nil
}

func clearProject(project structure.Project, tx *pg.Tx) error {
	_, err := tx.Model(&project).Where("key='" + project.Key + "'").Delete()
	if err != nil {
		return err
	}
	return nil
}
