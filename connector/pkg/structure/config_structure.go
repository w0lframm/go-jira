package structure

type ConnectorConfig struct {
	JiraURL              string `yaml:"jiraURL"`
	IssuesCountInRequest int    `yaml:"issuesCountInRequest"`
}

type DbConfig struct {
	PostgresUser     string `yaml:"postgresUser"`
	PostgresPassword string `yaml:"postgresPassword"`
	PostgresHost     string `yaml:"postgresHost"`
	PostgresPort     string `yaml:"postgresPort"`
	DbName           string `yaml:"dbName"`
}
