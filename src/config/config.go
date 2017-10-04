package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	StoragePassword string   `yaml:"storage_password"`
	BackupLocations []string `yaml:"backup_locations"`
}

func (c Config) String() string {
	return fmt.Sprint(c.BackupLocations)
}

func LoadConfig(filename string) *Config {
	var config Config
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Could not open config file: %s", filename)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatalf("Could not parse config:\n%s", err)
	}
	return &config
}
