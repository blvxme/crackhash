package config

func GetRedisConfigMap() (configMap map[string]string) {
	configMap = make(map[string]string)

	configMap["host"] = getEnvOrDefault("CRACKHASH_MANAGER_REDIS_HOST", DefaultRedisHost)
	configMap["port"] = getEnvOrDefault("CRACKHASH_MANAGER_REDIS_PORT", DefaultRedisPort)
	configMap["password"] = getEnvOrDefault("CRACKHASH_MANAGER_REDIS_PASSWORD", DefaultRedisPassword)
	configMap["db"] = getEnvOrDefault("CRACKHASH_MANAGER_REDIS_DB", DefaultRedisDb)

	return
}
