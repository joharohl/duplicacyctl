package config

import (
	"fmt"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	StoragePassword string   `yaml:"storage_password"`
	BackupLocations []string `yaml:"backup_locations"`
}

func (c Config) String() string {
	return fmt.Sprint(c.BackupLocations)
}

func New() *Config {
	return new(Config)
}

func (c *Config) Load(fs afero.Fs, filename string) {
	source, err := afero.ReadFile(fs, filename)
	if err != nil {
		log.Fatalf("Could not open config file: %s", filename)
	}
	err = yaml.Unmarshal(source, c)
	if err != nil {
		log.Fatalf("Could not parse config:\n%s", err)
	}
}
