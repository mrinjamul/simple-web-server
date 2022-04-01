package main

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config model structure for json and yaml
type Config struct {
	Port    string `json:"port"`
	Dir     string `json:"dir"`
	SslKey  string `json:"ssl_key"`
	SslCert string `json:"ssl_cert"`
	HTTPS   bool   `json:"https"`
}

// NewConfig creates new config
func NewConfig() *Config {
	return &Config{
		Port:    "8080",
		Dir:     "./",
		SslKey:  "",
		SslCert: "",
		HTTPS:   false,
	}
}

// LoadConfig loads config from json file
func LoadConfigfromJSON(file string) (*Config, error) {
	config := NewConfig()
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return config, err
	}
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(config); err != nil {
		return config, err
	}
	return config, nil
}

// SaveConfig saves config to json file
func SaveConfigfromJSON(file string, config *Config) {
	configFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	jsonEncoder := json.NewEncoder(configFile)
	if err := jsonEncoder.Encode(config); err != nil {
		log.Fatal(err)
	}
}

// LoadYAML loads config from yaml file
func LoadConfigfromYAML(file string) (*Config, error) {
	config := NewConfig()
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return config, err
	}
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	yamlParser := yaml.NewDecoder(configFile)
	if err := yamlParser.Decode(config); err != nil {
		return config, err
	}
	return config, nil
}

// SaveYAML saves config to yaml file
func SaveConfigfromYAML(file string, config *Config) {
	configFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	yamlEncoder := yaml.NewEncoder(configFile)
	if err := yamlEncoder.Encode(config); err != nil {
		log.Fatal(err)
	}
}

// GetConfig gets config from standard config file locations
// 1. /etc/sws/config.json
// 2. $HOME/.sws/config.json
// 3. ./config.json
// 4. /etc/sws/config.{yml,yaml}
// 5. $HOME/.sws/config.{yml,yaml}
// 6. config.{yml,yaml}
func GetConfig() *Config {
	// config := NewConfig()
	// Get Home directory
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	config, err := LoadConfigfromJSON("/etc/sws/config.json")
	if err != nil {
		config, err = LoadConfigfromJSON(home + "/.sws/config.json")
		if err != nil {
			config, err = LoadConfigfromJSON("./config.json")
			if err != nil {
				config, err = LoadConfigfromYAML("/etc/sws/config.yml")
				if err != nil {
					config, err = LoadConfigfromYAML(home + "/.sws/config.yml")
					if err != nil {
						config, err = LoadConfigfromYAML("./config.yml")
						if err != nil {
							config, err = LoadConfigfromYAML("/etc/sws/config.yaml")
							if err != nil {
								config, err = LoadConfigfromYAML(home + "/.sws/config.yaml")
								if err != nil {
									config, err = LoadConfigfromYAML("./config.yaml")
									if err != nil {
										config = NewConfig()
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return config
}
