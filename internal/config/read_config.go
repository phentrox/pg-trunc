package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig(name string) Config {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err.Error())
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err.Error())
	}

	return config
}
