package config

func GetAlphabet() (alphabet string) {
	alphabet = getEnvOrDefault("CRACKHASH_MANAGER_ALPHABET", DefaultAlphabet)

	return
}
