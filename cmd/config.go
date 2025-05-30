package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddr   string
	PostgresConn string
	RedisConn    string
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	config := Config{}

	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to read config: %w", err)
	}

	envPrefix := strings.ReplaceAll(strings.ToUpper(AppName), "-", "_")

	viper.SetEnvPrefix(envPrefix)
	viper.KeyDelimiter("__")
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}
