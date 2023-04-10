package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	configs "platform/pkg/config"
)

type (
	Config struct {
		configs.App  `yaml:"app"`
		configs.HTTP `yaml:"http"`
		configs.Log  `yaml:"logger"`
		UsersClient  `yaml:"users_client"`
		DataSource   `yaml:"datasource"`
		RabbitMQ     `yaml:"rabbitmq"`
	}

	DataSource struct {
		Type  string `env-required:"true" yaml:"type" env:"TYPE"`
		Mysql Mysql  `env-required:"true" yaml:"mysql" env:"MYSQL"`
		PG    PG     `env-required:"true" yaml:"postgres" env:"POSTGRES"`
	}

	PG struct {
		PoolMax int                  `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		DsnURL  configs.DBConnString `env-required:"true" yaml:"dsn_url" env:"PG_DSN_URL"`
	}

	Mysql struct {
		MaxOpenConns int                  `env-required:"true" yaml:"max_open_conns" env:"MAX_OPEN_CONNS"`
		MaxIdleConns int                  `env-required:"true" yaml:"max_idle_conns" env:"MAX_IDLE_CONNS"`
		URL          configs.DBConnString `env-required:"true" yaml:"url" env:"URL"`
	}

	RabbitMQ struct {
		URL string `env-required:"true" yaml:"url" env:"RABBITMQ_URL"`
	}

	UsersClient struct {
		URL string `env-required:"true" yaml:"url" env:"USER_CLIENT_URL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// debug
	fmt.Println(dir)

	err = cleanenv.ReadConfig(dir+"/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
