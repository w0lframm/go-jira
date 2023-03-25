package structure

type Project struct {
	tableName struct{} `pg:"project"`
	Id        int      `json:"id"`
	Key       string   `json:"key"`
	Name      string   `json:"name"`
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

type IssueDB struct {
	tableName struct{} `pg:"issue"`
	Id        int      `json:"id"`
	Key       string   `json:"key"`
	ProjectId int      `json:"projectId"`
	CreatorId int      `json:"creatorId"`
	Summary   string   `json:"summary"`
	Type      string   `json:"type"`
	Status    string   `json:"status"`
	Priority  string   `json:"priority"`
}

type Person struct {
	tableName   struct{} `pg:"author"`
	Id          int      `json:"id"`
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
}

type IssueFields struct {
	//CreatedTime   UnixTime      `json:"created"`
	//UpdatedTime   UnixTime      `json:"updated"`
	Summary       string        `json:"summary"`
	Type          IssueType     `json:"issuetype"`
	Status        IssueStatus   `json:"status"`
	Priority      IssuePriority `json:"priority"`
	ParentProject Project       `json:"project"`
	Creator       Person        `json:"creator"`
}

//type UnixTime struct {
//	Time time.Time
//}
//
//func (u *UnixTime) Unmarshal(b []byte) error {
//	var timestamp int64
//	err := json.Unmarshal(b, &timestamp)
//	if err != nil {
//		return err
//	}
//	u.Time = time.Unix(timestamp, 0)
//	return nil
//}

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
