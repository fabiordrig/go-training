package utils

import "github.com/spf13/viper"

type Config struct {
	// ServerAddress is the address of the server
	ServerAddress string `mapstructure:"server_address" default:":8080"`
	// DBDriver is the driver for the database
	DBDriver string `mapstructure:"DB_DRIVER" default:"postgres"`
	// DBSource is the source for the database
	DBSource string `mapstructure:"DB_URL"`
}

// LoadConfig loads the configuration from the environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
