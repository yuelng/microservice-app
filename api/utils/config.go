package config

import (
	"log"

	"fmt"
	"github.com/spf13/viper"
	"os"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func init() {
	v := viper.New() // or direct use viper
	env := os.Getenv("GOENV")
	v.SetConfigType("toml")
	v.SetConfigName(env)

	viper.AddConfigPath("$GOPATH/src/api/config")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal("error on parsing configuration file")
	}
	config = v
}

// can set default value
func ConfigGetString(table string, key string) (value string, err error) {
	value = viper.GetStringMapString(table)[key]
	if value == "" {
		err = fmt.Errorf("%v[%v] is empty", table, key)
	}
	return
}

func GetConfig() *viper.Viper {
	return config
}
