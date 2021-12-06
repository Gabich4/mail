package config

import (
	"auth/utils"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	init         bool
	Port         string
	Users        map[string]string
	JWTSecretKey string `yaml:"jwt_secret_key"`
}

var config = Config{
	init: false,
	Port: "8000",
	Users: map[string]string{
		"admin": utils.HashPassword("admin"),
		"user1": utils.HashPassword("user1"),
		"user2": utils.HashPassword("user2"),
	},
	JWTSecretKey: "my_secret_key",
}

func GetConfig() *Config {
	if !config.init {
		config.read()
		config.init = true
	}
	return &config
}

func (c Config) read() {

	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		utils.Logger.Print(err)
	}

	if len(data) == 0 {
		config.write()
	} else {
		err := yaml.Unmarshal(data, &config)
		if err != nil {
			utils.Logger.Fatal(err)
		}
	}
}

func (c Config) write() {

	data, err := yaml.Marshal(&config)
	if err != nil {
		utils.Logger.Fatal(err)
	}

	err = os.WriteFile("./config.yaml", data, 0666)
	if err != nil {
		utils.Logger.Print(err)
	}
}
