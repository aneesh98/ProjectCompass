package constants

import "project-tracker/utils"

var (
	IdeFilePath     = utils.ProcessPath("~/.projects/ides.json")
	ProjectFilePath = utils.ProcessPath("~/.projects/projects.json")
	RootPath        = utils.ProcessPath("~/.projects")
)
