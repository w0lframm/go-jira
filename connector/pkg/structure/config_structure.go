package structure

type ConnectorConfig struct {
	JiraURL              string `yaml:"jiraURL"`
	IssuesCountInRequest int    `yaml:"issuesCountInRequest"`
	ThreadCount          int    `yaml:"threadCount"`
	MinWaitTime          int    `yaml:"minWaitTime"`
	MaxWaitTime          int    `yaml:"maxWaitTime"`
}

type DbConfig struct {
	PostgresUser     string `yaml:"postgresUser"`
	PostgresPassword string `yaml:"postgresPassword"`
	PostgresHost     string `yaml:"postgresHost"`
	PostgresPort     string `yaml:"postgresPort"`
	DbName           string `yaml:"dbName"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}
