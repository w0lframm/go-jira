package structure

type RestProject struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RestProjects struct {
	Projects  []RestProject `json:"projects"`
	PageCount int           `json:"page_count"`
}
