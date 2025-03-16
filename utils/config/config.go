package config

import (
	"log"

	"github.com/sharmayajush/lumel_crud/utils/constant"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("app")     // No .yaml extension here
	viper.SetConfigType("yaml")    // Set file type
	viper.AddConfigPath("./conf/") // Relative path
	viper.AddConfigPath(constant.ConfigFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("unable to read config file app.yaml. err:%s", err.Error())
	}
}
