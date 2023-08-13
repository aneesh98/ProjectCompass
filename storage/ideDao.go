package storage

import (
	"fmt"
	"project-tracker/constants"
	"project-tracker/utils"
)

func SaveIde(ide utils.IDE) {
	ideList := ReadGenericJsonFile[utils.IDE](constants.IdeFilePath)
	ideList = append(ideList, ide)
	err := SaveJsonGeneric(ideList, constants.IdeFilePath)
	if err != nil {
		fmt.Printf("Fatal error occurred %s", err)
	}
	fmt.Printf("Successfully saved your IDE!!")
}

func ListIdes() {
	ideList := ReadGenericJsonFile[utils.IDE](constants.IdeFilePath)
	for i, ide := range ideList {
		fmt.Printf("%d. %s\n", i+1, ide.IdeName)
	}
}
