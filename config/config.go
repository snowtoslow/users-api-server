package config

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
)

type Config struct {
	AppPort      string `mapstructure:"app_port" json:"appPort" yaml:"appPort"`
	DatabaseName string `mapstructure:"databaseName" json:"databaseName" yaml:"databaseName"`
	SecretKey    string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`
	Issuer       string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}

func ParseConfigFile(path string) (Config, error) {
	log.Println(path)
	viper.AddConfigPath(filepath.Dir(path))
	if fileExtension := filepath.Ext(path); fileExtension != "" {
		viper.SetConfigType(strings.Trim(fileExtension, "."))
	}
	_, nameWithExtension := filepath.Split(path)
	if name := strings.TrimSuffix(nameWithExtension, filepath.Ext(path)); name != "" {
		viper.SetConfigName(name)
	}

	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
