package common

import (
	"os"

	"github.com/mcuadros/go-defaults"
	"gopkg.in/yaml.v3"
)

type Config struct {
	// Host is a service address (e.g. "localhost:5000")
	Host string `yaml:"host" default:"localhost:5000"`

	// Basepath is a postfix for accessing API (e.g. "/api/v1")
	Basepath string `yaml:"basepath" default:"/api/v1"`

	// MailsenderConnection is a HTTP client to mailsender service.
	MailsenderConnection string `yaml:"mailsenderconnection" default:"http://localhost:4000"`
}

var ServiceConfig Config

func init() {
	if err := ServiceConfig.read(); err != nil {
		defaults.SetDefaults(&ServiceConfig)
	}
}

func (c *Config) read() error {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &ServiceConfig); err != nil {
		return err
	}

	return nil
}
