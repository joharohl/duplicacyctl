package config_test

import (
	"github.com/joharohl/duplicacyctl/src/config"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"testing"
)

var rawYaml = []byte(`
storage_password: secrets_galore

backup_locations:
- /home/johan
- /media/secondary/Pictures
`)

func TestLoad(t *testing.T) {
	fs := afero.NewMemMapFs()
	fs.MkdirAll("/tmp", 0644)
	afero.WriteFile(fs, "/tmp/test.yaml", rawYaml, 0644)
	cf := config.New()
	cf.Load(fs, "/tmp/test.yaml")

	assert.Equal(t, "secrets_galore", cf.StoragePassword)
	assert.Equal(t, []string{"/home/johan", "/media/secondary/Pictures"}, cf.BackupLocations)
}
