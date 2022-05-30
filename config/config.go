package config

import "os"

func IsTestRun() bool {
	return os.Getenv("EXECUTION_ENVIRONMENT") == "test"
}
