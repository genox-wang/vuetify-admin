package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./app/config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
}

// GetString get string config by key
func GetString(key string) string {
	return viper.GetString(key)
}

// GetBool get bool config by key
func GetBool(key string) bool {
	return viper.GetBool(key)
}
