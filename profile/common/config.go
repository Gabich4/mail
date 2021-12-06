package common

import (
	"os"
	"profile/utils"

	"github.com/mcuadros/go-defaults"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string            `yaml:"host" default:"localhost:3000"`
	Basepath string            `yaml:"basepath" default:"/api/v1"`
	Users    map[string]string `yaml:"users"`
}

var ServiceConfig = Config{
	Users: map[string]string{
		"admin": utils.HashPassword("admin"),
		"user1": utils.HashPassword("user1"),
		"user2": utils.HashPassword("user2"),
	},
}

func init() {
	if err := ServiceConfig.read(); err != nil {
		defaults.SetDefaults(&ServiceConfig)
	}
}

func (c Config) read() error {

	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &ServiceConfig)
}
