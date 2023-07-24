package lib

import (
	"log"
	"os"
)

func GetProjectPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	return path
}
