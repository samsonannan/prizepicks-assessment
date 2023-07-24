package db

import "github.com/samsonannan/prizepicks-assessment/pkg/ent"

var (
	PostgresCfg    PostgresConfig
	PostgresClient *ent.Client
)

type PostgresConfig struct {
	Host     string `mapstructure:"host" toml:"host" json:"host"`
	Port     string `mapstructure:"port" toml:"port" json:"port"`
	User     string `mapstructure:"user" toml:"user" json:"user"`
	Password string `mapstructure:"password" toml:"password" json:"password"`
	Database string `mapstructure:"database" toml:"database" json:"database"`
}
