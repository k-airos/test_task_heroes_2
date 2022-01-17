package storage

type Config struct {
	ApplyURI       string `toml:"apply_uri"`
	DBname         string `toml:"db_name"`
	CollectionName string `toml:"marvel_collection_name"`
}

func NewConfig() *Config {
	return &Config{}
}
