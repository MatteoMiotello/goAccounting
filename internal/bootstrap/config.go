package bootstrap

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type Config struct {
}

// read configs from .env and all files contained in /config folder
func (c Config) Init() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	files, err := os.ReadDir("./config/")

	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath("./config")

	for _, file := range files {
		name := file.Name()

		parts := strings.Split(name, ".")
		n := parts[0]
		e := parts[1]

		viper.SetConfigName(n)
		viper.SetConfigType(e)
		err := viper.MergeInConfig()

		if err != nil {
			return err
		}
	}

	return nil
}
