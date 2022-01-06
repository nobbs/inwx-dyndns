package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	User     string
	Password string
}

func (c *Configuration) validate() error {
	if c.User == "" {
		return errors.New("a username is required in the configuration files")
	}

	if c.Password == "" {
		return errors.New("a password is required in the configuration files")
	}

	return nil
}

func Configure(v *viper.Viper) *Configuration {
	// add default config paths
	v.AddConfigPath("./config")
	v.AddConfigPath("/config")

	// try to get credentials
	v.SetConfigName("credentials")
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Unable to read config file, %v", err)
	}

	// put everything into a `Configuration` object
	c := &Configuration{}
	if err := v.Unmarshal(c); err != nil {
		log.Fatalf("Unable to decode config into struct, %v", err)
	}

	// validate config
	if err := c.validate(); err != nil {
		log.Fatalf("Failed configuration validation, %v", err)
	}

	return c
}
