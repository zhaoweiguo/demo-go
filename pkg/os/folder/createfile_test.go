package folder

import (
	"log"
	"os"
	"testing"
)

func TestCreateFolderIfNotExists(t *testing.T) {
	resultsPath := "./folder/"

	if _, err := os.Stat(resultsPath); os.IsNotExist(err) {
		err = os.Mkdir(resultsPath, 0777)
		if err != nil {
			log.Println(err, "Failed to create allure-results folder")
		}
	}
}
