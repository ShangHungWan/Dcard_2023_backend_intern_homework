package env

import (
	"io/ioutil"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	var path string
	if envFileInNowDir() {
		path = "./.env"
	} else {
		path = "../.env"
	}

	err := godotenv.Load(path)
	if err != nil {
		log.Print(err.Error())
		log.Fatal("Error loading .env file")
	}
}

func envFileInNowDir() bool {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == ".env" {
			return true
		}
	}

	return false
}
