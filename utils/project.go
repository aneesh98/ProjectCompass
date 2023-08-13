package utils

type Project struct {
	ProjectName  string `json:"projectName"`
	Path         string `json:"path"`
	Description  string `json:"description"`
	PreferredIDE int    `json:"preferredIDE,omitempty"`
}

func (p *Project) SetProjectName(projectName string) {
	p.ProjectName = projectName
}

func (p *Project) SetPath(path string) {
	p.Path = ProcessPath(path)
}

func (p *Project) SetDescription(desc string) {
	p.Description = desc
}

func (p *Project) SetPreferredIDE(ideId int) {
	p.PreferredIDE = ideId
}
