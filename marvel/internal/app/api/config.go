package api

type APIConfig struct {
	Port       string `toml:"bind_addr_marvel"`
	ConfigPath string
}

func NewConfig() *APIConfig {
	return &APIConfig{}
}
