package structure

type Project struct {
	tableName struct{} `pg:"project"`
	Id        int      `json:"id" pg:"id"`
	Key       string   `json:"key" pg:"key"`
	Name      string   `json:"name" pg:"name"`
	URL       string   `json:"self" pg:"url"`
}

type IssuesResponse struct {
	StartAt     int     `json:"startAt"`
	MaxResults  int     `json:"maxResults"`
	IssuesCount int     `json:"total"`
	Issues      []Issue `json:"issues"`
}

type Issue struct {
	Key       string      `json:"key"`
	Fields    IssueFields `json:"fields"`
	ChangeLog ChangeLog   `json:"changelog"`
}

type IssueDB struct {
	tableName  struct{} `pg:"issue"`
	Id         int      `pg:"id"`
	Key        string   `pg:"key"`
	ProjectId  int      `pg:"project_id"`
	CreatorId  int      `pg:"creator_id"`
	AssigneeId int      `pg:"assignee_id"`
	Summary    string   `pg:"summary"`
	Type       string   `pg:"type"`
	Status     string   `pg:"status"`
	Priority   string   `pg:"priority"`
	Created    string   `pg:"created"`
	Updated    string   `pg:"updated"`
	TimeSpent  int      `pg:"timespent"`
}

type Person struct {
	tableName   struct{} `pg:"author"`
	Id          int      `json:"id" pg:"id"`
	Key         string   `json:"key" pg:"key"`
	Name        string   `json:"name" pg:"name"`
	DisplayName string   `json:"displayName" pg:"display_name"`
}

type IssueFields struct {
	CreatedTime   string        `json:"created"`
	UpdatedTime   string        `json:"updated"`
	Summary       string        `json:"summary"`
	Type          IssueType     `json:"issuetype"`
	Status        IssueStatus   `json:"status"`
	Priority      IssuePriority `json:"priority"`
	ParentProject Project       `json:"project"`
	Creator       Person        `json:"creator"`
	Assignee      Person        `json:"reporter"`
	TimeSpent     int           `json:"timespent"`
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

type ChangeLog struct {
	Count     int       `json:"total"`
	Histories []History `json:"histories"`
}

type History struct {
	Author      Person                 `json:"author"`
	CreatedTime string                 `json:"created"`
	Items       []ChangeLogHistoryItem `json:"items"`
}

type ChangeLogHistoryItem struct {
	Field      string `json:"field"`
	FromString string `json:"fromString"`
	ToString   string `json:"toString"`
}

type StatusChangeDB struct {
	tableName   struct{} `pg:"status_change"`
	IssueId     int      `pg:"issue_id"`
	AuthorId    int      `pg:"author_id"`
	ChangedTime string   `pg:"changed"`
	FromStatus  string   `pg:"from_status"`
	ToStatus    string   `pg:"to_status"`
}
