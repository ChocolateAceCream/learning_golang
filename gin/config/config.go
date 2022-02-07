package config

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	Redis Redis `json: "redis"`
}

type Redis struct {
	Host     string `json: "host"`
	Password string `json: "password"`
	Key      string `json: "key"`
}

var path = "config.json"
var cfg *Config = nil
var once = sync.Once{}

func GetConfig() *Config {
	if cfg == nil {
		once.Do(func() { ParseCongif(path) })
	}
	return cfg
}

func ParseCongif(path string) *Config {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	if err = decoder.Decode(&cfg); err != nil {
		return nil
	}
	return cfg
}
