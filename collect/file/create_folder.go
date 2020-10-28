package file

import (
	"log"
	"os"
)

func createFolderIfNotExists() {
	resultsPath := "./folder/"

	if _, err := os.Stat(resultsPath); os.IsNotExist(err) {
		err = os.Mkdir(resultsPath, 0777)
		if err != nil {
			log.Println(err, "Failed to create allure-results folder")
		}
	}
}
