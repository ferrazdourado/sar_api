package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    JWT      JWTConfig     `mapstructure:"jwt"`
    VPN      VPNConfig     `mapstructure:"vpn"`
}

type ServerConfig struct {
    Port    string `mapstructure:"port"`
    Mode    string `mapstructure:"mode"`
    Timeout int    `mapstructure:"timeout"`
}

type DatabaseConfig struct {
    URI      string `mapstructure:"uri"`
    Database string `mapstructure:"database"`
}

type JWTConfig struct {
    Secret     string `mapstructure:"secret"`
    ExpiresIn  int    `mapstructure:"expires_in"`
    SigningKey string `mapstructure:"signing_key"`
}

type VPNConfig struct {
    ConfigDir string `mapstructure:"config_dir"`
    LogDir    string `mapstructure:"log_dir"`
}

func LoadConfig(path string) (*Config, error) {
    viper.SetConfigFile(path)
    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        return nil, err
    }

    var config Config
    err = viper.Unmarshal(&config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}