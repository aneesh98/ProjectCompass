package main

import (
	"fmt"
	"project-tracker/constants"
	"project-tracker/handlers"
	"project-tracker/storage"
	"project-tracker/utils"
	"strconv"
)

func initFn() {
	for _, path := range []string{constants.IdeFilePath, constants.ProjectFilePath} {
		utils.CreateFileWithPath(path)
	}
}

func main() {
	initFn()
	options := []string{"Register Project", "Register IDE", "List Available IDEs/Editors", "Search Project By Name", "Search Project By Description", "Start Project"}
	menuStr := "Select Option:"
	for i := 0; i < len(options); i++ {
		menuStr += "\n\t" + strconv.Itoa(i+1) + ". " + options[i]
	}
	fmt.Println(menuStr)
	var option int
	fmt.Scanf("%d", &option)
	switch option {
	case 1:
		{
			handlers.RegisterProject()
		}
	case 2:
		{
			handlers.RegisterIde()
		}
	case 3:
		{
			storage.ListIdes()
		}
	case 6:
		{
			storage.ListProjectsToStart()
		}
	}
}
