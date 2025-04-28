package config

type Config struct {
	Port         int        `toml:"port"`
	Dir          string     `toml:"dir"`
	AppendOnly   bool       `toml:"appendonly"`
	ReplicaOf    string     `toml:"replicaof"`
	RequirePass  string     `toml:"requirepass"`
}

func LoadConfig(path string) (*Config, error) {
	// Read from TOML/YAML/JSON file
}