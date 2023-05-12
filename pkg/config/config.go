package config

import (
	"fmt"
	"log"

	"github.com/alwisisva/twitter-app/pkg/helper"
	"github.com/joho/godotenv"
)

func LoadEnv(filepath string) {
	err := godotenv.Load(fmt.Sprintf("%s/%s", helper.ProjectRootPath, filepath))
	if err != nil {
		log.Printf("failed to read file %s \n", filepath)
		log.Println("reading environment variable")
	}

	log.Printf("Succeed read environment variable from file %s \n\n", filepath)

}
