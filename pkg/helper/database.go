package helper

import (
	"fmt"
)

func PsqlConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		EnvString("POSTGRES_USERNAME"),
		EnvString("POSTGRES_PASSWORD"),
		EnvString("POSTGRES_HOST"),
		EnvInt("POSTGRES_PORT"),
		EnvString("POSTGRES_DATABASE"),
	)
}
