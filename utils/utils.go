package utils

import "os"

func GetEnvString(name string) string {
	return os.Getenv(name)
}
