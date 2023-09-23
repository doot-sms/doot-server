package utils

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBUrl          string `mapstructure:"DB_URL"`
	DBName         string `mapstructure:"DB_NAME"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBRootPassword string `mapstructure:"DB_ROOT_PASSWORD"`
	DBType         string `mapstructure:"DB_TYPE"`

	StageStatus       string        `mapstructure:"STAGE_STATUS"`
	ServerHost        string        `mapstructure:"SERVER_HOST"`
	ServerPort        string        `mapstructure:"SERVER_PORT"`
	ServerReadTimeout time.Duration `mapstructure:"SERVER_READ_TIMEOUT"`

	DBMaxConnections         int `mapstructure:"DB_MAX_CONNECTIONS"`
	DBMaxIdleConnections     int `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	DBMaxLifetimeConnections int `mapstructure:"DB_MAX_LIFETIME_CONNECTIONS"`

	SupaBasePassword  string `mapstructure:"SUPABASE_PASSWORD"`
	DootEncryptionKey string `mapstructure:"DOOT_ENCRYPTION_KEY"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	// viper.AddConfigPath("../")
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
