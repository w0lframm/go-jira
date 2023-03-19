package structure

type Project struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type IssuesResponse struct {
	StartAt     int     `json:"startAt"`
	MaxResults  int     `json:"maxResults"`
	IssuesCount int     `json:"total"`
	Issues      []Issue `json:"issues"`
}

type Issue struct {
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type IssueFields struct {
	//CreatedTime time.Time     `json:"created"`
	//UpdatedTime time.Time     `json:"updated"`
	Type          IssueType     `json:"issuetype"`
	Status        IssueStatus   `json:"status"`
	Priority      IssuePriority `json:"priority"`
	ParentProject Project       `json:"project"`
}

type IssueType struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type IssueStatus struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type IssuePriority struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
