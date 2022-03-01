package config

func InsteadToken() string {
	insteadToken := Config("INSTEAD_TOKEN")

	return insteadToken
}

func RedisAPI() string {
	return Config("REDIS_URL")
}
