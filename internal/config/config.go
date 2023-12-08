package config

var config Config

type Config struct {
	AssetDir    string
	AssetPrefix string
	OutputFile  string
}

func GetConfig() *Config {
	return &config
}

func SetConfig(c Config) {
	config = c
}
