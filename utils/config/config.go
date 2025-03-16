package config

import (
	"log"

	"github.com/sharmayajush/lumel_crud/utils/constant"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.AddConfigPath(constant.ConfigFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("unable to read config file app.yaml. err:%s", err.Error())
	}
}
