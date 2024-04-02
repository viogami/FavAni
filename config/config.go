package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	Redis  RedisConfig  `yaml:"redis"`
	GRPC   GRPCConfig   `yaml:"grpc"`
}

type ServerConfig struct {
	Env       string `yaml:"env"`
	Address   string `yaml:"address"`
	Port      int    `yaml:"port"`
	JWTSecret string `yaml:"jwtSecret"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Migrate  bool   `yaml:"migrate"`
}

type RedisConfig struct {
	Enable   bool   `yaml:"enable"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type GRPCConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// 读取配置文件 app.yaml
func ConfigParse(appConfig *string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(*appConfig)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
