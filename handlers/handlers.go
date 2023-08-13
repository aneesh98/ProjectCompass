package handlers

import (
	"bufio"
	"fmt"
	"os"
	"project-tracker/storage"
	"project-tracker/utils"
	"strconv"
	"strings"
)

func RegisterIde() {
	options := []string{"Enter Name For your IDE", "Enter IDE Executable Path"}
	ide := utils.IDE{}
	for i := 0; i < len(options); i++ {
		fmt.Printf("%d. %s\n", i, options[i])
		switch i {
		case 0:
			{
				var ideName string
				fmt.Scanln(&ideName)
				ide.SetIdeName(ideName)
			}
		case 1:
			{
				var idePath string
				fmt.Scanln(&idePath)
				ide.SetExecutablePath(idePath)
			}
		}
	}
	storage.SaveIde(ide)
}

func RegisterProject() {
	project := utils.Project{}
	options := []string{"Enter Project Name", "Enter Project Path", "Enter Project Description", "Enter Preferred IDE"}
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(options); i++ {
		fmt.Printf("%d. %s\n", i+1, options[i])
		switch i {
		case 0:
			{
				projectName, _ := reader.ReadString('\n')
				project.SetProjectName(strings.TrimSpace(projectName))
			}
		case 1:
			{
				projectPath, _ := reader.ReadString('\n')
				project.SetPath(strings.TrimSpace(projectPath))
			}
		case 2:
			{
				projectDescription, _ := reader.ReadString('\n')
				project.SetDescription(strings.TrimSpace(projectDescription))
			}
		case 3:
			{
				fmt.Println("Available IDEs: ")
				storage.ListIdes()
				projectIDEidStr, _ := reader.ReadString('\n')
				projectIDEid, err := strconv.Atoi(strings.TrimSpace(projectIDEidStr))
				if err != nil {
					fmt.Println("Please enter a valid number for the IDE id.")
					i-- // Decrement i to repeat this step
					continue
				}
				project.SetPreferredIDE(projectIDEid)
			}
		}
	}
	storage.SaveProject(project)
}
