package converter

import "GoJira/pkg/structure"

func Convert(issue structure.Issue) structure.IssueDB {
	var result structure.IssueDB
	result.Key = issue.Key
	result.Summary = issue.Fields.Summary
	result.Priority = issue.Fields.Priority.Name
	result.Status = issue.Fields.Status.Name
	result.Type = issue.Fields.Type.Name
	return result
}
