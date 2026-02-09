package config

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	EnvLocal = "local"
)

type Config struct {
	Environment string `mapstructure:"environment"`
	Postgres    PostgresConfig
	Redis       RedisConfig
	HTTP        HTTPConfig
	WebhookURL  string
}

type PostgresConfig struct {
	Username string
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string
	SSLMode  string `mapstructure:"sslmode"`
	Password string
}

type RedisConfig struct {
	Host             string        `mapstructure:"host"`
	Port             int           `mapstructure:"port"`
	CacheTTL         time.Duration `mapstructure:"cacheTTL"`
	Password         string
	StreamKey        string `mapstructure:"streamKey"`
	Group            string `mapstructure:"group"`
	Consumer         string `mapstructure:"consumer"`
	MaxAttempts      uint   `mapstructure:"maxAttempts"`
	DeadLetterStream string `mapstructure:"deadLetterStream"`
}

type HTTPConfig struct {
	Host               string        `mapstructure:"host"`
	Port               int           `mapstructure:"port"`
	ReadTimeout        time.Duration `mapstructure:"readTimeout"`
	WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
	MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
}

func Init(configsDir string) (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	if err := parseConfigFile(configsDir, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	cfg.SetDefaults()
	cfg.OverrideFromEnv()

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &cfg, nil
}

func (c *Config) SetDefaults() {
	if c.Postgres.SSLMode == "" {
		c.Postgres.SSLMode = "disable"
	}
	if c.Redis.MaxAttempts == 0 {
		c.Redis.MaxAttempts = 5
	}
	if c.Redis.DeadLetterStream == "" && c.Redis.StreamKey != "" {
		c.Redis.DeadLetterStream = c.Redis.StreamKey + "_dl"
	}
	if c.HTTP.Port == 0 {
		c.HTTP.Port = 8080
	}
	if c.HTTP.ReadTimeout <= 0 {
		c.HTTP.ReadTimeout = 5 * time.Second
	}
	if c.HTTP.WriteTimeout <= 0 {
		c.HTTP.WriteTimeout = 5 * time.Second
	}
	if c.HTTP.MaxHeaderMegabytes <= 0 {
		c.HTTP.MaxHeaderMegabytes = 1 << 20
	}
}

func (c *Config) OverrideFromEnv() {
	if val := os.Getenv("POSTGRES_USER"); val != "" {
		c.Postgres.Username = val
	}
	if val := os.Getenv("POSTGRES_PASSWORD"); val != "" {
		c.Postgres.Password = val
	}
	if val := os.Getenv("POSTGRES_DB"); val != "" {
		c.Postgres.Name = val
	}
	if val := os.Getenv("REDIS_PASSWORD"); val != "" {
		c.Redis.Password = val
	}
	if val := os.Getenv("WEBHOOK_URL"); val != "" {
		c.WebhookURL = val
	}
}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read main config: %w", err)
	}

	if env == "" || env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)
	if err := viper.MergeInConfig(); err != nil {
		return fmt.Errorf("failed to read %s config: %w", env, err)
	}

	return nil
}

func (c *Config) Validate() error {
	if c.Environment == "" {
		return fmt.Errorf("environment is not set")
	}

	if c.Postgres.Host == "" {
		return fmt.Errorf("postgres host is required")
	}
	if c.Postgres.Port <= 0 || c.Postgres.Port > 65535 {
		return fmt.Errorf("postgres port must be between 1 and 65535")
	}
	if c.Postgres.Username == "" {
		return fmt.Errorf("postgres username is required")
	}
	if c.Postgres.Password == "" {
		return fmt.Errorf("postgres password is required")
	}
	if c.Postgres.Name == "" {
		return fmt.Errorf("postgres database name is required")
	}
	if c.Postgres.SSLMode == "" {
		return fmt.Errorf("postgres sslmode is required")
	}

	if c.Redis.Host == "" {
		return fmt.Errorf("redis host is required")
	}
	if c.Redis.Port <= 0 || c.Redis.Port > 65535 {
		return fmt.Errorf("redis port must be between 1 and 65535")
	}
	if c.Redis.CacheTTL <= 0 {
		return fmt.Errorf("redis cacheTTL must be > 0")
	}
	if c.Redis.StreamKey == "" || c.Redis.Group == "" || c.Redis.Consumer == "" {
		return fmt.Errorf("redis streamKey, group, and consumer are required")
	}
	if c.Redis.MaxAttempts == 0 {
		return fmt.Errorf("redis maxAttempts must be > 0")
	}
	if c.Redis.DeadLetterStream == "" {
		return fmt.Errorf("redis deadLetterStream is required")
	}

	if c.HTTP.Port <= 0 || c.HTTP.Port > 65535 {
		return fmt.Errorf("http port must be between 1 and 65535")
	}
	if c.HTTP.ReadTimeout <= 0 || c.HTTP.WriteTimeout <= 0 {
		return fmt.Errorf("http timeouts must be > 0")
	}
	if c.HTTP.MaxHeaderMegabytes <= 0 {
		return fmt.Errorf("http maxHeaderMegabytes must be > 0")
	}

	if c.WebhookURL == "" {
		return fmt.Errorf("webhook URL is required")
	}
	if _, err := url.ParseRequestURI(c.WebhookURL); err != nil {
		return fmt.Errorf("webhook URL is invalid: %w", err)
	}

	return nil
}
