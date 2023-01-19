package main

import (
	"time"
)

type DbConfig struct {
	Host            string
	Port            string
	User            string
	Pass            string
	Schema          string
	MaxIdleConn     int
	MaxOpenConn     int
	MaxConnLifetime time.Duration
	Debug           bool
}

var config *Config

type Config struct {
	Db *DbConfig
}

func AllConfig() *Config {
	return config
}

func ConfigLoad() {
	config = loadDefault()
}

func loadDefault() *Config {
	return &Config{
		Db: &DbConfig{
			Host:            "10.102.115.36",
			Port:            "3306",
			User:            "root",
			Pass:            "r00t",
			Schema:          "demo_db",
			MaxIdleConn:     1,
			MaxOpenConn:     2,
			MaxConnLifetime: 30,
			Debug:           true,
		},
	}
}
