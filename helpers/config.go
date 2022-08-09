package helpers

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config data key
type Config struct {
	Host     string
	Database string
}

// ReadConfig from given path
func ReadConfig(path string, dest interface{}) error {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("unable to read config file. %s", err.Error())
	}

	if err := v.Unmarshal(dest); err != nil {
		return fmt.Errorf("unable to parse config file. %s", err.Error())
	}

	return nil
}
