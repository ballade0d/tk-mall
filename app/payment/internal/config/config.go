package config

import (
	"github.com/BurntSushi/toml"
	"github.com/google/wire"
	"log"
)

type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type DatabaseConfig struct {
	Driver string `toml:"driver"`
	Source string `toml:"source"`
}

type RedisConfig struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	Database int    `toml:"database"`
}

type ElasticSearchConfig struct {
	Addresses []string `toml:"addresses"`
	APIKey    string   `toml:"api_key"`
	Indices   string   `toml:"indices"`
}

type Config struct {
	Server        ServerConfig        `toml:"server"`
	Database      DatabaseConfig      `toml:"database"`
	Redis         RedisConfig         `toml:"redis"`
	ElasticSearch ElasticSearchConfig `toml:"elasticsearch"`
}

var ProviderSet = wire.NewSet(NewConfig)

func NewConfig() (*Config, error) {
	var conf Config
	// 加载配置文件
	_, err := toml.DecodeFile("config-local.toml", &conf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &conf, nil
}
