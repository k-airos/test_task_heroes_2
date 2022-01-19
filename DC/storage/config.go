package storage

type Config struct {
	ApplyURI       string `toml:"apply_uri"`
	DBname         string `toml:"db_name"`
	CollectionName string `toml:"dc_collection_name"`
}

func NewConfig() *Config {
	return &Config{}
}
