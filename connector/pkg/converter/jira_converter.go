package converter

import (
	"GoJira/pkg/structure"
)

func ConvertIssueToIssueDB(issue structure.Issue) structure.IssueDB {
	var result structure.IssueDB
	result.Key = issue.Key
	result.Summary = issue.Fields.Summary
	result.Priority = issue.Fields.Priority.Name
	result.Status = issue.Fields.Status.Name
	result.Type = issue.Fields.Type.Name
	result.Created = issue.Fields.CreatedTime
	result.Updated = issue.Fields.UpdatedTime
	result.TimeSpent = issue.Fields.TimeSpent
	return result
}

func ConvertHistoryToStatusChangeDB(history structure.History) structure.StatusChangeDB {
	for i := 0; i < len(history.Items); i++ {
		if history.Items[i].Field == "status" {
			var result structure.StatusChangeDB
			result.ChangedTime = history.CreatedTime
			result.FromStatus = history.Items[i].FromString
			result.ToStatus = history.Items[i].ToString
			return result
		}
	}
	return structure.StatusChangeDB{}
}
