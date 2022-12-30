package bootstrap

import "github.com/spf13/viper"

type Config struct {
}

func (c Config) Init() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	return nil
}
