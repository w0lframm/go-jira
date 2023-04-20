package converter

import "GoJira/pkg/structure"

func ConvertProjectsToRestProjects(projects []structure.Project) structure.RestProjects {
	var result structure.RestProjects

	for i := 0; i < len(projects); i++ {
		var project structure.RestProject
		project.Name = projects[i].Name
		project.Key = projects[i].Key
		project.URL = projects[i].URL
		result.Projects = append(result.Projects, project)
	}

	return result
}
