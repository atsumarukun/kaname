package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const PROJECT_ROOT_DIR = "/go/src/backend"

var (
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseTimeZone string
)

func init() {
	err := godotenv.Load(fmt.Sprintf("%s/.env", PROJECT_ROOT_DIR)); if err != nil {
		panic(err.Error())
	}

	DatabaseHost = os.Getenv("DATABASE_HOST")
	DatabasePort, err = strconv.Atoi(os.Getenv("DATABASE_PORT")); if err != nil {
		panic(err.Error())
	}
	DatabaseUser = os.Getenv("DATABASE_USER")
	DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	DatabaseName = os.Getenv("DATABASE_NAME")
	DatabaseTimeZone = os.Getenv("DATABASE_TIME_ZONE")
}
