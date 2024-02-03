package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

const PROJECT_ROOT_DIR = "/go/src/backend"

func init() {
	err := godotenv.Load(fmt.Sprintf("%s/.env", PROJECT_ROOT_DIR)); if err != nil {
		panic(err.Error())
	}
}
