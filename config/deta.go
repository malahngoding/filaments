package config

func DetaID() string {
	return Config("DETA_PROJECT_ID")
}

func DetaKey() string {
	return Config("DETA_PROJECT_KEY")
}
