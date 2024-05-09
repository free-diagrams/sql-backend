package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig    `mapstructure:"app"`
	Logger LoggerConfig `mapstructure:"logger"`
	DB     DBConfig     `mapstructure:"database"`
	HTTP   HTTPConfig   `mapstructure:"http"`
}

type AppConfig struct {
	Name        string `mapstructure:"name" default:"sql-backend"`
	Environment string `mapstructure:"environment" required:"yes"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level" default:"info"`
}

type DBConfig struct {
	Host          string `mapstructure:"host" required:"yes"`
	Port          int    `mapstructure:"port" required:"yes"`
	Username      string `mapstructure:"username" required:"yes"`
	Password      string `mapstructure:"password" required:"yes"`
	Name          string `mapstructure:"name" required:"yes"`
	MigrationPath string
}

type HTTPConfig struct {
	Host string `mapstructure:"host" required:"yes"`
	Port string `mapstructure:"port" required:"yes"`
}

func Parse() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}

	config.DB.MigrationPath = "internal/infrastructure/repository/postgres/migration"

	return &config, nil
}
