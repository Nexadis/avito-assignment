package cfg

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
	EnvDev   = "dev"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"local"`
	DSN        string `yaml:"dsn" env:"DSN" env-required:"true"`
	HTTPServer `yaml:"http_server"`
	Cache      `yaml:"cache"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"ADDRESS" env-default:"localhost:8080"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"4s"`
}

type Cache struct {
	TTL  time.Duration `yaml:"ttl" env:"TTL" env-default:"5m"`
	Size uint          `yaml:"size" env:"SIZE" env-default:"1000"`
}

func MustLoad() *Config {
	conf, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		log.Fatal("CONFIG_FILE is not set")
	}
	_, err := os.Stat(conf)
	if errors.Is(err, os.ErrExist) {
		log.Fatalf("config file '%s' does not exist", conf)
	}
	cfg := Config{}
	err = cleanenv.ReadConfig(conf, &cfg)
	if err != nil {
		log.Fatalf("environment error: %s", err)
	}
	return &cfg
}
