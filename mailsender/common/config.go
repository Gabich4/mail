package common

import (
	"os"

	"github.com/mcuadros/go-defaults"
	"gopkg.in/yaml.v3"
)

// Config is a structure used to represents the service
// configuration.
type Config struct {
	// Host is a service address (e.g. "localhost:8000")
	Host string `yaml:"Host" default:"localhost:4000"`

	// Basepath is a postfix for accessing API (e.g. "/api/v1")
	Basepath string `yaml:"Basepath" default:"/api/v1"`

	// ProfileClient is a http.Client to profile service
	ProfileConnection string `yaml:"ProfileConnection" default:"http://localhost:3000"`

	MongoConnection string `yaml:"MongoConnection" default:"mongodb://localhost:27017"`
}

// ServiceConfig is a global variable containing all info about service.
var ServiceConfig Config

func init() {
	if err := ServiceConfig.read(); err != nil {
		defaults.SetDefaults(&ServiceConfig)
	}
}

// read get the data from config file or returns error
// if file does not exist.
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
