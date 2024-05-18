package config

import (
	"github.com/free-diagrams/sql-backend/internal/domain/errors"
	"github.com/free-diagrams/sql-backend/pkg/errs"
	"github.com/spf13/viper"
	"time"
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
	Host            string        `mapstructure:"host" required:"yes"`
	Port            int           `mapstructure:"port" required:"yes"`
	Username        string        `mapstructure:"username" required:"yes"`
	Password        string        `mapstructure:"password" required:"yes"`
	Name            string        `mapstructure:"name" required:"yes"`
	MaxOpenConns    int           `mapstructure:"max-open-conns" default:"25"`
	MaxIdleConns    int           `mapstructure:"max-idle-conns" default:"10"`
	ConnMaxLifeTime time.Duration `mapstructure:"conn-max-life-time" default:"5m"`
	MigrationPath   string
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
		return nil, errs.WrapErrorString(errors.InternalServer, "failed to read config file")
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, errs.WrapErrorString(errors.InternalServer, "failed to unmarshal config")
	}

	config.DB.MigrationPath = "internal/infrastructure/repository/postgres/migration"

	return &config, nil
}
