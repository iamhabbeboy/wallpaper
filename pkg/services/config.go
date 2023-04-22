package services

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigService struct {
	db *viper.Viper
}

func LoadConfig() *viper.Viper {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath("./.wallpaper_config")

	v.SetDefault("config.max_image", 5)
	v.SetDefault("api.url", "https://api.unsplash.com/")
	v.SetDefault("api.access_key", "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY")
	v.SetDefault("api.secret_key", "pseMeAYqR4G1I8cx8vbwkm4HTs1o56NzW6ZiKGHCMNs")
	v.SetDefault("config.image_path", ".wallpaper_config/images")

	// Write the configuration options to a YAML file
	if err := v.WriteConfigAs("./.wallpaper_config/config.yaml"); err != nil {
		log.Fatalf("Error writing configuration file: %s", err)
	}

	// Read the configuration options from the YAML file
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
	}
	return v
}

func NewConfigService() *ConfigService {
	config := LoadConfig()
	return &ConfigService{
		db: config,
	}
}

func (c *ConfigService) Get(key string) string {
	return c.db.GetString(key)
}

func (c *ConfigService) Set(key string, value string) error {
	c.db.Set(key, value)
	err := c.db.WriteConfig()
	if err != nil {
		fmt.Println("Error writing config file:", err)
	}
	return nil
}