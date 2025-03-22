package config

import "os"

func getEnvOrDefault(key string, defaultValue string) (value string) {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = defaultValue
	}

	return
}
