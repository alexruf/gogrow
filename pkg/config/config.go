package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type Config struct {
	logger *log.Logger
}

func NewConfig(logger *log.Logger) Config {
	return Config{logger: logger}
}

func (c Config) Init() error {
	c.logger.Println("Initializing config...")
	dir, err := os.Getwd()
	if err != nil {
		c.logger.Fatal(err)
	}
	usr, err := user.Current()
	if err != nil {
		c.logger.Fatal(err)
	}

	c.setDefaultConfig()
	viper.SetEnvPrefix("grow")
	viper.AutomaticEnv()
	viper.SetConfigName("gogrow")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)
	viper.AddConfigPath(filepath.Join(usr.HomeDir))

	if err := viper.ReadInConfig(); err == nil {
		c.logger.Println("Using config file: ", viper.ConfigFileUsed())
	}

	return nil
}

func (c Config) setDefaultConfig() {
	viper.SetDefault(ShootFromHour, 9)
	viper.SetDefault(ShootUntilHour, 22)
	viper.SetDefault(ShootIntervalMinutes, 30)
}

func (c Config) Write() error {
	c.logger.Println("Writing config file...")
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}
