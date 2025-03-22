package config

func GetManagerHost() (host string) {
	host = getEnvOrDefault("CRACKHASH_WORKER_MANAGER_HOST", DefaultManagerHost)

	return
}

func GetManagerPort() (port string) {
	port = getEnvOrDefault("CRACKHASH_WORKER_MANAGER_PORT", DefaultManagerPort)

	return
}
