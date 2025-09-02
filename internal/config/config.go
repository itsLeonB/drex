package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App
	DB
}

type App struct {
	Name    string        `default:"Drex"`
	Env     string        `default:"debug"`
	Port    string        `default:"50051"`
	Timeout time.Duration `default:"10s"`
}

func Load() Config {
	var app App
	envconfig.MustProcess("APP", &app)

	var db DB
	envconfig.MustProcess("DB", &db)

	return Config{
		App: app,
		DB:  db,
	}
}
