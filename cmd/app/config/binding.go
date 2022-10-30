package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	yamlConfig = "yaml"
)

type SystemConfig struct {
	configFile string
	configType string

	data []byte
}

func New() *SystemConfig {
	return &SystemConfig{}
}

func (c *SystemConfig) SetConfigFile(configFile string) {
	c.configFile = configFile
}

func (c *SystemConfig) SetConfigType(in string) {
	c.configType = in
}

func (c *SystemConfig) readInConfig() error {
	var err error
	c.data, err = ioutil.ReadFile(c.configFile)
	if err != nil {
		return err
	}
	return nil
}

func (c *SystemConfig) Binding(out interface{}) error {
	if err := c.readInConfig(); err != nil {
		return err
	}
	switch c.configType {
	case yamlConfig:
		if err := yaml.Unmarshal(c.data, out); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported config type %s", c.configType)
	}

	return nil
}
