package config

func GetNWorkers() (nWorkers string) {
	nWorkers = getEnvOrDefault("CRACKHASH_MANAGER_NWORKERS", DefaultNWorkers)

	return
}

func GetWorkerHost() (host string) {
	host = getEnvOrDefault("CRACKHASH_MANAGER_WORKER_HOST", DefaultWorkerHost)

	return
}

func GetWorkerPort() (port string) {
	port = getEnvOrDefault("CRACKHASH_MANAGER_WORKER_PORT", DefaultWorkerPort)

	return
}
