package config

import "github.com/spf13/viper"

type Loader struct {
	key string
}

func NewLoader(key string) *Loader {
	return &Loader{key}
}

func (c *Loader) Get() string {
	return viper.GetString(c.key)
}

func (c *Loader) GetInt() int {
	return viper.GetInt(c.key)
}
