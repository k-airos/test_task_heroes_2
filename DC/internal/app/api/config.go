package api

type APIConfig struct {
	Port       string `toml:"bind_addr_dc"`
	ConfigPath string
}

func NewConfig() *APIConfig {
	return &APIConfig{}
}
