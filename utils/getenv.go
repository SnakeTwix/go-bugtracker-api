package utils

import (
	"fmt"
	"os"
)

func GetEnv(envName string) string {
	env := os.Getenv(envName)

	if env == "" {
		panic(fmt.Sprintf("%s is not defined", envName))
	}

	return env
}
