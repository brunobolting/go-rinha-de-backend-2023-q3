package env

import "os"

var (
	REDIS      string
	POSTGRESQL string
	PORT       string
)

func init() {
	REDIS = os.Getenv("REDIS")
	POSTGRESQL = os.Getenv("POSTGRESQL")
	PORT = os.Getenv("PORT")
}
