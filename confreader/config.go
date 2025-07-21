package confreader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name      string
	Logger    LoggerConfig
	Databases DatabaseSection
	SecretKey string
}

type LoggerConfig struct {
	Level  string
	File   string
	Format string
}

type DatabaseSection struct {
	Mongos MongosConfig
}

type MongosConfig struct {
	SSO_Service MongoConfig
}
type MongoConfig struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type SQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type ConfigReader struct {
	Path string
}

func NewConfigReader(path string) *ConfigReader {
	return &ConfigReader{Path: path}
}

var globalConfig *Config

func (c *ConfigReader) LoadConfig() (*Config, error) {
	data, err := os.ReadFile(c.Path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("cannot parse config file: %w", err)
	}
	globalConfig = &cfg
	return globalConfig, nil
}
func GetConfig() *Config {
	return globalConfig
}
