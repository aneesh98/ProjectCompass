package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"project-tracker/constants"
	"project-tracker/utils"
)

func ReadGenericJsonFile[T any](filePath string) []T {
	content, err := os.ReadFile(filePath)
	if len(content) == 0 {
		content = []byte("[]")
	}
	if err != nil {
		log.Fatal(err)
	}
	var parsedContents []T
	err = json.Unmarshal(content, &parsedContents)
	if err != nil {
		log.Fatal(err)
	}
	return parsedContents
}

func SaveJsonGeneric(data interface{}, path string) error {
	jsonData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, jsonData, 0644); err != nil {
		return err
	}
	return nil
}

func RetrieveProjectById(id int) utils.Project {
	return utils.Project{}
}

func RetrieveProjectByName(name string) utils.Project {
	return utils.Project{}
}

func DeleteProjectEntry(id int) {

}

func QueryProjectByDescription(desc string) utils.Project {
	return utils.Project{}
}

func SaveProject(project utils.Project) {
	projectList := ReadGenericJsonFile[utils.Project](constants.ProjectFilePath)
	projectList = append(projectList, project)
	err := SaveJsonGeneric(projectList, constants.ProjectFilePath)
	if err != nil {
		fmt.Printf("Fatal error occurred %s", err)
	}
	fmt.Printf("Successfully saved your Project!!")
}

func ListProjects() []utils.Project {
	projectList := ReadGenericJsonFile[utils.Project](constants.ProjectFilePath)
	for i, project := range projectList {
		fmt.Printf("%d. %s\n", i+1, project.ProjectName)
	}
	return projectList
}

func ListProjectsToStart() {
	fmt.Println("Following are your registered projects")
	projectList := ListProjects()
	fmt.Println("Which project you want to start?")
	var projectId int
	fmt.Scanf("%d", &projectId)
	StartProject(projectList[projectId-1])
}

func StartProject(project utils.Project) {
	ideList := ReadGenericJsonFile[utils.IDE](constants.IdeFilePath)
	ide := ideList[project.PreferredIDE-1]
	cmd := exec.Command(ide.ExecutablePath, project.Path)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
}
